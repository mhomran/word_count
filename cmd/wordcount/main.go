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
	SplittedData := splitter.Split("../../ExampleIn.txt")

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

	oldVal := m[" "]
	m[" "] = oldVal - 1

	fileOut, err := os.Create("../../out.txt")
	if err != nil {
		panic(err)
	}

	for _, k := range p {
		str := k.Key + " : " + strconv.Itoa(k.Value) + " \n"
		fmt.Fprint(fileOut, str)
	}
}
