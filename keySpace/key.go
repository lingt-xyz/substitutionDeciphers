package keySpace

import "strings"

var KeySpace = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

// FilterText removes all letters that are out of key space
func FilterText(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		// remove letters that out of key space
		if 'a' <= c && c <= 'z' {
			c -= 'a' - 'A'
			b.WriteByte(c)
		} else if 'A' <= c && c <= 'Z' {
			b.WriteByte(c)
		}
	}
	return b.String()
}