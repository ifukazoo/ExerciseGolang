package countbit

import (
	"crypto/sha256"
	"testing"
)

func Test_countDifBit(t *testing.T) {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("x"))
	b[0] = b[0] &^ 1
	var c SHA256
	d := SHA256{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	tests := []struct {
		input [2]SHA256
		want  int
	}{
		{[2]SHA256{a, a}, 0},
		{[2]SHA256{a, b}, 1},
		{[2]SHA256{c, d}, 256},
	}
	for _, test := range tests {
		if got := countDifBit(test.input[0], test.input[1]); got != test.want {
			t.Errorf("test(%v) return[%v], want[%v]", test.input, got, test.want)
		}
	}
}
