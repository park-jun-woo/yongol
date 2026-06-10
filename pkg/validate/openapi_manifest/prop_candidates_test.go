//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what propCandidates — 빈 맵 힌트 / 정렬된 후보 응답 필드명 목록 검증

package openapi_manifest

import (
	"testing"
)

func TestPropCandidates(t *testing.T) {
	t.Run("empty map returns hint", func(t *testing.T) {
		got := propCandidates(map[string]bool{})
		want := "No OpenAPI 2xx response declares object properties"
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	t.Run("nil map returns hint", func(t *testing.T) {
		got := propCandidates(nil)
		want := "No OpenAPI 2xx response declares object properties"
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	t.Run("renders sorted field list", func(t *testing.T) {
		props := map[string]bool{"token": true, "expiresIn": true, "user": true}
		got := propCandidates(props)
		want := "Available 2xx response fields: expiresIn, token, user"
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})
}
