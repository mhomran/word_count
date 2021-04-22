package main

import (
	"fmt"

	"github.com/mhomran/word_count/pkg/mapper"
	"github.com/mhomran/word_count/pkg/reducer"
	"github.com/mhomran/word_count/pkg/splitter"
)

const (
	INPUT_FILE  = "../../test/input/ExampleIn.txt"
	OUTPUT_FILE = "../../test/output/WordCountOutput.txt"
	THREADS_NUM = 5
)

func main() {
	var res bool

	//splitting
	SplittedData := splitter.Split(INPUT_FILE, THREADS_NUM)
	if SplittedData == nil {
		fmt.Println("[ERROR]\t Split function")
		return
	}

	//mapping
	WorkersOutput := mapper.Mapper(SplittedData)
	if WorkersOutput == nil {
		fmt.Println("[ERROR]\t Mapper function")
		return
	}

	//reducing
	res = reducer.Reducer(OUTPUT_FILE, WorkersOutput)
	if !res {
		fmt.Println("[ERROR]\t Reducer function")
		return
	}
}
