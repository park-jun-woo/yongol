//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=trusted-proxy
//ff:what blockRouter — trusted_proxies 미설정 시 SetTrustedProxies(nil 기본) + fail-fast 방출 회귀

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockRouter_TrustedProxiesDefault(t *testing.T) {
	// No manifest backend.http.trusted_proxies → emitted default is nil
	// (trust no proxy), still overridable via BACKEND_HTTP_TRUSTED_PROXIES.
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
	block := blockRouter(fs, "example.com/zenflow")

	if !strings.Contains(strings.Join(block.Lines, "\n"),
		`trustedProxies := envStringList("BACKEND_HTTP_TRUSTED_PROXIES", nil)`) {
		t.Errorf("default must resolve trusted proxies with nil fallback, got:\n%s",
			strings.Join(block.Lines, "\n"))
	}

	// SetTrustedProxies must come immediately after gin.Default(), before
	// any middleware registration (otelgin etc.).
	ginIdx, setIdx := -1, -1
	for i, l := range block.Lines {
		if strings.Contains(l, "gin.Default()") {
			ginIdx = i
		}
		if strings.Contains(l, "trustedProxies := envStringList") {
			setIdx = i
		}
	}
	if ginIdx == -1 || setIdx != ginIdx+1 {
		t.Errorf("trusted proxy resolution must directly follow gin.Default() (gin=%d set=%d):\n%s",
			ginIdx, setIdx, strings.Join(block.Lines, "\n"))
	}

	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "r.SetTrustedProxies(trustedProxies)") {
		t.Errorf("must emit SetTrustedProxies call, got:\n%s", body)
	}
	if !strings.Contains(body, "os.Exit(1)") {
		t.Errorf("invalid CIDR must fail-fast at bootstrap, got:\n%s", body)
	}

	imports := strings.Join(block.Imports, "\n")
	for _, want := range []string{`"log/slog"`, `"os"`} {
		if !strings.Contains(imports, want) {
			t.Errorf("missing import %s, got:\n%s", want, imports)
		}
	}
}
