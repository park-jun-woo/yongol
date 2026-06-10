//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what collectFrontendAuthProps — nil doc / 프로퍼티 합집합 / refresh_op 매칭·미매칭·빈 이름 검증

package openapi_manifest

import (
	"testing"
)

func TestCollectFrontendAuthProps(t *testing.T) {
	t.Run("nil doc returns empty sets and not found", func(t *testing.T) {
		all, refresh, found := collectFrontendAuthProps(nil, "Refresh")
		if len(all) != 0 || len(refresh) != 0 || found {
			t.Errorf("expected empty/empty/false, got %v/%v/%v", all, refresh, found)
		}
	})

	doc := buildDocWithResponseFields(map[string][]string{
		"Login":   {"access_token", "refresh_token"},
		"Refresh": {"access_token"},
	})

	t.Run("collects union and refresh op props", func(t *testing.T) {
		all, refresh, found := collectFrontendAuthProps(doc, "Refresh")
		if !found {
			t.Fatalf("expected refresh op found")
		}
		if !all["access_token"] || !all["refresh_token"] {
			t.Errorf("expected union of all props, got %v", all)
		}
		if !refresh["access_token"] || refresh["refresh_token"] {
			t.Errorf("expected refresh op props {access_token}, got %v", refresh)
		}
	})

	t.Run("unknown refresh op is not found", func(t *testing.T) {
		_, refresh, found := collectFrontendAuthProps(doc, "Nonexistent")
		if found || len(refresh) != 0 {
			t.Errorf("expected not found and empty refresh props, got %v/%v", found, refresh)
		}
	})

	t.Run("empty refresh op id is not found", func(t *testing.T) {
		_, _, found := collectFrontendAuthProps(doc, "")
		if found {
			t.Errorf("expected not found for empty refreshOp")
		}
	})
}
