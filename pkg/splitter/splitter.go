package splitter

import (
	"bufio"
	"math"
	"os"
	"strings"
)

// Takes an input text file and outputs a string slice containing all
// the splitted strings inside the text file
func Split(FileName string, SlicesNumber int) [][]string {
	var SplittedArray []string
	var SlicedPart []string
	var SplittedOutput [][]string
	var ferr error
	var file *os.File
	var i int
	var step int

	if SlicesNumber == 0 {
		return nil
	}

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

	step = len(SplittedArray) / SlicesNumber
	for i := 0; i < (step * SlicesNumber); i += step {
		SlicedPart = SplittedArray[i:int(math.Min(float64(i+step), float64(len(SplittedArray))))]
		SplittedOutput = append(SplittedOutput, SlicedPart)
	}

	SplittedOutput[SlicesNumber-1] = SplittedArray[step*(SlicesNumber-1):]

	ferr = file.Close()
	if ferr != nil {
		return nil
	}

	return SplittedOutput
}
