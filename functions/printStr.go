package functions

import (
	"fmt"
	"strings"
)

func PrintStr(str string, asciiRep [][]string) {
	subStrs := SubStrs(str)
	len := len(subStrs)
	for index, subStr := range subStrs {
		if subStr == "\\n" {
			if index == len - 1 {
				fmt.Print("\n")
			} else if index == 0 {
				fmt.Print("\n") 
			} else {
				if subStrs[index - 1] == "\\n" {
					fmt.Print("\n")
				}
			}
		} else {
			PrintChar(subStr, asciiRep)
		}
	}
}

func SubStrs(str string) []string {
	parts := strings.Split(str, "\\n")
	var result []string
	for i, part := range parts {
		if part != "" {
			result = append(result, part)
		}
		if i < len(parts) - 1 {
			result = append(result, "\\n")
		}
	}
	return result
}


