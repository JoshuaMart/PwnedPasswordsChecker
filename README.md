# PwnedPasswordsChecker
![enter image description here](https://zupimages.net/up/20/04/zsp8.png)
**PwnedPasswordsChecker** is a tool that checks if the hash of a known password (in SHA1 or NTLM format) is present in the list of I Have Been Pwned leaks and the number of occurrences.

You can download the hash-coded version for SHA1 [here](https://downloads.pwnedpasswords.com/passwords/pwned-passwords-sha1-ordered-by-hash-v5.7z) or the hash-coded version for NTLM [here](https://downloads.pwnedpasswords.com/passwords/pwned-passwords-ntlm-ordered-by-hash-v5.7z)

*This script only works with the version sorted by hash and entry hashes must be in uppercase and preferably ordered by hashs*

## Usage :
`./PwnedPasswordsChecker ./inputHashList.txt ./OutputFile.txt ./pwned-passwords-{hash}-ordered-by-hash-v5.txt`

Output format : `{hash}:{occurence}`

## Installation :
Download the compiled version for Windows or Linux from [release page](https://github.com/JoshuaMart/PwnedPasswordsChecker/releases)

If you wish to compile it yourself, you will need to have golang installed on your system and perform the following commands:
```bash
git clone https://github.com/JoshuaMart/PwnedPasswordsChecker && cd PwnedPasswordsChecker
go get github.com/stoicperlman/fls
go build main.go
```

## Screenshots
Example of use on a list of more than 20,000 hashes (9sec) (Intel Core i7 8565U)
![enter image description here](https://zupimages.net/up/20/05/cudb.png)

## Improvements
Feel free to contact me on [Twitter](https://twitter.com/J0_mart) or do a PR to improve the script.
