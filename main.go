package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type wordFound struct {
	Key   rune
	Count int
}

func main() {
	const bookPath = "books/frankenstein.txt"
	var bookContent string

	bookContent, err := getBookContent(bookPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	wordsCount := countWords(bookContent)
	letterCount := countLetters(bookContent)

	fmt.Printf("Total words: %v\n\n", wordsCount)
	displayLetterCount(letterCount)
	fmt.Println("--- End report ---")
}

func getBookContent(bookPath string) (string, error) {
	var scannedLines []string

	file, err := os.Open(bookPath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scannedLines = append(scannedLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}

	return strings.Join(scannedLines, "\n"), nil
}

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func countLetters(text string) map[rune]int {
	text = strings.ToLower(text)

	letterAppearances := make(map[rune]int)

	for _, letter := range text {
		if unicode.IsLetter(letter) {
			letterAppearances[letter]++
		}
	}
	return letterAppearances
}

func displayLetterCount(letterCount map[rune]int) {
	var wordsFound []wordFound

	for k, v := range letterCount {
		wordsFound = append(wordsFound, wordFound{k, v})
	}

	sort.Slice(wordsFound, func(i, j int) bool {
		return wordsFound[i].Count > wordsFound[j].Count
	})

	for _, wordFound := range wordsFound {
		fmt.Printf("The '%s' character was found %d times\n", string(wordFound.Key), wordFound.Count)
	}
}
