package editor

// Parrot will not change input
func Parrot() Editor {
	return &parrot{}
}

// parrot through output raw.
type parrot struct{}

func (e *parrot) Edit(line string) (string, error) {
	return line, nil
}
