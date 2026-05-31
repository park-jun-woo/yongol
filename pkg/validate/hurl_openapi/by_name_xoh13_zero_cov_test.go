//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — checkEntryURLMethod / xoh12 / xoh13 분기 직접 호출
package hurl_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestByNameXoh13_ZeroCov(t *testing.T) {
	if d := xoh13GuardCoverage(nil); d != nil {
		t.Errorf("nil fs should yield nil")
	}
	// missing service funcs → nil early return.
	fs := &yongol.Fullstack{OpenAPIDoc: miniDoc(), HurlEntries: []hurl.HurlEntry{{Method: "GET", Path: "/widgets"}}}
	if d := xoh13GuardCoverage(fs); d != nil {
		t.Errorf("no service funcs should yield nil, got %v", d)
	}
}
