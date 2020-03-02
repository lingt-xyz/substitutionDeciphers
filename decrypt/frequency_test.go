package decrypt

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

// convert the raw data of 2-letter counting to 2-letter bi-gram
func TestBiGram(t *testing.T) {
	matrix := [26][26]float64{}
	for _, v := range biGramArray {
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

// test the correctness of converting
func TestConverting(t *testing.T) {
	for _, v := range biGramArray {
		//t.Logf("%v", v)
		letters := strings.TrimSpace(v.letters)
		column := letters[0] - 'A'
		row := letters[1] - 'A'
		expecting := math.Round(v.frequency*1000000) / 10000
		got := BiGramMatrix[row][column]
		if got != expecting {
			t.Errorf("Matrix is not correct with the letters %v at %v:%v, expecting: %v, got: %v", letters, row, column, expecting, got)
		}
	}
}

// show matrix
func TestMatrix(t *testing.T) {
	fmt.Printf("|%6s", " ")
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("|%3s%3s", string(i), " ")
	}
	fmt.Printf("|\n")

	for index, r := range BiGramMatrix {
		fmt.Printf("|%3s%3s", string(index+'A'), " ")
		for _, v := range r {
			fmt.Printf("|%.4f", v)
		}
		fmt.Printf("|\n")
	}
}
