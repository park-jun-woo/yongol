//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what hasExplicit2xx — nil/없음/2xx 존재 검증

package openapi_ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestHasExplicit2xx(t *testing.T) {
	t.Run("nil op returns false", func(t *testing.T) {
		if hasExplicit2xx(nil) {
			t.Error("expected false")
		}
	})

	t.Run("nil responses returns false", func(t *testing.T) {
		if hasExplicit2xx(&openapi3.Operation{}) {
			t.Error("expected false")
		}
	})

	t.Run("no 2xx returns false", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("404", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		op := &openapi3.Operation{Responses: resps}
		if hasExplicit2xx(op) {
			t.Error("expected false")
		}
	})

	t.Run("200 returns true", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		op := &openapi3.Operation{Responses: resps}
		if !hasExplicit2xx(op) {
			t.Error("expected true")
		}
	})

	t.Run("201 returns true", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("201", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		op := &openapi3.Operation{Responses: resps}
		if !hasExplicit2xx(op) {
			t.Error("expected true")
		}
	})
}
