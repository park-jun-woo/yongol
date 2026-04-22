//ff:func feature=validate type=test control=sequence topic=funcspec-structural
//ff:what F-1 테스트 (TODO: 케이스 추가)

package funcspec

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestF01BuiltinOverride(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
