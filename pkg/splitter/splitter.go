package splitter

import (
	"bufio"
	"os"
	"strings"
)

func Split(FileName string) []string {
	var SplittedArray []string
	var ferr error
	var file *os.File
	var i int

	file, ferr = os.Open(FileName)
	if ferr != nil {
		return nil
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		for i = 0; i < len(items); i++ {
			SplittedArray = append(SplittedArray, strings.ToLower(items[i]))
		}
	}

	ferr = file.Close()
	if ferr != nil {
		return nil
	}

	return SplittedArray
}
