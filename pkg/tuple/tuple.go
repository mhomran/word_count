package tuple

import "sort"

//A data structure representing and tuple which takes a string (key)
// and an integer (value).
type Tuple struct {
	Key   string
	Value int
}

type TupleList []Tuple

//return the length of a list of tuples.
func (p TupleList) Len() int { return len(p) }

//swap two elements inside a list of tuples.
func (p TupleList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//comparison function to compare between two tuples.
//Returns true if the first element's value is less than the second's.
//In case of two elements with the same value, returns true if
//the first element's key is smaller alphabetically than the second's.
func (p TupleList) Less(i, j int) bool {
	if p[i].Value != p[j].Value {
		return p[i].Value < p[j].Value
	} else {
		strs := []string{p[i].Key, p[j].Key}
		sort.Strings(strs)
		return strs[0] != p[i].Key
	}
}
