//ff:func feature=validate-contract type=test control=sequence
//ff:what TestRun — arts 디렉토리 preserved 파일 대상 PRV 규칙 오케스트레이션 검증

package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("nil fs / empty dir → nil", func(t *testing.T) {
		if d := Run(nil, "x"); d != nil {
			t.Errorf("nil fs should return nil, got %+v", d)
		}
		if d := Run(buildFSWithOp(), ""); d != nil {
			t.Errorf("empty artsDir should return nil, got %+v", d)
		}
	})

	t.Run("no preserved files → nil", func(t *testing.T) {
		dir := t.TempDir()
		if d := Run(buildFSWithOp(), dir); d != nil {
			t.Errorf("empty arts dir should return nil, got %+v", d)
		}
	})

	t.Run("flags panic in preserved file", func(t *testing.T) {
		dir := t.TempDir()
		p := filepath.Join(dir, "do_thing.go")
		writePreserved(t, p, "package service\nfunc DoThing() { panic(\"x\") }\n")
		diags := Run(buildFSWithOp(), dir)
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, "[PRV-10]") {
				found = true
			}
		}
		if !found {
			t.Fatalf("expected PRV-10 diag from Run, got %+v", diags)
		}
	})
}
