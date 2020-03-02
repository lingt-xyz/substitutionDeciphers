package keySpace

import (
	"strings"
	"testing"
)

func TestFilterText(t *testing.T) {
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
		output := FilterText(testCase.input)
		if output != strings.ToUpper(testCase.output) {
			t.Errorf("Excepting: %v\n\t\tGot: %v", strings.ToUpper(testCase.output), output)
		}
	}
}
