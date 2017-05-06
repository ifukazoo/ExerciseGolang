package format

import (
	"html/template"
	"os"
)

// PrintFmt :出力形式
type PrintFmt struct {
	// Url
	URL        string
	Transcript string
}

const templ = `
================================================================================
URL:{{.URL}}
Transcript:
vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
{{.Transcript}}
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`

var commic *template.Template

// Print 出力
func Print(url, transcript string) error {
	if commic == nil {
		commic = template.Must(template.New("commic").Parse(templ))
	}
	return commic.Execute(os.Stdout, PrintFmt{url, transcript})
}
