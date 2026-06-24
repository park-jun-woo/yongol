//ff:func feature=gen-gogin type=test control=sequence
//ff:what generate_integ — dummy specs 로 실제 Fullstack 구성 후 gogin/ssac.Generate 통합 커버리지
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGoginSSaCGenerate_EmptyEarlyReturn_Integ(t *testing.T) {
	if err := Generate(&yongol.Fullstack{}, t.TempDir(), "", ""); err != nil {
		t.Fatalf("empty Generate should be a no-op nil, got %v", err)
	}
}
