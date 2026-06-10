//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what opCandidates — 빈 맵 힌트 / 정렬된 후보 operationId 목록 검증

package openapi_manifest

import (
	"testing"
)

func TestOpCandidates(t *testing.T) {
	t.Run("empty map returns hint", func(t *testing.T) {
		got := opCandidates(map[string]bool{})
		want := "The OpenAPI document declares no operationIds"
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	t.Run("nil map returns hint", func(t *testing.T) {
		got := opCandidates(nil)
		want := "The OpenAPI document declares no operationIds"
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	t.Run("renders sorted candidate list", func(t *testing.T) {
		ids := map[string]bool{"listUsers": true, "createUser": true, "getUser": true}
		got := opCandidates(ids)
		want := "Declared operationIds: createUser, getUser, listUsers"
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})
}
