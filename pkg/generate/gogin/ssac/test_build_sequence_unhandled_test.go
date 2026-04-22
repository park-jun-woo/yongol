//ff:func feature=gen-gogin type=test control=selection
//ff:what TestBuildSequenceUnhandled — 알 수 없는 seq.Type 에 대해 buildSequence 가 error 반환

package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// Phase011 — silent fall-through 대신 명시 error 반환 보장.
func TestBuildSequenceUnhandled(t *testing.T) {
	g := &methodGen{
		FuncName: "FakeHandler",
		FileName: "fake_service.ssac",
	}
	seq := ssacparser.Sequence{
		Type: "foobar",
		Line: 42,
	}

	lines, imports, isPost, err := g.buildSequence(seq, nil)
	if err == nil {
		t.Fatalf("expected error for unhandled seq.Type=%q, got nil (lines=%v)", seq.Type, lines)
	}
	if lines != nil {
		t.Errorf("expected nil lines on error, got %v", lines)
	}
	if imports != nil {
		t.Errorf("expected nil imports on error, got %v", imports)
	}
	if isPost {
		t.Errorf("expected isPost=false on error, got true")
	}

	msg := err.Error()
	wants := []string{
		"unhandled",
		`"foobar"`,
		"fake_service.ssac",
		"line=42",
		"FakeHandler",
	}
	for _, w := range wants {
		if !strings.Contains(msg, w) {
			t.Errorf("error message missing %q: got %q", w, msg)
		}
	}
}
