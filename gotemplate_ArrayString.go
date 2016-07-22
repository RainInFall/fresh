package fresh

import (
	"sort"
)

// template type Array(A)

//Array has js-like functions
type ArrayString []string

/*
Some tests whether some element in the array passes the test implemented by the provided function
*/
func (array ArrayString) Some(f func(string, int, ArrayString) bool) bool {
	for index, value := range array {
		if f(value, index, array) {
			return true
		}
	}
	return false
}

/*
Reverse the Array
*/
func (array ArrayString) Reverse() ArrayString {
	mid := len(array) / 2
	for i, j := 0, len(array)-1; i < mid; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

/*
Sort the Array
*/
func (array ArrayString) Sort() ArrayString {
	sort.Sort(array)
	return array
}
