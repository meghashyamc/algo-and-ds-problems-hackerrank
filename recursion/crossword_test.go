package recursion

import (
	"reflect"
	"testing"
)

var inputsCrossword = [][]string{{"+-++++++++",
	"+-++++++++",
	"+-++++++++",
	"+-----++++",
	"+-+++-++++",
	"+-+++-++++",
	"+++++-++++",
	"++------++",
	"+++++-++++",
	"+++++-++++",
	"LONDON;DELHI;ICELAND;ANKARA"},
	{"+-++++++++",
		"+-++++++++",
		"+-------++",
		"+-++++++++",
		"+-++++++++",
		"+------+++",
		"+-+++-++++",
		"+++++-++++",
		"+++++-++++",
		"++++++++++",
		"AGRA;NORWAY;ENGLAND;GWALIOR"},
	{"XXXXXX-XXX",
		"XX------XX",
		"XXXXXX-XXX",
		"XXXXXX-XXX",
		"XXX------X",
		"XXXXXX-X-X",
		"XXXXXX-X-X",
		"XXXXXXXX-X",
		"XXXXXXXX-X",
		"XXXXXXXX-X",
		"ICELAND;MEXICO;PANAMA;ALMATY"}}
var outputsCrossword = [][]string{{"+L++++++++",
	"+O++++++++",
	"+N++++++++",
	"+DELHI++++",
	"+O+++C++++",
	"+N+++E++++",
	"+++++L++++",
	"++ANKARA++",
	"+++++N++++",
	"+++++D++++"},
	{"+E++++++++",
		"+N++++++++",
		"+GWALIOR++",
		"+L++++++++",
		"+A++++++++",
		"+NORWAY+++",
		"+D+++G++++",
		"+++++R++++",
		"+++++A++++",
		"++++++++++"},
	{"XXXXXXIXXX",
		"XXMEXICOXX",
		"XXXXXXEXXX",
		"XXXXXXLXXX",
		"XXXPANAMAX",
		"XXXXXXNXLX",
		"XXXXXXDXMX",
		"XXXXXXXXAX",
		"XXXXXXXXTX",
		"XXXXXXXXYX"}}

func TestCrosswordPuzzle(t *testing.T) {

	for i, crossword := range inputsCrossword {

		res := crosswordPuzzle(crossword[:len(crossword)-1], crossword[len(crossword)-1])

		if !reflect.DeepEqual(res, outputsCrossword[i]) {
			t.Fatal("expected", outputsCrossword[i], "for", crossword, "got ", res)
		}
	}
}

func TestGetBlanksToFill(t *testing.T) {

	res := getBlanksToFill(getCrosswordBytes(inputsCrossword[0][:len(inputsCrossword[0])-1]))

	if len(res) != 4 {
		t.Fatal("expected 4 blanks, got", len(res))
	}
	if res[0].startx != 1 && res[0].endx != 6 {
		t.Fatal("expected first horizontal blank's coordinates to be 1,6 got", res[0].startx, res[0].endx)
	}
	if res[1].startx != 2 && res[1].endx != 8 {
		t.Fatal("expected second horizontal blank's coordinates to be 2,8 got", res[1].startx, res[1].endx)
	}
	if res[2].starty != 0 && res[2].endy != 6 {
		t.Fatal("expected first vertical blank's coordinates to be 0,6 got", res[2].starty, res[2].endy)
	}

}

func TestGetCrosswordBytes(t *testing.T) {

	crossword := inputsCrossword[0][:len(inputsCrossword[0])-1]
	res := getCrosswordBytes(crossword)

	if len(res) != len(crossword) {
		t.Fatal("expected length of crossword bytes to be", len(crossword), "got", len(res))
	}
	if len(crossword[0]) != len(res[0]) {
		t.Fatal("expected length of crossword bytes' first array to be", len(crossword[0]), "got", len(res[0]))

	}
}

func TestAddWordToBlank(t *testing.T) {

	crosswordBytes := getCrosswordBytes(inputsCrossword[0][:len(inputsCrossword[0])-1])
	blanks := getBlanksToFill(crosswordBytes)

	if !addWordToBlank(blanks[0], "DELHI", crosswordBytes) {
		t.Fatal("expected blank", blanks[0], "to be filled successfully with DELHI but addWordToBlank returned false")
	}

	if string(crosswordBytes[3]) != "+DELHI++++" {
		t.Fatal("expected blank to be filled successfully by DELHI, got", crosswordBytes[3])
	}

	if !addWordToBlank(blanks[2], "LONDON", crosswordBytes) {
		t.Fatal("expected blank", blanks[2], "to be filled successfully with LONDON but addWordToBlank returned false")
	}

	verticalBytes := []byte{crosswordBytes[0][1], crosswordBytes[1][1], crosswordBytes[2][1], crosswordBytes[3][1], crosswordBytes[4][1], crosswordBytes[5][1]}
	if string(verticalBytes) != "LONDON" {
		t.Fatal("expected blank to be filled successfully by LONDON, got", string(verticalBytes))
	}

}
