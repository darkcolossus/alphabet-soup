package soup_test

import(
	"testing"
	"strings"
	"HomeTest/main/soup"
)

func TestGenerate(t *testing.T) {
	ls := &soup.LetterSoup{}
	occur := 0

	if occur = ls.Generate(0, 10, ""); occur != -1 {
		t.Error("Error when testing rows edge cases.")
		t.Fail()
	} else {
		t.Log("Edge case for rows tested correctly.")
	}

	if occur = ls.Generate(-1, 300, ""); occur != -1 {
		t.Error("Error when testing columns edge cases.")
		t.Fail()
	} else {
		t.Log("Edge case for columns tested correctly.")
	}

	if occur = ls.Generate(2, 1, ""); occur != 0 {
		t.Error("Error when testing little matrix.")
		t.Fail()
	} else {
		t.Log("Edge case for little matrix tested correctly.")
	}
}

func TestSearchTerms(t *testing.T) {
	m1 := [][]string{
		[]string{"O", "I", "E"},
		[]string{"I", "I", "X"},
		[]string{"E", "X", "E"},
	}

	f := 3
	c := 3
	searchTerm := "OIE"
	occurs := 0

	ls := &soup.LetterSoup{ F: f, C: c, M: m1, Term: searchTerm, Occurs: occurs}

	if ls.SearchTerm(); ls.Occurs != 3 {
		t.Error("Wrong amount of occurrences.")
		t.Fail()
	} else {
		t.Log("Correct amount of occurrences.")
	}

	m2 := [][]string{
		[]string{"E", "I", "O", "I", "E", "I", "O", "E", "I", "O"},
	}

	ls.F = 1
	ls.C = 10
	ls.M = m2
	ls.Occurs = 0

	if ls.SearchTerm(); ls.Occurs != 4 {
		t.Error("Wrong amount of occurrences.")
		t.Fail()
	} else {
		t.Log("Correct amount of occurrences.")
	}

	m3 := [][]string{
		[]string{"E", "A", "E", "A", "E"},
		[]string{"A", "I", "I", "I", "A"},
		[]string{"E", "I", "O", "I", "E"},
		[]string{"A", "I", "I", "I", "A"},
		[]string{"E", "A", "E", "A", "E"},
	}

	ls.F = 5
	ls.C = 5
	ls.M = m3
	ls.Occurs = 0

	if ls.SearchTerm(); ls.Occurs != 8 {
		t.Error("Wrong amount of occurrences.")
		t.Fail()
	} else {
		t.Log("Correct amount of occurrences.")
	}

	m4 := [][]string{
		[]string{"O", "X"},
		[]string{"I", "O"},
		[]string{"E", "X"},
		[]string{"I", "I"},
		[]string{"O", "X"},
		[]string{"I", "E"},
		[]string{"E", "X"},
	}

	ls.F = 7
	ls.C = 2
	ls.M = m4
	ls.Occurs = 0

	if ls.SearchTerm(); ls.Occurs != 3 {
		t.Error("Wrong amount of occurrences.")
		t.Fail()
	} else {
		t.Log("Correct amount of occurrences.")
	}

	m5 := [][]string{
		[]string{"E"},
	}

	ls.F = 1
	ls.C = 1
	ls.M = m5
	ls.Occurs = 0

	if ls.SearchTerm(); ls.Occurs != 0 {
		t.Error("Wrong amount of occurrences.")
		t.Fail()
	} else {
		t.Log("Correct amount of occurrences.")
	}
}

func TestPickUpLetter(t *testing.T) {
	l := soup.PickUpLetter()
	
	if 	!strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZ", l) {
		t.Error("Generated character is not in the alphabet.")
		t.Fail()
	} else {
		t.Log("Generated character is in the alphabet.")
	}
}

func TestCalculateIndex(t *testing.T) {
	r1 := 0
	
	for i := 0; i < 10; i++ {
		r1 = soup.CalculateIndex()

		if r1 < 0 || r1 > 25 {
			t.Error("Random number is out of range.")
			t.Fail()
		} else {
			t.Log("Random number is in the range.")
		}
	}
}