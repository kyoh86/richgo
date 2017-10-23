package emptytest

// Dummy is dummy.
func Dummy() int {
	var j int
	for k := range make([]struct{}, 3) {
		j += k
		j *= 2
	}
	return j
}
