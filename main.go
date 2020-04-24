package main

import (
	"flag"
	"github.com/lingt-xyz/substitutionDeciphers/decrypt"
	"github.com/lingt-xyz/substitutionDeciphers/encrypt"
	"github.com/lingt-xyz/substitutionDeciphers/text"
	"io/ioutil"
	"log"
)

func main() {
	verbose := flag.Bool("v", false, "output verbose debug information")
	cipherType := flag.String("c", "encipher", "type of cipher operation; should be 'encipher', 'decipher' or 'demo'")

	givenText := flag.String("i", "", "full path of the file for testing")
	givenKey := flag.String("k", "", "key for encipher")
	flag.Parse()

	content, err := ioutil.ReadFile(*givenText)
	if err != nil {
		log.Fatalf("Cannot open the file '%v', please check its existence and the proper permission is given.", *givenText)
	}
	normalizedText := string(content)
	if *verbose {
		log.Printf("Given raw text: %v", normalizedText)
	}
	normalizedText = text.FilterText(normalizedText)

	if *cipherType == "encipher" {
		if *verbose {
			log.Printf("Plain text (%v letters): %v", len(normalizedText), normalizedText)
		}
		key := []byte(*givenKey)
		if len(key) == 0 {
			key = encrypt.GenerateKey(text.KeySpace)
			log.Printf("Key is not provided, using randomly generated key: %v", string(key))
		} else {
			log.Printf("Using given key: %v", string(key))
		}
		cipherText := encrypt.Encipher(normalizedText, key)
		log.Printf("Cipher text (length %v): %v", len(cipherText), cipherText)
	} else if *cipherType == "decipher" {
		if *verbose {
			log.Printf("Cipher text (%v letters): %v", len(normalizedText), normalizedText)
		}
		//decrypt.TabulateLetterFrequency(decrypt.GetLetterFrequencies(normalizedText))
		putativePlainText := decrypt.Decipher(normalizedText, *verbose)
		log.Printf("Plain text (length %v): %v", len(normalizedText), putativePlainText)
	} else if *cipherType == "demo" {
		if *verbose {
			log.Printf("Plain text (%v letters): %v", len(normalizedText), normalizedText)
		}
		//log.Printf("Fact letter frequency:")
		//decrypt.TabulateLetterFrequency(decrypt.LetterFrequencyFactArray)
		//log.Printf("Plain-text letter frequency:")
		//decrypt.TabulateLetterFrequency(decrypt.GetLetterFrequencies(normalizedText))
		key := encrypt.GenerateKey(text.KeySpace)
		log.Printf("Generated the encipher key: %v", string(key))
		//log.Printf("Expecting the decipher givenKey: %v", string(encrypt.InverseKey(key)))
		cipherText := encrypt.Encipher(normalizedText, key)
		log.Printf("Cipher text: %v", cipherText)
		//log.Printf("Cipher text letter frequency:")
		//decrypt.TabulateLetterFrequency(decrypt.GetLetterFrequencies(cipherText))
		putativePlainText := decrypt.Decipher(cipherText, *verbose)
		//log.Printf("Decipher text: %v", putativePlainText)
		numErr := 0
		for i, c := range putativePlainText {
			if byte(c) != normalizedText[i] {
				numErr++
			}
		}
		log.Printf("Plain text (length %v, error %.2f): %v", len(normalizedText), float64(numErr)/float64(len(normalizedText)), putativePlainText)
		//log.Printf("Plain text (length %v, error %.2f)", len(normalizedText), float64(numErr)/float64(len(normalizedText)))
	} else {
		log.Fatalf("Unsupported cipher operation %v; `encipher` or `decipher` was expected.", *cipherType)
	}
}