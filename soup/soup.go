package soup

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

type LetterSoup struct {
	F int
	C int
	M [][]string
	Term string
	Occurs int
}

// Create matrix.
func (ls *LetterSoup) CreateM() {
	ls.M = make([][]string, ls.F)

	for i := range ls.M {
		ls.M[i] = make([]string, ls.C)
	}
}

// Fill in matrix.
func (ls *LetterSoup) FillInM() {
	rand.Seed(time.Now().UnixNano())

	for i := range ls.M {
		for j := range ls.M[i] {
			ls.M[i][j] = PickUpLetter()
		}
		fmt.Printf("%v\n", ls.M[i])
	}
}

// Search term.
func (ls *LetterSoup) SearchTerm() {
	if ls.F >= 3 && ls.C >= 3 {
		for i := 0; i < ls.F; i++ {
			for j := 0; j < ls.C; j++ {
				ls.LR(i, j)
				ls.RL(i, j)
				ls.Up(i, j)
				ls.Down(i, j)
				ls.LRBD(i, j)
				ls.RLUD(i, j)
				ls.LRUD(i, j)
				ls.RLBD(i, j)
			}
		}
	} else if ls.F >= 3 && ls.C < 3 {
		for i := 0; i < ls.F; i++ {
			for j := 0; j < ls.C; j++ {
				ls.Up(i, j)
				ls.Down(i, j)
			}
		}
	} else {
		for i := 0; i < ls.F; i++ {
			for j := 0; j < ls.C; j++ {
				ls.LR(i, j)
				ls.RL(i, j)
			}
		}
	}
}

var tmp = []string{}

// Search term from left to right.
func (ls *LetterSoup) LR(i, j int) {
	if j+2 < ls.C {
		r:= ls.M[i][j:j+3]

		if strings.Join(r, "") == ls.Term {
			ls.Occurs++
		}
	}
}

// Search term from right to left.
func (ls *LetterSoup) RL(i, j int) {
	if j-2 >= 0 {
		// l = m[i][j:j-3]
		tmp = append(tmp, ls.M[i][j])
		tmp = append(tmp, ls.M[i][j-1])
		tmp = append(tmp, ls.M[i][j-2])

		if strings.Join(tmp, "") == ls.Term {
			ls.Occurs++
		}

		tmp = []string{}
	}
}

// Search up.
func (ls *LetterSoup) Up(i, j int) {
	if i-2 >= 0 {
		tmp = append(tmp, ls.M[i][j])
		tmp = append(tmp, ls.M[i-1][j])
		tmp = append(tmp, ls.M[i-2][j])

		if strings.Join(tmp, "") == ls.Term {
			ls.Occurs++
		}

		tmp = []string{}
	}
}

// Search down.
func (ls *LetterSoup) Down(i, j int) {
	if i+2 < ls.F {
		tmp = append(tmp, ls.M[i][j])
		tmp = append(tmp, ls.M[i+1][j])
		tmp = append(tmp, ls.M[i+2][j])

		if strings.Join(tmp, "") == ls.Term {
			ls.Occurs++
		}

		tmp = []string{}
	}
}

// Search from top left to bottom right diagonal.
func (ls *LetterSoup) LRBD(i, j int) {
	if i+2 < ls.F && j+2 < ls.C {
		tmp = append(tmp, ls.M[i][j])
		tmp = append(tmp, ls.M[i+1][j+1])
		tmp = append(tmp, ls.M[i+2][j+2])

		if strings.Join(tmp, "") == ls.Term {
			ls.Occurs++
		}

		tmp = []string{}
	}
}

// Search from bottom right to top left diagonal.
func (ls *LetterSoup) RLUD(i, j int) {
	if i-2 >= 0 && j-2 >= 0 {
		tmp = append(tmp, ls.M[i][j])
		tmp = append(tmp, ls.M[i-1][j-1])
		tmp = append(tmp, ls.M[i-2][j-2])

		if strings.Join(tmp, "") == ls.Term {
			ls.Occurs++
		}

		tmp = []string{}
	}
}

// Search from bottom left to top right diagonal.
func (ls *LetterSoup) LRUD(i, j int) {
	if i-2 >= 0 && j+2 < ls.C {
		tmp = append(tmp, ls.M[i][j])
		tmp = append(tmp, ls.M[i-1][j+1])
		tmp = append(tmp, ls.M[i-2][j+2])

		if strings.Join(tmp, "") == ls.Term {
			ls.Occurs++
		}

		tmp = []string{}
	}
}

// Search from top right to bottom left diagonal.
func (ls *LetterSoup) RLBD(i, j int) {
	if i+2 < ls.F && j-2 >= 0 {
		tmp = append(tmp, ls.M[i][j])
		tmp = append(tmp, ls.M[i+1][j-1])
		tmp = append(tmp, ls.M[i+2][j-2])

		if strings.Join(tmp, "") == ls.Term {
			ls.Occurs++
		}
		
		tmp = []string{}
	}
}

func (ls *LetterSoup) Generate(f int, c int, term string) int {

	if f < 1 || f > 100 || c < 1 || c > 100 {
		return -1
	} 

	if f < 3 && c < 3 {
		return 0
	}

	ls.F = f
	ls.C = c
	ls.Term = term
	ls.Occurs = 0

	ls.CreateM()
	ls.FillInM()
	ls.SearchTerm()
	fmt.Printf("Occurrences: %v\n", ls.Occurs)

	return ls.Occurs
}

// English Letter Frequency (based on a sample of 40,000 words).
var freqLetters = []rune { 'E', 'T', 'A', 'O', 'I', 'N', 'S', 'R', 'H',
		 'D', 'L', 'U', 'C', 'M', 'F', 'Y', 'W', 'G', 'P', 'B', 'V', 'K', 'Q', 'J', 'Z'}
const HalfHigherFreq = 0.834
const HalfHighestFreq = 0.5746


func PickUpLetter() string {
	return string(freqLetters[CalculateIndex()])
}

func CalculateIndex() int {
	rand.Seed(time.Now().UnixNano())
	randProb := rand.Float32()

	if randProb >= 0 && randProb <= HalfHigherFreq {
		randProbAccurate := rand.Float32()

		if randProbAccurate >= 0 && randProbAccurate <= HalfHighestFreq {
			return rand.Intn(6 - 0)
		} else {
			return rand.Intn(12 - 7) + 7
		}
	} else {
		return rand.Intn(25-13) + 13
	}	
}