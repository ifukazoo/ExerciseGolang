package reverse

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		input [5]int
		want  [5]int
	}{
		{[5]int{0, 1, 2, 3, 4}, [5]int{4, 3, 2, 1, 0}},
		{[5]int{}, [5]int{}},
	}
	for _, test := range tests {
		if reverse(&test.input); test.input != test.want {
			t.Errorf("test(%v) return[%v], want[%v]", test.input, test.input, test.want)
		}
	}
}
