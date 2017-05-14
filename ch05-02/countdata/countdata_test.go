package countdata

import (
	"strings"
	"testing"
)

func Test_Count(t *testing.T) {
	tests := []struct {
		input string
	}{
		{`<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`},
	}
	for _, test := range tests {
		if got, err := CountData(strings.NewReader(test.input)); err != nil {
			t.Errorf("test(%v) error[%v]", test.input, err)
		} else {
			if got["html"] != 1 {
				t.Errorf("test(%v) got[%v]", test.input, got)
			}
			if got["head"] != 1 {
				t.Errorf("test(%v) got[%v]", test.input, got)
			}
			if got["body"] != 1 {
				t.Errorf("test(%v) got[%v]", test.input, got)
			}
			if got["p"] != 1 {
				t.Errorf("test(%v) got[%v]", test.input, got)
			}
			if got["a"] != 2 {
				t.Errorf("test(%v) got[%v]", test.input, got)
			}
			if got["ul"] != 1 {
				t.Errorf("test(%v) got[%v]", test.input, got)
			}
			if got["li"] != 2 {
				t.Errorf("test(%v) got[%v]", test.input, got)
			}
		}
	}
}
