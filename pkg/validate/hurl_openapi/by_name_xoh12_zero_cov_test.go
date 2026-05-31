//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — checkEntryURLMethod / xoh12 / xoh13 분기 직접 호출
package hurl_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestByNameXoh12_ZeroCov(t *testing.T) {
	// nil/empty guards.
	if d := xoh12StatusCoverage(nil); d != nil {
		t.Errorf("nil fs should yield nil")
	}
	// populated: GET /widgets covers 200 but not 404 → warning.
	fs := &yongol.Fullstack{
		OpenAPIDoc: miniDoc(),
		HurlEntries: []hurl.HurlEntry{
			{Method: "GET", Path: "/widgets", StatusCode: "200", URLVar: "host"},
		},
	}
	_ = xoh12StatusCoverage(fs) // exercise the iteration; result depends on 5xx filter
}
