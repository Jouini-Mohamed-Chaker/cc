package main

import (
	"bufio"
	"bytes"
	"io"
)

// Handles the "-clw" options when no option is provided
func GetStats(reader io.Reader) (*Stats, error) {
	var stats Stats

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		stats.Lines++

		stats.Bytes += len(scanner.Bytes()) + 1 // Adding 1 to account for the consumed newline character

		stats.Words += len(bytes.Fields(scanner.Bytes()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &stats, nil
}

// Handles the "-w" option
func CountNumberOfWords(reader io.Reader) (int, error) {
	var numberOfWords int

	scanner := bufio.NewScanner(reader)
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
func CountNumberOfLines(reader io.Reader) (int, error) {
	var numberOfLines int
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		numberOfLines++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return numberOfLines, nil
}

// Handles the "-c" option
func CountNumberOfBytes(reader io.Reader) (int, error) {
	var numberOfBytes int
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		numberOfBytes++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return numberOfBytes, nil
}
