package encrypt

import (
	"github.com/lingt-xyz/substitutionDeciphers/text"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	t.Logf("Before shuffle: %v", string(text.KeySpace))
	keys := GenerateKey(text.KeySpace)
	t.Logf("After shuffle: %v", string(keys))
	if string(keys) == string(text.KeySpace){
		t.Errorf("Nothing changed.")
	}
}

func TestEncipher(t *testing.T) {
	plainText := "defend the east wall of the castle"
	plainText = text.FilterText(plainText)
	keys := []byte{'Q', 'L', 'D', 'N', 'M', 'G', 'A', 'F', 'H', 'J', 'X', 'I', 'Y', 'R', 'O', 'B', 'T', 'C', 'K', 'V', 'S', 'Z', 'W', 'P', 'U', 'E'}
	cipherText := Encipher(plainText, keys)
	expecting := "NMGMRNVFMMQKVWQIIOGVFMDQKVIM"
	if cipherText != expecting {
		t.Errorf("Expecting: %v, got: %v", expecting, cipherText)
	}
}
