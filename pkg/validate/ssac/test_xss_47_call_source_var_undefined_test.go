//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what XSS-47 test (TODO: add cases)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXss47CallSourceVarUndefined(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
