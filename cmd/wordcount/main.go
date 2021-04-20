package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/mhomran/word_count/pkg/splitter"
	"github.com/mhomran/word_count/pkg/tuple"
)

var m map[string]int

func main() {
	//splitting
	SplittedData := splitter.Split("../../test/input/ExampleIn.txt")
	if SplittedData == nil {
		fmt.Println("[ERROR]\t Split function")
		return
	}

	//mapping
	//reducing

	m = make(map[string]int)
	//occurance
	for i := 0; i < len(SplittedData); i++ {
		m[SplittedData[i]] += 1
	}

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
}
