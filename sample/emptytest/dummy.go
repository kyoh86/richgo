package emptytest

func Dummy() {
	for k := range make([]struct{}, 3) {
		k *= 2
	}
}
