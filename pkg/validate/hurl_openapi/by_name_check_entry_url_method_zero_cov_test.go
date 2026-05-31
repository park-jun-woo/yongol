//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — checkEntryURLMethod / xoh12 / xoh13 분기 직접 호출
package hurl_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestByNameCheckEntryURLMethod_ZeroCov(t *testing.T) {
	routes := collectOpenAPIRoutes(miniDoc())

	// external-service entry (URLVar set, not host) → nil.
	ext := hurl.HurlEntry{Method: "GET", Path: "/x", URLVar: "authurl"}
	if d := checkEntryURLMethod(ext, routes); d != nil {
		t.Errorf("external entry should be skipped, got %v", d)
	}

	// matching own-API entry → nil.
	match := hurl.HurlEntry{Method: "GET", Path: "/widgets", URLVar: "host"}
	if d := checkEntryURLMethod(match, routes); d != nil {
		t.Errorf("matching entry should yield nil, got %v", d)
	}

	// drift entry → diagnostic.
	drift := hurl.HurlEntry{Method: "GET", Path: "/nonexistent", URLVar: "host", File: "t.hurl", Line: 1}
	if d := checkEntryURLMethod(drift, routes); d == nil {
		t.Errorf("drift entry should yield diagnostic")
	}
}
