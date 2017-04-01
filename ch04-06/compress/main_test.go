package compress

import (
	"reflect"
	"testing"
)

func Test_compress(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte("hello world"), []byte("hello world")},
		{[]byte("hello  world"), []byte("hello world")},
		{[]byte("世界"), []byte("世界")},
		{[]byte("世 界"), []byte("世 界")},
		{[]byte("世  界"), []byte("世 界")},
		{[]byte("世界  "), []byte("世界 ")},
		{[]byte("世\n界"), []byte("世 界")},
		{[]byte("世\n 界"), []byte("世 界")},
		{[]byte("世 \n界"), []byte("世 界")},
		{[]byte("世\u0085界"), []byte("世 界")},
		{[]byte("世\u00A0界"), []byte("世 界")},
		{[]byte("世\u0085\u00A0界"), []byte("世 界")},
		{[]byte("世\n\u0085\u00A0界"), []byte("世 界")},
	}
	for _, test := range tests {
		if got := compress(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("return[% x], want[% x]", got, test.want)
		}
	}
}
