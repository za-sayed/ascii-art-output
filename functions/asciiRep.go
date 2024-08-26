package functions

func AsciiRep(fileLines []string) [][]string {
	var asciiRep [][]string
	var arr []string
	counter := 0
	for _, line := range fileLines {
		if counter == 0 {
			counter++
			continue
		}
		arr = append(arr, line)
		counter++
		if counter == 9 {
			asciiRep = append(asciiRep, arr)
			arr = nil
			counter = 0
		}
	}
	return asciiRep
}
