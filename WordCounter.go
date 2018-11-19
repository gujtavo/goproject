package main

import (
	"fmt"
	"strings"
	"unicode"
	"io/ioutil"
)


func CountWords(text string) map[string]int {

	freq := make(map[string]int)
	for _, word := range strings.Fields(strings.ToLower(text)) {
		word = strings.Map(func(r rune) rune {
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

func main() {
	
	dataBin, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Print(err)
	}
	textFile := string(dataBin)
	
	count := CountWords(textFile)

	fmt.Printf("%+v", count)



	
	
	
}
