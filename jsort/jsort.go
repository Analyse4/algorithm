package jsort

type Comparable interface {
	//Less(Comparable, Comparable) int
	Less(i interface{}, j interface{}) int
	Exch([]interface{}, int, int)
}

func SelectS(a []Comparable) {

}

func IsSort(a []Comparable) int {

}
