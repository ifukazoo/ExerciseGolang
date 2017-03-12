package countbit

var pc [256]byte

// SHA256 sha256
type SHA256 [32]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func countDifBit(a, b SHA256) int {
	diffBit := 0
	for i := range a {
		diffBit += int(pc[a[i]^b[i]])
	}
	return diffBit
}
