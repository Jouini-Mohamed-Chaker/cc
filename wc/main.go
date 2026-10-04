package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var options = []string{"-c", "-l", "-w"}

func main() {
	if len(os.Args) > 3 {
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
	case "":
		if filename == "" {
			printUsage()
		os.Exit(1)
		}
		numberOfBytes, err := CountNumberOfBytes(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		numberOfLines, err := CountNumberOfLines(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		numberOfWords, err := CountNumberOfWords(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		fmt.Printf("%d %d %d %s\n", numberOfBytes, numberOfLines, numberOfWords, filename)
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Printf("ccwc [%s] textfile\n", strings.Join(options, "|"))
}
