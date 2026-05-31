//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestDomainSecurityHelpers — unit tests for the pure domain_security helper functions
package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestClassifyPageDomain(t *testing.T) {
	dirs := map[string]string{
		"admin":  "admin",
		"public": "public",
	}
	if got := classifyPageDomain(stml.PageSpec{FileName: "admin/x.stml"}, dirs); got != "admin" {
		t.Errorf("got %q, want admin", got)
	}
	if got := classifyPageDomain(stml.PageSpec{FileName: "other/x.stml"}, dirs); got != "" {
		t.Errorf("unmatched should give empty, got %q", got)
	}
}
