//ff:func feature=cli type=test control=sequence
//ff:what TestVersionCmd — versionCmd 서브커맨드 출력 검증

package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	cmd := versionCmd()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetArgs([]string{})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := buf.String()
	if !strings.HasPrefix(out, "yongol ") {
		t.Errorf("expected output to start with 'yongol ', got: %q", out)
	}
}
