package decrypt

import (
	"github.com/lingt-xyz/substitutionDeciphers/keySpace"
	"sort"
)

// http://norvig.com/mayzner.html
// https://gist.github.com/lydell/c439049abac2c9226e53

// getLetterFrequencies gets letter frequencies and sort letters by their frequencies
func getLetterFrequencies(s string) (string, []LetterFrequency) {
	s = keySpace.FilterText(s)
	frequencyArray := make([]LetterFrequency, 26)
	// initialize letters
	for i := 0; i < len(frequencyArray); i++ {
		frequencyArray[i].letter = byte(i + 'A')
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		frequencyArray[c-'A'].frequency++
	}

	// sort by frequency
	sort.SliceStable(frequencyArray, func(i, j int) bool {
		return frequencyArray[i].frequency > frequencyArray[j].frequency
	})

	return s, frequencyArray
}

// guessKeyByFrequencyAnalysis gets keys by analysing letter frequencies
func guessKeyByFrequencyAnalysis(frequencyArray []LetterFrequency) []byte {
	keys := make([]byte, 26)
	for i, f := range frequencyArray {
		key := LetterFrequencyFactArray[i].letter
		keys[f.letter-'A'] = key
	}
	return keys
}

func parseCipherText(s string) ([]byte, [26][26]int) {
	cipherText, letterFrequencyArray := getLetterFrequencies(s)
	keys := guessKeyByFrequencyAnalysis(letterFrequencyArray)
	biGramMatrix := [26][26]int{}
	for i := 0; i < len(cipherText)-1; i++ {
		c1, c2 := s[i], s[i+1]
		t1, t2 := keys[c1-'A'], keys[c2-'A']
		biGramMatrix[t2-'A'][t1-'A']++
	}
	return keys, biGramMatrix
}

// have to parse the text every time
func fastMethodAlgorithm1() {
	// 1. Construct an initial keySpace guess, `k`, based upon the symbol frequencies of the expected language and the ciphertext
	// 2. Calculate `v=f(d(letter,k))`
	// 3. Let `k'=k`
	// 4. Change `k'` by swapping two elements, `a` and `b`, in `k'`
	// 5. Let `v'=f(d(letter,k'))`
	// 6. If `v'<v` then let `v=v'` and let `k=k'`
	// 7. Go to step 3
}

func fastMethodAlgorithm2() {
	// 1. Construct an initial keySpace guess, `k`, based upon the symbol frequencies of the expected language and the ciphertext
	// 2. Let `D=D(d)
	// 2. Calculate `v=f(d(letter,k))`
	// 3. Let `k'=k`
	// 4. Change `k'` by swapping two elements, `a` and `b`, in `k'`
	// 5. Let `v'=f(d(letter,k'))`
	// 6. If `v'<v` then let `v=v'` and let `k=k'`
	// 7. Go to step 3
}
