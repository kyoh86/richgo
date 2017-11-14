// +build sample

package cover05

// Cover05 will be covered about 50%
func Cover05(arg int) string {
	switch arg {
	case 1:
		return "case-1"
	case 2:
		return "case-2"
	default:
		return "others"
	}
}
