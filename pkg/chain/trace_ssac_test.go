//ff:func feature=chain type=test control=sequence
//ff:what traceSSaC 가 sequence 타입을 중복없이 요약하고 SSaC 파일 위치를 찾는지 검증
package chain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestTraceSSaC(t *testing.T) {
	specsDir := t.TempDir()
	svcDir := filepath.Join(specsDir, "service")
	if err := os.MkdirAll(svcDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	content := "service x\n\nfunc GetCourse(id int) {\n}\n"
	if err := os.WriteFile(filepath.Join(svcDir, "course.ssac"), []byte(content), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	sf := &ssac.ServiceFunc{
		Name:     "GetCourse",
		FileName: "course.ssac",
		Sequences: []ssac.Sequence{
			{Type: "get"},
			{Type: "get"}, // duplicate type must be deduped
			{Type: "response"},
		},
	}

	link := traceSSaC(sf, specsDir)
	if link.Kind != "SSaC" || link.File != filepath.Join("service", "course.ssac") {
		t.Errorf("link fields: %+v", link)
	}
	if link.Summary != "@get @response" {
		t.Errorf("summary: got %q, want %q", link.Summary, "@get @response")
	}
	if link.Line != 3 {
		t.Errorf("line: got %d, want 3", link.Line)
	}
}
