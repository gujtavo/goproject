package main

import (
	"fmt"
	"strings"
	"unicode"
	"io/ioutil"
	"sort"
)


func CountWords(text string) map[string]int {

	freq := make(map[string]int)
	for _, word := range strings.Fields(strings.ToLower(text)) {
		word = strings.Map(func(r rune) rune { //Rune literals are just 32-bit integer values. (however they're untyped constants, so their type can change)
			if unicode.IsLetter(r) {
				return r
			} else {
				return -1
			}
		}, word)
		if word != "" {
			freq[word] = freq[word] + 1
		}
		
	}
	
	return freq

}

func order(Map map[string]int) []string {
	var words []string
	for word := range Map {
		words = append(words, word)
	}
	sort.Strings(words)
	return words
}

func main() {
	
	data, err := ioutil.ReadFile("words.txt")
	if err != nil { // A successful call
		fmt.Print(err)
	}
	textFile := string(data)
	
	countedWords := CountWords(textFile)
	orderedCountList := order(countedWords)

	for _, word := range orderedCountList {
		fmt.Println(word, countedWords[word])
	}

	



	
	
	
}
