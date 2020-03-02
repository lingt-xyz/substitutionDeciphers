package encrypt

import (
	"github.com/lingt-xyz/substitutionDeciphers/keySpace"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	t.Logf("Before shuffle: %v", keySpace.KeySpace)
	keys := GenerateKey(keySpace.KeySpace)
	t.Logf("After shuffle: %v", keys)
	t.Logf("KeySpace: %v", keySpace.KeySpace)
}

func TestEncipher(t *testing.T) {
	plainText := "defend the east wall of the castle"
	keys := []byte{'Q', 'L', 'D', 'N', 'M', 'G', 'A', 'F', 'H', 'J', 'X', 'I', 'Y', 'R', 'O', 'B', 'T', 'C', 'K', 'V', 'S', 'Z', 'W', 'P', 'U', 'E'}
	_, cipherText := Encipher(plainText, keys)
	expecting := "NMGMRNVFMMQKVWQIIOGVFMDQKVIM"
	if cipherText != expecting {
		t.Errorf("Expecting: %v, got: %v", expecting, cipherText)
	}
}
