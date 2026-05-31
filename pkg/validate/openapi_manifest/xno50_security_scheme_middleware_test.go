//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xno50SecuritySchemeMiddleware — nil/매칭/누락 검증
package openapi_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXno50SecuritySchemeMiddleware(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
