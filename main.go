package main

import (
	"flag"
	"github.com/lingt-xyz/substitutionDeciphers/decrypt"
	"github.com/lingt-xyz/substitutionDeciphers/encrypt"
	"github.com/lingt-xyz/substitutionDeciphers/text"
	"io/ioutil"
	"log"
	"runtime"
)

func main() {
	log.Printf("Number of CPUs: %v\n", runtime.NumCPU())

	testFile := flag.String("input", "test", "full path of the file for testing")
	flag.Parse()

	content, err := ioutil.ReadFile(*testFile)
	if err != nil {
		log.Panicf("Cannot open the file '%v', please check its existence and the reading permission is given.", *testFile)
	}
	plainText := string(content)
	//log.Printf("Raw text: %v", plainText)
	plainText = text.FilterText(plainText)
	log.Printf("Plain text (%v): %v", len(plainText), plainText)
	key := encrypt.GenerateKey(text.KeySpace)
	log.Printf("Generated the key: %v", string(key))
	log.Printf("Expecting the key: %v", string(encrypt.InverseKey(key)))
	cipherText := encrypt.Encipher(plainText, key)
	//log.Printf("Cipher text: %v", cipherText)
	plainText = decrypt.Decipher(cipherText)
	log.Printf("Plain text: %v", plainText)
}