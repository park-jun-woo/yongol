//ff:func feature=validate type=test control=sequence topic=funcspec-structural
//ff:what XFF-41 테스트 (TODO: 케이스 추가)

package funcspec

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXff41FuncForbiddenImport(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
