package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/mhomran/word_count/pkg/tuple"
)

var m map[string]int

func main() {
	file, ferr := os.Open("../../ExampleIn.txt")
	arr := []string{""}
	m = make(map[string]int)

	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		for i := 0; i < len(items); i++ {
			arr = append(arr, strings.ToLower(items[i]))
		}
	}

	//occurance
	for i := 0; i < len(arr); i++ {
		m[arr[i]] += 1
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

	fileOut, err := os.Create("out.txt")
	if err != nil {
		panic(err)
	}

	for _, k := range p {
		str := k.Key + " : " + strconv.Itoa(k.Value) + " \n"
		fmt.Fprint(fileOut, str)
	}
}
