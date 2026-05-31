//ff:func feature=gen-gogin type=test control=sequence topic=rate-limit
//ff:what blockRateLimit — manifest.backend.rate_limit → RouteRateLimit 미들웨어 등록
package boot

import (
	"testing"
)

func TestBlockRateLimit_NilFullstack(t *testing.T) {
	block := blockRateLimit(nil, "example.com/zenflow")
	if len(block.Lines) != 0 {
		t.Errorf("nil Fullstack must yield inert block, got %v", block.Lines)
	}
}
