package main

import (
	"algorithm/jsort"
)

type word int

func (w word) Less(i interface{}) int {
	if int(w)-int(i.(word)) < 0 {
		return -1
	} else if int(w)-int(i.(word)) > 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	arr := []word{word(8), word(7), word(4), word(2)}
	arrCov := make([]jsort.Comparable, len(arr))
	for i, v := range arr {
		arrCov[i] = jsort.Comparable(v)
	}

	jsort.SelectSort(arrCov)

	jsort.Show(arrCov)
}
