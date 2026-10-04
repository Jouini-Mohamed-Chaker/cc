package main

import (
	"bufio"
	"os"
)

// Handles the "-w" option
func CountNumberOfWords(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var numberOfWords int
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		numberOfWords++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return numberOfWords, nil
}

// Handles the "-l" option
func CountNumberOfLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var numberOfLines int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numberOfLines++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return numberOfLines, nil
}

// Handles the "-c" option
func CountNumberOfBytes(filename string) (int64, error) {
	stat, err := os.Stat(filename)
	if err != nil {
		return 0, nil
	}

	return stat.Size(), nil
}