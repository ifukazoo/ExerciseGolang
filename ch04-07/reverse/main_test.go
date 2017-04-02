package reverse

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte("h"), []byte("h")},
		{[]byte("hello"), []byte("olleh")},
		{[]byte("世"), []byte("世")},
		{[]byte("世 界"), []byte("界 世")},
	}
	for _, test := range tests {
		if got := reverse(test.input); !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("return[% x], want[% x]", got, test.want)
		}
	}
}
