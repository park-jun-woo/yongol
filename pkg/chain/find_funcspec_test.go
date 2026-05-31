//ff:func feature=chain type=test control=iteration dimension=1
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

	cases := []findFuncSpecCase{
		{
			name:      "match case-insensitive",
			pkg:       "mail",
			funcName:  "SendWelcome",
			wantOK:    true,
			wantFile:  "func/mail/send_welcome.go",
			wantSumm:  "@func mail.SendWelcome",
			wantLineP: true,
		},
		{name: "wrong package", pkg: "auth", funcName: "SendWelcome", wantOK: false},
		{name: "wrong name", pkg: "mail", funcName: "SomethingElse", wantOK: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assertFindFuncSpec(t, tc, specs, specsDir)
		})
	}
}
