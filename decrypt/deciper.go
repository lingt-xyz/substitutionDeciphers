package decrypt

import "strings"

// http://norvig.com/mayzner.html
// https://gist.github.com/lydell/c439049abac2c9226e53

func getSymbolFrequency(s string) (string, map[byte]int) {
	symbolFrequency := make(map[byte]int, 'z'-'a'+1)
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'a' <= c && c <= 'z' {
			c -= 'a' - 'A'
			b.WriteByte(c)
			symbolFrequency[c]++
		} else if 'A' <= c && c <= 'Z' {
			b.WriteByte(c)
			symbolFrequency[c]++
		}
	}
	return b.String(), symbolFrequency
}

// have to parse the text every time
func fastMethodAlgorithm1() {
	// 1. Construct an initial wordSpace guess, `k`, based upon the symbol frequencies of the expected language and the ciphertext
	// 2. Calculate `v=f(d(letter,k))`
	// 3. Let `k'=k`
	// 4. Change `k'` by swapping two elements, `a` and `b`, in `k'`
	// 5. Let `v'=f(d(letter,k'))`
	// 6. If `v'<v` then let `v=v'` and let `k=k'`
	// 7. Go to step 3
}

func fastMethodAlgorithm2() {
	// 1. Construct an initial wordSpace guess, `k`, based upon the symbol frequencies of the expected language and the ciphertext
	// 2. Let `D=D(d)
	// 2. Calculate `v=f(d(letter,k))`
	// 3. Let `k'=k`
	// 4. Change `k'` by swapping two elements, `a` and `b`, in `k'`
	// 5. Let `v'=f(d(letter,k'))`
	// 6. If `v'<v` then let `v=v'` and let `k=k'`
	// 7. Go to step 3
}

func parseText(s string){
	symbolFrequency := make(map[byte]int, 'z'-'a'+1)
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s)-1; i++ {
		c := s[i]
		if 'a' <= c && c <= 'z' {
			c -= 'a' - 'A'
			b.WriteByte(c)
			symbolFrequency[c]++
		} else if 'A' <= c && c <= 'Z' {
			b.WriteByte(c)
			symbolFrequency[c]++
		}
	}
}