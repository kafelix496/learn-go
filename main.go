package main

import (
	"fmt"

	"github.com/kafelix496/learn-go/dictionary"
)

func main() {
	dict := dictionary.Dictionary{}

	addError := dict.Add("my", "word")
	if addError != nil {
		fmt.Println("add error", addError)
	}

	def, searchError := dict.Search("my")
	if searchError != nil {
		fmt.Println("search error 1", searchError)
	} else {
		fmt.Println("search result 1", def)
	}

	def2, searchError2 := dict.Search("my2")
	if searchError2 != nil {
		fmt.Println("search error 2", searchError2)
	} else {
		fmt.Println("search result 2", def2)
	}

	searchError3 := dict.Update("my", "new word")
	if searchError3 != nil {
		fmt.Println("search error 3", searchError)
	}

	searchError4 := dict.Update("my2", "new word")
	if searchError4 != nil {
		fmt.Println("search error 4", searchError4)
	}

	dict.Delete("my")

	def3, searchError5 := dict.Search("my")
	if searchError5 != nil {
		fmt.Println("search error 5", searchError5)
	} else {
		fmt.Println("search result 5", def3)
	}
}
