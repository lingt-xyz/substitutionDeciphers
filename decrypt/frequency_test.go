package decrypt

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestLetterFrequencies(t *testing.T) {
	if len(LetterFrequencyFactArray) != 26 {
		t.Errorf("LetterFrequencyFactArray should be size of 26, got %v", len(LetterFrequencyFactArray))
	}
	for i := 1; i < len(LetterFrequencyFactArray); i++ {
		if LetterFrequencyFactArray[i].frequency >= LetterFrequencyFactArray[i-1].frequency {
			t.Errorf("LetterFrequencyFactArray is not sorted between %v and %v", LetterFrequencyFactArray[i-1], LetterFrequencyFactArray[i])
		}
	}
}

// convert the raw data of 2-letter counting to 2-letter bi-gram
func TestBiGram(t *testing.T) {
	matrix := [26][26]float64{}
	for _, v := range biGramFactArray {
		//t.Logf("%v", v)
		letters := strings.TrimSpace(v.letters)
		column := letters[0] - 'A'
		row := letters[1] - 'A'
		matrix[row][column] = v.frequency
	}
	//t.Logf("%v", matrix)

	for _, m := range matrix {
		fmt.Printf("%v", "{")
		for i, f := range m {
			if i == len(m)-1 {
				fmt.Printf("%v", math.Round(f*1000000)/10000)
			} else {
				fmt.Printf("%v,", math.Round(f*1000000)/10000)
			}
		}
		fmt.Printf("%v\n", "},")
	}
}

func TestConvertAlphabetOrderToFrequencyOrder(t *testing.T){
	matrix := ConvertAlphabetOrderToFrequencyOrder(BiGramFactMatrixByAlphabet)
	for _, r := range matrix{
		fmt.Print("{")
		for j, c := range r{
			fmt.Print(c)
			if j!= len(r)-1{
				fmt.Printf(",")
			}
		}
		fmt.Print("}")
		fmt.Printf(",\n")
	}
}

func TestTabulateMatrix(t *testing.T) {
	TabulateMatrix(BiGramFactMatrixByAlphabet, false)
	TabulateMatrix(BiGramFactMatrixByFrequency, true)
}