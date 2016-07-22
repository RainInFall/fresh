package fresh

//go:generate js-like array string

func (array Arraystring) Len() int {
	return len(array)
}

func (array Arraystring) Less(i, j int) bool {
	return false
}

func (array Arraystring) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}
