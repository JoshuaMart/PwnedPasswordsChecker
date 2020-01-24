package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/stoicperlman/fls"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 3 {
		fmt.Println("Input and/or output and/or hash file is missing.")
		os.Exit(1)
	}
	inputFile := args[0]
	outputFile := args[1]
	hashFile := args[2]

	// Open & read hashs files
	hashs, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer hashs.Close()
	scanner := bufio.NewScanner(hashs)

	out, err := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Starting the time counter
	startTimer := time.Now()
	// Open file
	f, _ := os.OpenFile(hashFile, os.O_CREATE|os.O_RDONLY, 0400)
	defer f.Close()
	file := fls.LineFile(f)

	// Regex for extraction
	r, _ := regexp.Compile("([A-Z0-9]{32}):[0-9]{1,}")

	// Whence is the point of reference for offset
	// 0 = Beginning of file | 1 = Current position | 2 = End of file
	var whence int = 0

	// Count the number of bytes
	fStat, _ := f.Stat()

	v := 0
	x := 0

	for scanner.Scan() {
		hash := scanner.Text()
		// Start of file in Int64
		start := int64(whence)
		//End of file in Int64
		end := fStat.Size()
		// Retrieves the middle of the file (in bytes)
		mid := start + (end-start)/2

		// Go to the middle at the beginning of the line
		file.Seek(mid, whence)
		file.SeekLine(0, io.SeekCurrent)

		// Extract hash from the line
		length1 := make([]byte, 32)
		n1, _ := f.Read(length1)
		extract := string(length1[:n1])
		for {
			if hash == extract {
				start, _ = file.SeekLine(0, io.SeekCurrent)
				length1 = make([]byte, 42)
				n1, _ = f.Read(length1)
				extract = string(length1[:n1])

				out.WriteString(r.FindString(extract) + "\n")
				x++

				break
			} else if hash > extract {
				start = mid
				mid = start + (end-start)/2

				file.Seek(mid, whence)
				file.SeekLine(0, io.SeekCurrent)

				length1 = make([]byte, 32)
				n1, _ = f.Read(length1)
				extract = string(length1[:n1])
			} else {
				end = mid
				mid = start + (end-start)/2

				file.Seek(mid, whence)
				file.SeekLine(0, io.SeekCurrent)

				length1 = make([]byte, 32)
				n1, _ = f.Read(length1)
				extract = string(length1[:n1])
			}

			if start == mid || end == mid {
				break
			}
		}
		v++
	}

	// Time calculation
	fmt.Printf("%d hashs analyzed and %d found \n", v, x)
	log.Printf("It took : %s", time.Since(startTimer))
}
