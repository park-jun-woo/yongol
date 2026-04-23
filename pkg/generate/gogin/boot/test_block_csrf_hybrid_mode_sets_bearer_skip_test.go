//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what TestBlockCsrf_HybridMode_SetsBearerSkip — hybrid 모드에서 HybridBearerSkip true 설정 확인

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockCsrf_HybridMode_SetsBearerSkip(t *testing.T) {
	raw := &pmanifest.Auth{
		Mode: "hybrid",
		Csrf: &pmanifest.CsrfConfig{Enabled: true},
	}
	a := prepared.Auth{Present: true, Mode: "hybrid", Raw: raw}
	block := blockCsrf(a, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "HybridBearerSkip: true") {
		t.Fatalf("hybrid mode should set HybridBearerSkip:true, got:\n%s", body)
	}
}
