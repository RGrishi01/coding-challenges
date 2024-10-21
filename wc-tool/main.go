package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func countLines(s string) {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d %s\n", lineCount, s)
}

func countWords(s string) {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d %s\n", wordCount, s)
}

func countChars(s string) {
	data, err := os.ReadFile(s)

	if err != nil {
		log.Fatal(err)
	}

	charCount := utf8.RuneCount(data)

	fmt.Printf("%d %s\n", charCount, s)
}

func countBytes(s string) {
	fileInfo, err := os.Stat(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d %s\n", fileInfo.Size(), s)
}

func main() {
	fileBytes := flag.Bool("c", false, "")
	fileLines := flag.Bool("l", false, "")
	fileWords := flag.Bool("w", false, "")
	fileChars := flag.Bool("m", false, "")
	flag.Parse()

	fileName := flag.Arg(0)

	if fileName == "" {
		log.Fatal("Please provide a file name with appropriate flag")
	}

	s := fileName

	if !(*fileBytes || *fileChars || *fileLines || *fileWords) {
		countLines(s)
		countWords(s)
		countBytes(s)
	} else if *fileBytes {
		countBytes(s)
	} else if *fileLines {
		countLines(s)
	} else if *fileWords {
		countWords(s)
	} else if *fileChars {
		countChars(s)
	}
}
