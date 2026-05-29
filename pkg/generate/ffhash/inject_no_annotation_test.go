//ff:func feature=gen-ffhash type=test control=sequence
//ff:what test: TestInjectSkipsFilesWithoutAnnotation — //ff: 블록이 없는 파일은 수정하지 않음

package ffhash

import (
	"bytes"
	"testing"
)

func TestInjectSkipsFilesWithoutAnnotation(t *testing.T) {
	plain := []byte("package demo\n\nfunc Demo() {}\n")
	out := InjectCheckedLine(plain)
	if bytes.Contains(out, []byte("//ff:checked")) {
		t.Errorf("expected no injection when annotation block is missing, got:\n%s", out)
	}
}
