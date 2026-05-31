//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockAuthInit — auth.Configure + RefreshStore 주입 (라우트 마운트 없음)
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockAuthInit_Present(t *testing.T) {
	a := prepared.Auth{Present: true, Mode: "cookie", Raw: &pmanifest.Auth{}}
	block := blockAuthInit(a, "example.com/zenflow")
	if block.Active != nil {
		t.Errorf("present auth block should leave Active nil (caller-gated)")
	}
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "configureAuth(") {
		t.Errorf("must wire configureAuth, got:\n%s", body)
	}
	if !strings.Contains(body, "auth.Init(infraauth.NewPostgres(queries))") {
		t.Errorf("must install RefreshStore singleton, got:\n%s", body)
	}
	// parseSameSite + configureAuth helpers emitted.
	if len(block.Funcs) != 2 {
		t.Errorf("present auth must emit 2 helper funcs, got %d", len(block.Funcs))
	}
}
