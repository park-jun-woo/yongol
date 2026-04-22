//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what test: TestFileNameForMethod — receiver 유무에 따른 파일명 구성 검증

package fffile

import "testing"

func TestFileNameForMethod(t *testing.T) {
	type in struct {
		recv   string
		method string
	}
	cases := map[in]string{
		{"Server", "ActivateWorkflow"}: "server_activate_workflow.go",
		{"", "ActivateWorkflow"}:       "activate_workflow.go",
		{"Server", ""}:                 "",
	}
	for k, want := range cases {
		if got := FileNameForMethod(k.recv, k.method); got != want {
			t.Errorf("FileNameForMethod(%q,%q) = %q, want %q", k.recv, k.method, got, want)
		}
	}
}
