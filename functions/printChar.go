package functions

import (
	"fmt"
)

func PrintChar(str string, asciiRep [][]string) {
	arr := LineNum(str)
	for i := 0; i < 8; i++ {
		for _, num := range arr {
			fmt.Print(asciiRep[num][i])
		}
		fmt.Println("")
	}
}

func LineNum(str string) []int {
	var arr []int
	for _, char := range str {
		arr = append(arr, int(char)-32)
	}
	return arr
}
