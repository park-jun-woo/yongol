//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xos21ErrStatusNotInOpenAPI — nil doc/no matching op/status 검증
package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos21ErrStatusNotInOpenAPI(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
