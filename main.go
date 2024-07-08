package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	a "ascii/ascii_art"
)

// Error messages to be displayed when the usage is incorrect
const ErrorText = `Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"`

const fsError = `Usage: go run . [STRING] [BANNER]

EX: go run . something standard`

const outError = `Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`

func main() {
	// Check if there are any arguments passed
	if len(os.Args) <= 1 {
		fmt.Println(fsError)
		os.Exit(0)
	}

	// Default arguments
	input := ""
	bannerFile := "standard"
	subString := input

	// Define the color flag

	var flgColor = flag.String("color", "", "Color")
	var flgOutput = flag.String("output", "", "output")
	flag.Parse()

	if *flgOutput != "" && !strings.HasSuffix(*flgOutput, ".txt") {
		fmt.Println("Error: the file must be a .txt file")
		os.Exit(0)
	}
	// Assuming flag.NFlag() correctly counts the number of flags set by the user
	nflags := flag.NFlag()

	if nflags == 1 {
		if !strings.Contains(os.Args[1], "--output=") && !strings.Contains(os.Args[1], "--color=") {
			switch nflags {
			case 1:
				fmt.Println(outError)
			case 2:
				fmt.Println(ErrorText)
			default:
				fmt.Println("Only two flags allowed")
			}
			os.Exit(0)
		}
	}

	color := a.ColorPicker(*flgColor)

	// Get non-flag arguments
	args := flag.Args() // Non-flag arguments
	nArgs := len(args)  // Count of non-flag arguments

	// Handle arguments based on the number of flags and non-flag arguments
	if nflags >= 1 {
		switch nArgs {
		case 1:
			input = args[0]
			subString = input
		case 2:
			input = args[0]
			subString = input
			bannerFile = args[1]
		case 3:
			subString = args[0]
			input = args[1]
			bannerFile = args[2]
		default:
			fmt.Println(ErrorText)
			os.Exit(0)

		}
	} else {
		switch nArgs {
		case 1:
			input = args[0]
		case 2:
			input = args[0]
			bannerFile = args[1]
		default:
			fmt.Println(fsError)
			os.Exit(0)
		}
	}

	if strings.Contains(bannerFile, ".") {
		fmt.Println(fsError)
		os.Exit(0)
	}

	bannerFile += ".txt"

	contents, err := a.GetFile(bannerFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Verify the length of the contents
	if len(contents) != 856 {
		fmt.Println("Error")
		return
	}

	// Process and print the input with the specified color and substring
	output := a.ProcessInput(contents, input, color, subString)

	if *flgOutput != "" {

		err = os.WriteFile(*flgOutput, []byte(output), 0644)
		if err != nil {
			fmt.Println("Erorr writing file", err)
			os.Exit(0)
		}
	} else {
		fmt.Print(output)
	}

}
