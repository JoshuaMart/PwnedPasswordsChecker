# PwnedPasswordsChecker
![banner](https://zupimages.net/up/20/04/zsp8.png)![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg) ![made-with-go](https://img.shields.io/badge/made%20with-go-blue)  ![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)

**PwnedPasswordsChecker** is a tool that checks if the hash of a known password (in SHA1 or NTLM format) is present in the list of I Have Been Pwned leaks and the number of occurrences.

You can download the hash-coded version for SHA1 [here](https://downloads.pwnedpasswords.com/passwords/pwned-passwords-sha1-ordered-by-hash-v7.7z) or the hash-coded version for NTLM [here](https://downloads.pwnedpasswords.com/passwords/pwned-passwords-ntlm-ordered-by-hash-v7.7z)

Once the list is downloaded it is then necessary to convert it to binary by using my other tool [HIBP PasswordList Slimmer](https://github.com/JoshuaMart/HIBP_PasswordList_Slimmer)

*This script only works with the HIBP version sorted by hash and entry hashes must be in lowercase and preferably ordered by hashs*

## Usage :
```
./PwnedPasswordsChecker {InputHashList} {HashType} {OutputFile} {CompressedHIBPHashList}
./PwnedPasswordsChecker .\NTLM_LIST.txt NTLM .\Output.txt .\ntlm_hibp_compressed.bin
```

Output format : `{hash}:{occurence}`

## Installation :
Download the compiled version for Windows or Linux from [release page](https://github.com/JoshuaMart/PwnedPasswordsChecker/releases)

If you wish to compile it yourself, you will need to have golang installed on your system and perform the following commands:
```bash
git clone https://github.com/JoshuaMart/PwnedPasswordsChecker && cd PwnedPasswordsChecker
go build main.go
```

## Screenshots
Thanks to the use of a "compressed" format the tool has largely gained in performance, example of use between the old version and the new one with a list of 20,000 hashes (Intel Core I7 8565U) :

![Screenshot](https://zupimages.net/up/20/05/cudb.png)

## Improvements
Feel free to contact me on [Twitter](https://twitter.com/J0_mart) or do a PR to improve the script.
