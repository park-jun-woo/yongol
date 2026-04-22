//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what XDD-61 test (TODO: add cases)

package ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdd61SensitiveNoAnnotation(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
