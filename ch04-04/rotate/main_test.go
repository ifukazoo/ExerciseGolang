package rotate

import (
	"reflect"
	"testing"
)

func Test_rotateSimple(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 int
		want   []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{[]int{1}, 1, []int{1}},
	}
	for _, test := range tests {
		if rotateSimple(test.input1, test.input2); !reflect.DeepEqual(test.input1, test.want) {
			t.Errorf("test(%v) return[%v], want[%v]", test.input2, test.input1, test.want)
		}
	}
}

func Test_rotateFor(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 int
		want   []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{[]int{1}, 1, []int{1}},
	}
	for _, test := range tests {
		if rotateFor(test.input1, test.input2); !reflect.DeepEqual(test.input1, test.want) {
			t.Errorf("test(%v) return[%v], want[%v]", test.input2, test.input1, test.want)
		}
	}
}
