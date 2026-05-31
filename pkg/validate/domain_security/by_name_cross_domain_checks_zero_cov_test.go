//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestByName_ZeroCov — domain_security 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestByNameCrossDomainChecks_ZeroCov(t *testing.T) {
	opDomain := map[string]string{"OtherOp": "other"}

	fetch := stml.FetchBlock{
		OperationID:   "OtherOp",
		NestedFetches: []stml.FetchBlock{{OperationID: "OtherOp"}},
	}
	if d := checkFetchCrossDomain(fetch, "core/page.html", "core", opDomain); len(d) == 0 {
		t.Errorf("checkFetchCrossDomain expected cross-domain warning")
	}

	page := stml.PageSpec{
		FileName: "core/page.html",
		Fetches:  []stml.FetchBlock{fetch},
		Actions:  []stml.ActionBlock{{OperationID: "OtherOp"}},
	}
	if d := checkPageCrossDomain(page, "core", opDomain); len(d) == 0 {
		t.Errorf("checkPageCrossDomain expected cross-domain warnings")
	}
}
