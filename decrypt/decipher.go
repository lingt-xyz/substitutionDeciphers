package decrypt

import (
	"github.com/lingt-xyz/substitutionDeciphers/encrypt"
	"log"
	"math"
)

// http://norvig.com/mayzner.html
// https://gist.github.com/lydell/c439049abac2c9226e53

// guessKeyByFrequencyAnalysis gets keys by analysing letter frequencies
func guessKeyByFrequencyAnalysis(frequencyArray []LetterFrequency) []byte {
	keys := make([]byte, 26)
	for i, f := range frequencyArray {
		key := LetterFrequencyFactArray[i].letter
		keys[f.letter-'A'] = key
		//keys[i] = f.letter
	}
	return keys
}

func Decipher(s string, verbose bool) string {
	letterFrequencyArray := GetLetterFrequencies(s)
	putativeKey := guessKeyByFrequencyAnalysis(letterFrequencyArray)
	if verbose {
		log.Printf("Putative key: %v", string(putativeKey))
	}
	putativePlainText := encrypt.Encipher(s, putativeKey)
	key := fastMethodAlgorithm2(putativePlainText, putativeKey, verbose)
	return encrypt.Encipher(s, key)
}

// fastMethodAlgorithm1 has to parse the text every time
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
func fastMethodAlgorithm2(putativePlaintext string, key []byte, verbose bool) []byte {
	matrix := parseText(putativePlaintext)
	//matrix = ConvertAlphabetOrderToFrequencyOrder(matrix)
	if verbose {
		TabulateMatrix(matrix, false)
	}
	distance := getMatricesDistance(matrix, BiGramFactMatrixByAlphabet, /* BiGramFactMatrixByFrequency*/)
start:
	for i := 1; i < 26; i++ {
		for j := 0; j < 25-i; j++ {
			newMatrix := swapMatrix(matrix, j, j+i)
			newDistance := getMatricesDistance(newMatrix, BiGramFactMatrixByAlphabet, /* BiGramFactMatrixByFrequency*/)
			if newDistance < distance {
				// update keys and matrix
				key = swapKey(key, byte(j+'A'), byte(j+i+'A'))
				if verbose{
					log.Printf("Found lower error %v\n", newDistance)
					TabulateMatrix(newMatrix, false)
					log.Printf("Swap %v and %v", j, j+i)
					log.Printf("Putative key: %v", string(key))

				}
				matrix = newMatrix
				distance = newDistance
				// restart outer loop
				i = 0
				continue start
			}
		}
	}
	return key
}

// swapKey swaps keys at index `i` and `j`.
func swapKey(keys []byte, k1 byte, k2 byte) []byte {
	newKeys := make([]byte, 26)
	copy(newKeys, keys)
	indexK1 := -1
	indexK2 := -1
	for i, v := range keys {
		if k1 == v {
			indexK1 = i
		} else if k2 == v {
			indexK2 = i
		}
	}
	newKeys[indexK1] = keys[indexK2]
	newKeys[indexK2] = keys[indexK1]
	return newKeys
}

func swapMatrix(matrix [26][26]float64, j int, i int) [26][26]float64 {
	newMatrix := matrix
	newMatrix[j] = matrix[i]
	newMatrix[i] = matrix[j]
	for row := range newMatrix {
		newMatrix[row][j] = matrix[row][i]
		newMatrix[row][i] = matrix[row][j]
	}
	return newMatrix
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
