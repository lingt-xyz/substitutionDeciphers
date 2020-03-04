package encrypt

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateKey(keySpace []byte) []byte {
	keys := make([]byte, len(keySpace))
	copy(keys, keySpace)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })
	return keys
}

func InverseKey(key []byte) []byte {
	inverseKey := make([]byte, len(key))
	for i, c := range key {
		inverseKey[c-'A'] = byte(i + 'A')
	}
	return inverseKey
}

func Encipher(s string, keys []byte) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		b.WriteByte(keys[c-'A'])
	}
	return b.String()
}
