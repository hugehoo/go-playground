package main

import (
	"fmt"
	"nomad-go/chapter2/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	dictionary[word] = word
	fmt.Println(dictionary["first"])

	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	// Add section
	def := "Greeting"
	er := dictionary.Add(word, def)
	if er != nil {
		fmt.Println(er)
	}
	definition, err2 := dictionary.Search(word)
	fmt.Println(definition)
	if err2 != nil {
		fmt.Println(err2)
	}

	// update Section
	err3 := dictionary.Update(word, "Second")
	if err3 != nil {
		fmt.Println(err3)
	}
	foundWord, _ := dictionary.Search(word)
	fmt.Println(foundWord)

	// Delete section
	baseWord := "Running"
	dictionary.Add(baseWord, "is what i like")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	searchWord, err4 := dictionary.Search(baseWord)
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(searchWord)
}
