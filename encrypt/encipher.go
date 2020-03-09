package encrypt

import (
	"math/rand"
	"strings"
	"time"
)

// GenerateKey generates a key randomly from the given key space.
func GenerateKey(keySpace []byte) []byte {
	keys := make([]byte, len(keySpace))
	copy(keys, keySpace)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })
	return keys
}

// InverseKey generates the decipher key from the given encipher key
func InverseKey(key []byte) []byte {
	inverseKey := make([]byte, len(key))
	for i, c := range key {
		inverseKey[c-'A'] = byte(i + 'A')
	}
	return inverseKey
}

// Encipher enciphers a given plain text with the given key and return the cipher.
func Encipher(s string, key []byte) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		b.WriteByte(key[c-'A'])
	}
	return b.String()
}
