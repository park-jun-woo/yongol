//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-25 test (TODO: add cases)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS25UnknownSeqType(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
