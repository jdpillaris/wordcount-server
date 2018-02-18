package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	var number = 10

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file path: ")
	file, _ := reader.ReadString('\n')
	// file := "/home/jdpillaris/Documents/Workspaces/Tests/Ulventech/servlet.txt"
	fmt.Println(file)

	var lines []string
	var wordCounter map[string]int
	var mostFrequentWords PairList

	lines = getLines(file)
	wordCounter = getWordFrequency(lines)
	mostFrequentWords = getMostFrequentWords(wordCounter, number)

	// fmt.Printf("%v", mostFrequentWords)
	fmt.Println("The", number, "most frequent words are:")
	for _, elem := range mostFrequentWords {
		fmt.Println("Word:", elem.Word, "; Count:", elem.Count)
	}
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
		words = strings.Fields(line)
		for _, word := range words {
			if _, exists := wordFrequencies[word]; exists {
				wordFrequencies[word]++
			} else {
				wordFrequencies[word] = 1
			}
		}
	}

	return wordFrequencies
}

func getMostFrequentWords(wordCounter map[string]int, number int) PairList {
	frequentWords := make(PairList, len(wordCounter))
	i := 0
	for k, v := range wordCounter {
		frequentWords[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(frequentWords))
	return frequentWords[:number]
}

// Pair struct
type Pair struct {
	Word  string
	Count int
}

// PairList : Array of Pairs
type PairList []Pair

func (pairs PairList) Len() int           { return len(pairs) }
func (pairs PairList) Less(i, j int) bool { return pairs[i].Count < pairs[j].Count }
func (pairs PairList) Swap(i, j int)      { pairs[i], pairs[j] = pairs[j], pairs[i] }
