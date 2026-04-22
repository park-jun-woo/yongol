//ff:func feature=chain type=test control=sequence
//ff:what Chain 호출 시 OpenAPIDoc 이 nil 이면 에러 반환 검증
package chain

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestChain_NoOpenAPI(t *testing.T) {
	fs := &yongol.Fullstack{SpecsDir: t.TempDir()}
	if _, err := Chain(fs, "Anything"); err == nil {
		t.Fatal("expected error when OpenAPIDoc is nil")
	}
}
