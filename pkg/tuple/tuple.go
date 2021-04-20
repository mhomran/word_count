package tuple

import "sort"

type Tuple struct {
	Key   string
	Value int
}

type TupleList []Tuple

func (p TupleList) Len() int { return len(p) }

func (p TupleList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p TupleList) Less(i, j int) bool {
	if p[i].Value != p[j].Value {
		return p[i].Value < p[j].Value
	} else {
		strs := []string{p[i].Key, p[j].Key}
		sort.Strings(strs)
		return strs[0] != p[i].Key
	}
}
