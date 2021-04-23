package reducer

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/mhomran/word_count/pkg/tuple"
)

//Takes a map and sorts it and write it in a file with a name
//FileName
func Reducer(FileName string, MapInput map[string]int) bool {
	var file *os.File
	var ferr error
	var i int

	if MapInput == nil {
		return false
	}

	Tuples := make(tuple.TupleList, len(MapInput))

	i = 0
	for k, v := range MapInput {
		Tuples[i] = tuple.Tuple{Key: k, Value: v}
		i++
	}

	//sort.Reverse: just to sort the list in reverse order.
	sort.Sort(sort.Reverse(Tuples))

	file, ferr = os.Create(FileName)
	if ferr != nil {
		return false
	}

	for _, k := range Tuples {
		str := k.Key + " : " + strconv.Itoa(k.Value) + " \n"
		fmt.Fprint(file, str)
	}

	ferr = file.Close()
	if ferr != nil {
		return ferr == nil
	}

	return true
}
