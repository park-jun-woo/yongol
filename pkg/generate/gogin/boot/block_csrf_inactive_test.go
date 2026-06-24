//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what blockCsrf — middleware.Csrf 등록 (쿠키 인증 조건부, Phase005)
package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

func TestBlockCsrf_Inactive(t *testing.T) {
	// Default bearer (no csrf) → inert block with always-inactive predicate.
	block := blockCsrf(nil, prepared.Auth{Present: false}, "example.com/zenflow")
	if block.Active == nil || block.Active(nil) {
		t.Errorf("non-csrf block must carry csrfAlwaysInactive predicate")
	}
	if len(block.Lines) != 0 {
		t.Errorf("inactive csrf must emit no lines, got %v", block.Lines)
	}
}
