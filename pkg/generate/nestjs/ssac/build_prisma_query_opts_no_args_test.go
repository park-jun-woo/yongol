//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestBuildPrismaQueryOptsNoArgs — TestBuildPrismaQueryOpts — GetOp → where + pagination 옵션 목록 구성 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestBuildPrismaQueryOptsNoArgs(t *testing.T) {
	op := &ir.GetOp{}
	if opts := buildPrismaQueryOpts(op); len(opts) != 0 {
		t.Errorf("opts = %v, want empty when no args/pagination", opts)
	}
}
