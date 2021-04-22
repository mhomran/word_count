package main

import (
	"fmt"

	"github.com/mhomran/word_count/pkg/mapper"
	"github.com/mhomran/word_count/pkg/reducer"
	"github.com/mhomran/word_count/pkg/splitter"
)

func main() {
	var res bool

	//splitting
	SplittedData := splitter.Split("../../test/input/ExampleIn.txt", 5)
	if SplittedData == nil {
		fmt.Println("[ERROR]\t Split function")
		return
	}

	//mapping
	WorkersOutput := mapper.Mapper(SplittedData)

	//reducing
	res = reducer.Reducer("../../test/output/WordCountOutput.txt", WorkersOutput)
	if !res {
		fmt.Println("[ERROR]\t Reducer function")
		return
	}
	/*
		p := make(tuple.TupleList, len(m))

		i := 0
		for k, v := range m {
			p[i] = tuple.Tuple{Key: k, Value: v}
			i++
		}

		sort.Sort(sort.Reverse(p))

		fileOut, err := os.Create("../../test/output/out.txt")
		if err != nil {
			panic(err)
		}

		for _, k := range p {
			str := k.Key + " : " + strconv.Itoa(k.Value) + " \n"
			fmt.Fprint(fileOut, str)
		}
	*/
}
