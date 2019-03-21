package jsort

import "fmt"

type Comparable interface {
	//Less(Comparable, Comparable) int
	Less(i interface{}) int
	//Exch([]interface{}, int, int)
}

func SelectSort(a []Comparable) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i].Less(a[j]) > 0 {
				min := a[j]
				a[j] = a[i]
				a[i] = min
			}
		}
	}
}

func IsSort(a []Comparable) bool {
	for i, v := range a[:len(a)-2] {
		if v.Less(a[i+1]) > 0 {
			return false
		}
	}
	return true
}

func Show(a []Comparable) {
	fmt.Println(a)
}
