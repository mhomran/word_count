package tuple

import "sort"

type Tuple struct {
	Key   string
	Value int
}

type TupleList []Tuple

func (p TupleList) Less(i, j int) bool {
	if p[i].Value != p[j].Value {
		return p[i].Value < p[j].Value
	} else {
		strs := []string{p[i].Key, p[j].Key}
		sort.Strings(strs)
		return strs[0] == p[i].Key
	}
}
