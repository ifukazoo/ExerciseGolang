package compress

import (
	"unicode"
	"unicode/utf8"
)

func compress(bytes []byte) []byte {

	var (
		r    rune
		size int
	)
	const asciiSpace = byte(0x20)

	idx := 0
	for i := 0; i < len(bytes); {
		r, size = utf8.DecodeRune(bytes[i:])
		if unicode.IsSpace(r) {
			bytes[idx] = asciiSpace
			idx++
			i += size

			// 残りの空白は捨てる
			for i < len(bytes) {
				r, size = utf8.DecodeRune(bytes[i:])
				if !unicode.IsSpace(r) {
					break
				}
				i += size
			}
		} else {
			for j := 0; j < size; j++ {
				bytes[idx] = bytes[i]
				idx++
				i++
			}
		}
	}
	return bytes[:idx]
}
