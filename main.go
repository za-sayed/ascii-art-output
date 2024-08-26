package main

import (
	"ascii-art/functions"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var text string
	var style string
	var outputFile string
	if len(os.Args) != 4 {
		if len(os.Args) == 2 {
			outputFile = ""
			text = os.Args[1]
			if !functions.ValidateASCII(text) {
				fmt.Println("Error: Non-ASCII character detected")
				return
			}
			style = "standard"
		} else if len(os.Args) == 3 {
			if len(os.Args[1]) > 9 && strings.ToLower(os.Args[1][:9]) == "--output=" {
				outputFile = os.Args[1][9:]
				if filepath.Ext(outputFile) != ".txt" {
					fmt.Println("Error: Output file must have a .txt extension")
					return
				}
			} else {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
				fmt.Println("Example: go run . --output=<fileName.txt> something standard")
				return
			}
			text = os.Args[2]
			if !functions.ValidateASCII(text) {
				fmt.Println("Error: Non-ASCII character detected")
				return
			}
			style = "standard"
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("Example: go run . --output=<fileName.txt> something standard")
			return	
		}
		
	} else {
		if len(os.Args[1]) > 9 && strings.ToLower(os.Args[1][:9]) == "--output=" {
			outputFile = os.Args[1][9:]
			if filepath.Ext(outputFile) != ".txt" {
				fmt.Println("Error: Output file must have a .txt extension")
				return
			}
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("Example: go run . --output=<fileName.txt> something standard")
			return
		}
		text = os.Args[2]
		if !functions.ValidateASCII(text) {
			fmt.Println("Error: Non-ASCII character detected")
			return
		}
		if strings.ToLower(os.Args[3]) == "shadow" || strings.ToLower(os.Args[3]) == "standard" || strings.ToLower(os.Args[3]) == "thinkertoy" {
			style = strings.ToLower(os.Args[3])
		} else {
			fmt.Println("Invalid style. Please use either shadow, standard, or thinkertoy")
			return
		}
	}
	fileLines := functions.Read(style)
	asciiRep := functions.AsciiRep(fileLines)
	if outputFile == "" {
		functions.PrintStr(text, asciiRep)
	} else {
		err := functions.WriteToFile(outputFile, text, asciiRep)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}
}
