//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerate — nil fs early-return + copy 에러 + 전체 성공 경로 검증
package middleware

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

func TestGenerate_NilFullstack(t *testing.T) {
	if err := Generate(nil, prepared.State{}, t.TempDir()); err != nil {
		t.Errorf("nil fs should return nil, got: %v", err)
	}
}
