package encrypt

import (
	"math/rand"
	"time"
)

func GenerateKey(keySpace []byte) []byte {
	keys := make([]byte, len(keySpace))
	copy(keys, keySpace)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })
	return keys
}

