# Methods for the Cryptanalysis of Substitution Ciphers

A Fast Method for the Cryptanalysis of Substitution Ciphers by Thomas Jakobsen

## Compile

- GNU/Linux: `go build`
- Windows: `env GOOS=windows GOARCH=amd64 go build`

## How to run it

> Use `substitutionDeciphers.exe` if you have a Windows system.

```shell script
./substitutionDeciphers -c=encipher -i=.test/t2 -v=true

./substitutionDeciphers -c=encipher -i=.test/t2 -k=TPULEMYNOKXDCAHFJSIRGBVQWZ -v=true

./substitutionDeciphers -c=decipher -i=.test/t2 -v=true

./substitutionDeciphers -c=demo -i=.test/t1 -v=true
```
