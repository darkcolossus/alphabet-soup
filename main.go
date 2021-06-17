package main

import (
	"fmt"
	"HomeTest/main/soup"
)

func main() {
	fmt.Println("Enter Rows Number: ")
	var f int
	fmt.Scanln(&f)

	var c int
	fmt.Println("Enter Cols Number: ")
	fmt.Scanln(&c)

	fmt.Println("Enter term to be searched: ")
	var term string
	fmt.Scanln(&term)

	ls := &soup.LetterSoup{}

	ls.Generate(f, c, term)
}


