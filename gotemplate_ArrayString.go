package fresh

// template type Array(A)

//Array has js-like functions
type ArrayString []string

/*
Some
*/
func (array ArrayString) Some(f func(string, int, ArrayString) bool) bool {
	for index, value := range array {
		if f(value, index, array) {
			return true
		}
	}
	return false
}
