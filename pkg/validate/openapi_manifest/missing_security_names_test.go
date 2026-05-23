//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what missingSecurityNames — nil op/nil security/매칭/누락 검증

package openapi_manifest

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestMissingSecurityNames(t *testing.T) {
	mwSet := map[string]bool{"bearerAuth": true}

	t.Run("nil op returns nil", func(t *testing.T) {
		got := missingSecurityNames(nil, mwSet)
		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})

	t.Run("nil security returns nil", func(t *testing.T) {
		op := &openapi3.Operation{}
		got := missingSecurityNames(op, mwSet)
		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})

	t.Run("all in set returns nil", func(t *testing.T) {
		sec := openapi3.SecurityRequirements{
			{"bearerAuth": {}},
		}
		op := &openapi3.Operation{Security: &sec}
		got := missingSecurityNames(op, mwSet)
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("missing name returned", func(t *testing.T) {
		sec := openapi3.SecurityRequirements{
			{"oauth2": {}},
		}
		op := &openapi3.Operation{Security: &sec}
		got := missingSecurityNames(op, mwSet)
		if len(got) != 1 || got[0] != "oauth2" {
			t.Errorf("expected [oauth2], got %v", got)
		}
	})
}
