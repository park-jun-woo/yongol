//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what XDD-61 테스트 (TODO: 케이스 추가)

package ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdd61SensitiveNoAnnotation(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
