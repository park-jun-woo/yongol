//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what TestByName_ZeroCov — O-6 스키마 워커들을 이름으로 직접 호출해 커버리지 귀속
package openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestByNameXOE01_ZeroCov(t *testing.T) {
	doc := byNameO06Doc()
	fs := byNameO06FS(doc)
	_ = xoe01ErrorResponseRequired(fs)
	_ = xoe01ErrorResponseRequired(&yongol.Fullstack{})
}
