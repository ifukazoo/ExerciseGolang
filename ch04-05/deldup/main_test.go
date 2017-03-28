package deldup

import (
	"reflect"
	"testing"
)

func Test_deldup(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "a"}, []string{"a"}},
		{[]string{"a", "a", "a"}, []string{"a"}},
		{[]string{"abc", "abc"}, []string{"abc"}},
	}
	for _, test := range tests {
		if got := deldup(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("test[%v] return[%v], want[%v]", test.input, got, test.want)
		}
	}
}
