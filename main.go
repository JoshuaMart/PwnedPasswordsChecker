package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// READ INPUT FLAG
	inputFile, hashType, outputFile, hashFile := flagCheck()

	// OPEN INPUT HASH FILE
	hashs, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer hashs.Close()
	scanner := bufio.NewScanner(hashs)

	// PREPARE OUTPUT FILE
	out, err := os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// OPEN HIBP HASH FILE
	f, _ := os.OpenFile(hashFile, os.O_CREATE|os.O_RDONLY, 0400)
	defer f.Close()

	// START TIME COUNTER
	startTimer := time.Now()

	// COUNT THE NUMBER OF BYTES FOR HIBP HASH FILE
	fStat, _ := f.Stat()

	v := 0 // Hashs analyzed
	x := 0 // Hashs found

	// START OF FILE IN INT64
	start := int64(0)

	var byteLength int64
	var hashLength int

	// DEFINE HASH TYPE
	if hashType == "NTLM" {
		byteLength = 20 // Necessary byte length for extract in binary
		hashLength = 16 // Hash length in bytes inside the binary
	} else if hashType == "SHA1" {
		byteLength = 24
		hashLength = 20
	}

	for scanner.Scan() {
		hash := scanner.Text()

		// END OF FILE IN INT64
		end := fStat.Size()

		// RETRIEVES THE MIDDLE OF FILE IN BYTES
		mid := start + (end-start)/2

		for mid%byteLength != 0 {
			mid += 1
		}

		f.Seek(mid, 0)

		// Extract hash from the line
		extract := make([]byte, byteLength) // 20 for NTLM && 24 for SHA1
		f.Read(extract)
		hashExtract := hex.EncodeToString(extract[:hashLength])
		occurenceHash := binary.BigEndian.Uint32(extract[hashLength:byteLength])

		for {
			if hash == hashExtract {
				x++
				out.WriteString(hashExtract + ":" + strconv.FormatUint(uint64(occurenceHash), 10) + "\n") // Write HASH & OCCURENCE TO OUTPUT FILE
				break
			} else if hash > hashExtract {
				start = mid
				mid = start + (end-start)/2
				for mid%byteLength != 0 { // Necessary to make sure the position is at the beginning of a hash.
					mid -= 1
				}
			} else {
				end = mid
				mid = start + (end-start)/2
				for mid%byteLength != 0 { // Necessary to make sure the position is at the beginning of a hash.
					mid -= 1
				}
			}
			f.Seek(mid, 0) // Repositioning in the middle
			extract = make([]byte, byteLength)
			f.Read(extract)
			hashExtract = hex.EncodeToString(extract[:hashLength])
			occurenceHash = binary.BigEndian.Uint32(extract[hashLength:byteLength])

			if start == mid || end == mid {
				if hash == hashExtract {
					out.WriteString(hashExtract + ":" + strconv.FormatUint(uint64(occurenceHash), 10) + "\n") //Write HASH & OCCURENCE TO OUTPUT FILE
					x++
					break
				} else {
					break
				}
			}
		}
		v++
	}

	// Time calculation
	fmt.Printf("%d hashs analyzed and %d found \n", v, x)
	log.Printf("It took : %s", time.Since(startTimer))
}

// Check if flag is provided
func flagCheck() (string, string, string, string) {
	flag.Parse()

	args := flag.Args()
	if len(args) < 4 {
		fmt.Println("Input and/or hashType and/or output and/or hash file is missing.")
		os.Exit(1)
	}
	return args[0], args[1], args[2], args[3]
}
