package controller

import (
	"testing"
)

func Test_getWordFrequency(t *testing.T) {
	type args struct {
		lineList []string
	}
	simpleSentences := []string{
		"I went for lunch.",
		"I went, to the store, later.",
	}
	wordCounterForSimpleSentences := map[string]int{
		"i": 2, "went": 2, "for": 1, "lunch": 1, "to": 1, "the": 2, "store": 2, "later": 2,
	}

	manySpecialChars := []string{
		"Hey!! Your birthday @today?",
		"So - what did your 'family' give?",
	}
	wordCounterForSpecialChars := map[string]int{
		"hey": 1, "your": 2, "birthday": 1, "today": 1, "so": 1, "what": 1, "did": 1,
		"family": 1, "give": 1,
	}

	emptySentences := []string{
		"", "", "", "",
	}
	wordCounterForEmptySentences := map[string]int{}

	tables := []struct {
		sentences           []string
		expectedWordCounter map[string]int
	}{
		{simpleSentences, wordCounterForSimpleSentences},
		{manySpecialChars, wordCounterForSpecialChars},
		{emptySentences, wordCounterForEmptySentences},
	}

	for _, tt := range tables {
		wordsList := getWordFrequency(tt.sentences)
		for word := range tt.expectedWordCounter {
			if count, ok := wordsList[word]; ok {
				_ = count
			} else {
				t.Errorf("getWordFrequency() on %v does not contain %s", tt.sentences, word)
			}
		}
	}
}

func Test_getMostFrequentWords(t *testing.T) {
	equalSightings := map[string]int{
		"all": 3, "words": 3, "occur": 3, "equal": 3, "times": 3,
	}
	singleSightings := map[string]int{
		"all": 1, "words": 1, "appear": 1, "once": 1,
	}

	var found = 0
	wordFreqs := getMostFrequentWords(equalSightings, 8)
	for _, wf := range wordFreqs {
		if wf.Word == "all" {
			found = 1
		}
	}
	if found == 0 {
		t.Errorf("getWordFrequency() on %v - %v does not contain 'all'", equalSightings, wordFreqs)
	}

	sortedForSingleSightings := getMostFrequentWords(singleSightings, 10)
	for _, wf := range sortedForSingleSightings {
		if wf.Word == "once" {
			found = 1
		}
	}
	if found == 0 {
		t.Errorf("getWordFrequency() on %v - %v does not contain 'once'", singleSightings, sortedForSingleSightings)
	}
}
