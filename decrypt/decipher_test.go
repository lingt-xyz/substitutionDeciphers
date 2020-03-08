package decrypt

import (
	"fmt"
	"github.com/lingt-xyz/substitutionDeciphers/encrypt"
	"github.com/lingt-xyz/substitutionDeciphers/text"
	"io/ioutil"
	"strings"
	"testing"
)

func TestGetLetterFrequencies(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"const nihongo = \"日本語（にほんご、にっぽんご）\"", "constnihongo"},
		{"fmt.Printf(\"%#U starts at byte position %d\n\", runeValue, index)", "fmtPrintfUstartsatbytepositiondruneValueindex"},
		{"[Exercise: Put an invalid UTF-8 byte sequence into the string. (How?) What happens to the iterations of the loop?]", "ExercisePutaninvalidUTFbytesequenceintothestringHowWhathappenstotheiterationsoftheloop"},
		{"Le français est une langue indo-européenne de la famille des langues romanes. Le français s'est formé en France (variété de la « langue d'oïl », qui est la langue de la partie septentrionale du pays). Le français est déclaré langue officielle en France en 15394.", "LefranaisestunelangueindoeuropennedelafamilledeslanguesromanesLefranaissestformenFrancevaritdelalanguedolquiestlalanguedelapartieseptentrionaledupaysLefranaisestdclarlangueofficielleenFranceen"},
		{`U+65E5 '日' starts at byte position 0
			U+672C '本' starts at byte position 3
			U+8A9E '語' starts at byte position 61`, "UEstartsatbytepositionUCstartsatbytepositionUAEstartsatbyteposition"},
	}
	for _, testCase := range testCases {
		frequencyArray := getLetterFrequencies(strings.ToUpper(testCase.output))
		t.Logf("Frequency: %v", frequencyArray)
	}
}

func TestGuessKeyByFrequencyAnalysis(t *testing.T){
	cipherText := strings.ToUpper("LefranaisestunelangueindoeuropennedelafamilledeslanguesromanesLefranaissestformenFrancevaritdelalanguedolquiestlalanguedelapartieseptentrionaledupaysLefranaisestdclarlangueofficielleenFranceen")
	frequencyArray := getLetterFrequencies(cipherText)
	keys := guessKeyByFrequencyAnalysis(frequencyArray)
	t.Logf("Guessed keys: %v", keys)
}

func TestParseCipherText(t *testing.T){
	plainText := text.FilterText("defend the east wall of the castle")
	encrypt.Encipher(plainText, encrypt.GenerateKey(text.KeySpace))
	matrix := parseText(plainText)
	t.Logf("%v", matrix)
}

func TestParsePlainText(t *testing.T){
	testFile := "../data/sample.txt"
	content, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Cannot open the file '%v', please check its existence and the reading permission is given.", testFile)
	}
	plainText := text.FilterText(string(content))

	matrix := parseText(plainText)
	fmt.Print("{")
	for _, r := range matrix{
		fmt.Print("{")
		for _, c := range r{
			fmt.Print(c)
			fmt.Printf(",")
		}
		fmt.Print("}")
		fmt.Printf(",")
	}
	fmt.Print("}")
	//t.Logf("%v", matrix)
}