package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/mhomran/word_count/pkg/mapper"
	"github.com/mhomran/word_count/pkg/splitter"
	"github.com/mhomran/word_count/pkg/tuple"
)

func main() {
	//splitting
	SplittedData := splitter.Split("../../test/input/ExampleIn.txt", 5)
	if SplittedData == nil {
		fmt.Println("[ERROR]\t Split function")
		return
	}

	//mapping
	m := mapper.Mapper(SplittedData)

	//reducing
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
