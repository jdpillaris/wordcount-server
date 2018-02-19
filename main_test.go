package main

import (
	"reflect"
	"testing"
)

func Test_getLines(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLines(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWordFrequency(t *testing.T) {
	type args struct {
		lineList []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWordFrequency(tt.args.lineList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getWordFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMostFrequentWords(t *testing.T) {
	type args struct {
		wordCounter map[string]int
		number      int
	}
	tests := []struct {
		name string
		args args
		want PairList
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMostFrequentWords(tt.args.wordCounter, tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMostFrequentWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_main(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			main()
// 		})
// 	}
// }

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
