package functions

import (
	"fmt"
	"os"
)

func WriteToFile(outputFile string, text string, asciiRep [][]string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating file '%s': %v", outputFile, err)
	}
	defer file.Close()
	subStrs := SubStrs(text)
	len := len(subStrs)
	for index, subStr := range subStrs {
		if subStr == "\\n" {
			if index == len - 1 {
				_, err = fmt.Fprintln(file)
				if err != nil {
					return fmt.Errorf("error writing to file: %v", err)
				}
			} else if index == 0 {
				_, err = fmt.Fprintln(file)
				if err != nil {
					return fmt.Errorf("error writing to file: %v", err)
				} 
			} else {
				if subStrs[index - 1] == "\\n" {
					_, err = fmt.Fprint(file)
					if err != nil {
						return fmt.Errorf("error writing to file: %v", err)
					}
				}
			}
		} else {
			err = writeASCIIToFile(file, subStr, asciiRep)
			if err != nil {
				return fmt.Errorf("error writing to file: %v", err)
			}
		}
	}
	return nil
}


func writeASCIIToFile(file *os.File, str string, asciiRep [][]string) error {
	arr := LineNum(str)
	for i := 0; i < 8; i++ {
		line := ""
		for _, num := range arr {
			line += asciiRep[num][i]
			_, err := fmt.Fprint(file, line)
			if err != nil {
				return fmt.Errorf("error writing to file: %v", err)
			}
		}
		_,err := fmt.Fprintln(file)
		if err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}
	}
	return nil
}
