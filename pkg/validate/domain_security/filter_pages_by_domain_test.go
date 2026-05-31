//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestDomainSecurityHelpers — unit tests for the pure domain_security helper functions
package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFilterPagesByDomain(t *testing.T) {
	pages := []stml.PageSpec{
		{FileName: "admin/dashboard.stml"},
		{FileName: "public/home.stml"},
	}
	got := filterPagesByDomain(pages, "admin")
	if len(got) != 1 || got[0].FileName != "admin/dashboard.stml" {
		t.Errorf("filterPagesByDomain = %+v", got)
	}
	// empty dir → nil.
	if got := filterPagesByDomain(pages, ""); got != nil {
		t.Errorf("empty dir → %v, want nil", got)
	}
	// trailing slash tolerated.
	if got := filterPagesByDomain(pages, "admin/"); len(got) != 1 {
		t.Errorf("trailing slash dir → %+v", got)
	}
}
