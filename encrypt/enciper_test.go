package encrypt

import (
	"github.com/lingt-xyz/substitutionDeciphers/wordSpace"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	t.Logf("Before shuffle: %v", wordSpace.KeySpace)
	keys := GenerateKey(wordSpace.KeySpace)
	t.Logf("After shuffle: %v", keys)
	t.Logf("KeySpace: %v", wordSpace.KeySpace)
}
