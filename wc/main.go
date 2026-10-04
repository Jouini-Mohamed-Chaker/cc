package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var options = []string{"-c", "-l", "-w", "-m"}

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		printUsage()
		os.Exit(1)
	}

	var option string
	var filename string

	// If first argument after command name is an option
	if slices.Contains(options, os.Args[1]) {
		// There has to be an argument after the option which is the filename
		if len(os.Args) != 3 {
			printUsage()
			os.Exit(1)
		}
		option = os.Args[1]
		filename = os.Args[2]
	} else {
		filename = os.Args[1]
	}

	switch option {
	case "-c":
		num, err := CountNumberOfBytes(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		fmt.Printf("%d %s\n", num, filename)
	case "-l":
		num, err := CountNumberOfLines(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		fmt.Printf("%d %s\n", num, filename)
	case "-w":
		num, err := CountNumberOfWords(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		fmt.Printf("%d %s\n", num, filename)
	case "-m":
		num, err := CountNumberOfCharacters(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		fmt.Printf("%d %s\n", num, filename)
	default:
		printUsage()
		os.Exit(1)
	}
}

// Handles the "-m" option
func CountNumberOfCharacters(filename string) (int, error) {
	panic("unimplemented")
}

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

	fmt.Printf("%d %s\n", numberOfLines, filename)
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

func printUsage() {
	fmt.Printf("ccwc [%s] textfile\n", strings.Join(options, "|"))
}
