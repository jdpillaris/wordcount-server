package controller

import (
	"bufio"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"unicode"
)

// WordFrequency struct
type WordFrequency struct {
	Word  string
	Count int
}

// WordFrequencyList : Array of WordFrequency
type WordFrequencyList []WordFrequency

func (pairs WordFrequencyList) Len() int           { return len(pairs) }
func (pairs WordFrequencyList) Less(i, j int) bool { return pairs[i].Count < pairs[j].Count }
func (pairs WordFrequencyList) Swap(i, j int)      { pairs[i], pairs[j] = pairs[j], pairs[i] }

func Upload(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{}
	t, _ := template.ParseFiles("assets/upload.html")
	if r.Method == "GET" {
		t.Execute(w, nil)
	} else {
		var wordsToDisplay = 10
		filePath := storeFile(r)

		var lines = getLines(filePath)
		var wordCounter = getWordFrequency(lines)
		var mostFrequentWords = getMostFrequentWords(wordCounter, wordsToDisplay)

		m["myList"] = mostFrequentWords
		t.Execute(w, m)
	}
}

func storeFile(r *http.Request) string {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f, err := os.OpenFile("./.store/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(f, file)

	filePath := "./.store/" + handler.Filename
	return filePath
}

func getLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func getWordFrequency(lineList []string) map[string]int {
	var wordFrequencies = make(map[string]int)
	var words []string

	for _, line := range lineList {
		isNotAlphaNumeric := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		words = strings.FieldsFunc(line, isNotAlphaNumeric)

		for _, word := range words {
			word = strings.ToLower(word)
			if _, exists := wordFrequencies[word]; exists {
				wordFrequencies[word]++
			} else {
				wordFrequencies[word] = 1
			}
		}
	}

	return wordFrequencies
}

func getMostFrequentWords(wordCounter map[string]int, numberOfWords int) WordFrequencyList {
	frequentWords := make(WordFrequencyList, len(wordCounter))
	i := 0
	for k, v := range wordCounter {
		frequentWords[i] = WordFrequency{k, v}
		i++
	}
	sort.Sort(sort.Reverse(frequentWords))
	if numberOfWords > cap(frequentWords) {
		return frequentWords[:]
	}
	return frequentWords[:numberOfWords]

}
