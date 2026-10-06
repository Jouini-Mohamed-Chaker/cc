package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 3 {
		printUsage()
		os.Exit(1)
	}

	var option string
	var filename string

	if len(os.Args) == 1 {
		printUsage()
		os.Exit(1)

	} else if len(os.Args) == 2 {
		// Means we have one argument : can be either a filename
		// Or an option to pass stdin
		if strings.HasPrefix(os.Args[1], "-") {
			option := os.Args[1]

			printResultDependingOnOption(option, "", os.Stdin)

		} else {
			filename = os.Args[1]
			file, err := os.Open(filename)
			if err != nil {
				printErrorAndExit(err)
			}
			defer file.Close()

			// One function for the counting to reuse the buffer
			stats, err := GetStats(file)
			if err != nil {
				printErrorAndExit(err)
			}

			fmt.Printf("%d %d %d %s\n", stats.Bytes, stats.Lines, stats.Words, filename)
		}

	} else if len(os.Args) == 3 {
		// Means we have two arguments, probably an option and filename
		option = os.Args[1]
		filename = os.Args[2]

		file, err := os.Open(filename)
		if err != nil {
			printErrorAndExit(err)
		}
		defer file.Close()

		// We look what argument was provided
		printResultDependingOnOption(option, filename, file)
	}

}

func printResultDependingOnOption(option string, filename string, reader io.Reader) {
	switch option {
	case "-c":
		num, err := CountNumberOfBytes(reader)
		if err != nil {
			printErrorAndExit(err)
		}
		fmt.Printf("%d %s\n", num, filename)
	case "-l":
		num, err := CountNumberOfLines(reader)
		if err != nil {
			printErrorAndExit(err)
		}
		fmt.Printf("%d %s\n", num, filename)
	case "-w":
		num, err := CountNumberOfWords(reader)
		if err != nil {
			printErrorAndExit(err)
		}
		fmt.Printf("%d %s\n", num, filename)

	default:
		printUsage()
		os.Exit(1)
	}
}

func printErrorAndExit(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}

func printUsage() {
	fmt.Println("ccwc [-c|-l|-w] textfile")
}
