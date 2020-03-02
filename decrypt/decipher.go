package decrypt

import (
	"github.com/lingt-xyz/substitutionDeciphers/text"
	"log"
	"math"
	"sort"
)

// http://norvig.com/mayzner.html
// https://gist.github.com/lydell/c439049abac2c9226e53

// getLetterFrequencies gets letter frequencies and sort letters by their frequencies
func getLetterFrequencies(s string) []LetterFrequency {
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

	return frequencyArray
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

func parseCipherText(s string) ([]byte, [26][26]float64) {
	letterFrequencyArray := getLetterFrequencies(s)
	keys := guessKeyByFrequencyAnalysis(letterFrequencyArray)
	biGramCountMatrix := [26][26]int{}
	biGramPercentageMatrix := [26][26]float64{}
	for i := 0; i < len(s)-1; i++ {
		c1, c2 := s[i], s[i+1]
		t1, t2 := keys[c1-'A'], keys[c2-'A']
		biGramCountMatrix[t2-'A'][t1-'A']++
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			f := float64(biGramCountMatrix[i][j]) / float64(len(s)-1)
			biGramPercentageMatrix[i][j] = math.Round(f*1000000) / 10000
		}
	}
	return keys, biGramPercentageMatrix
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

// 1. Construct an initial keySpace guess, `k`, based upon the symbol frequencies of the expected language and the ciphertext
// 2. Let `D=D(d(c,k))
// 3. Calculate `v=\sum(D-E))`
// 4. Let `k'=k`
// 5. Let `D'=D`
// 6. Change `k'` by swapping two elements, `a` and `b`, in `k'`
// 7. Exchange the corresponding rows in `D'`; exchange the corresponding columns in `D'`
// 8. Let `v'=\sum(D'-E)`
// 9. If `v'>v` then go to step 4
// 10. let `v=v'
// 11, Ket `k=k'`
// 12. Let `D=D'`
// 13. Go to step 6
func fastMethodAlgorithm2(input string) {
	cipherText := text.FilterText(input)
	keys, matrix := parseCipherText(cipherText)
	d := getMatricesDistance(matrix, BiGramFactMatrix)
	swapKeys(keys, 1)
	log.Printf("%v, %v", keys, matrix)
	log.Println(d)
}

func swapKeys(keys []byte, distance int) {
	newKeys := make([]byte, 26)
	for i := 0; i < len(keys)-distance; i++ {
		newKeys[i] = keys[i+distance]
	}
}

func getMatricesDistance(m1, m2 [26][26]float64) float64 {
	distance := 0.0
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			distance += math.Abs(m1[i][j] - m2[i][j])
		}
	}
	return distance
}
