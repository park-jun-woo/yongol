//ff:func feature=chain type=test control=selection dimension=2
//ff:what TestFindFuncSpec — findFuncSpecLink 의 매칭/불일치/대소문자 분기 검증

package chain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestFindFuncSpec(t *testing.T) {
	specsDir := t.TempDir()
	pkgDir := filepath.Join(specsDir, "func", "mail")
	if err := os.MkdirAll(pkgDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	// File body so grepLine resolves a concrete line number for the func name.
	if err := os.WriteFile(filepath.Join(pkgDir, "send_welcome.go"), []byte("// @func\n// header\nfunc SendWelcome() {}\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	specs := []funcspec.FuncSpec{{Package: "mail", Name: "sendWelcome"}}

	cases := []struct {
		name      string
		pkg       string
		funcName  string
		wantOK    bool
		wantFile  string
		wantSumm  string
		wantLineP bool // whether to assert a positive line number
	}{
		{
			name:      "match case-insensitive",
			pkg:       "mail",
			funcName:  "SendWelcome",
			wantOK:    true,
			wantFile:  "func/mail/send_welcome.go",
			wantSumm:  "@func mail.SendWelcome",
			wantLineP: true,
		},
		{
			name:     "wrong package",
			pkg:      "auth",
			funcName: "SendWelcome",
			wantOK:   false,
		},
		{
			name:     "wrong name",
			pkg:      "mail",
			funcName: "SomethingElse",
			wantOK:   false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			callRef := tc.pkg + "." + tc.funcName
			link, ok := findFuncSpecLink(callRef, tc.pkg, tc.funcName, specs, specsDir)
			if ok != tc.wantOK {
				t.Fatalf("ok: got %v, want %v", ok, tc.wantOK)
			}
			if !tc.wantOK {
				return
			}
			if link.Kind != "FuncSpec" {
				t.Errorf("kind: got %q, want FuncSpec", link.Kind)
			}
			if link.File != tc.wantFile {
				t.Errorf("file: got %q, want %q", link.File, tc.wantFile)
			}
			if link.Summary != tc.wantSumm {
				t.Errorf("summary: got %q, want %q", link.Summary, tc.wantSumm)
			}
			if tc.wantLineP && link.Line <= 0 {
				t.Errorf("line: got %d, want > 0", link.Line)
			}
		})
	}
}
