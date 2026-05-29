//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what test: TestFileNameForFunc — PascalCase → snake_case.go 변환 케이스 검증

package fffile

import "testing"

func TestFileNameForFunc(t *testing.T) {
	cases := map[string]string{
		"ActivateWorkflow": "activate_workflow.go",
		"convertWorkflow":  "convert_workflow.go",
		"HTTPServer":       "http_server.go",
		"":                 "",
	}
	for in, want := range cases {
		if got := FileNameForFunc(in); got != want {
			t.Errorf("FileNameForFunc(%q) = %q, want %q", in, got, want)
		}
	}
}
