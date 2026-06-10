//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what opRequiresAuth — op nil / 명시 security / opt-out / doc 상속 / doc nil 분기 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestOpRequiresAuth(t *testing.T) {
	sec := openapi3.SecurityRequirements{openapi3.SecurityRequirement{"bearerAuth": {}}}
	empty := openapi3.SecurityRequirements{}

	t.Run("nil op returns false", func(t *testing.T) {
		if opRequiresAuth(nil, &openapi3.T{}) {
			t.Errorf("expected false for nil op")
		}
	})

	t.Run("explicit non-empty security returns true", func(t *testing.T) {
		op := &openapi3.Operation{Security: &sec}
		if !opRequiresAuth(op, nil) {
			t.Errorf("expected true for explicit security")
		}
	})

	t.Run("explicit empty security opts out", func(t *testing.T) {
		op := &openapi3.Operation{Security: &empty}
		if opRequiresAuth(op, &openapi3.T{Security: sec}) {
			t.Errorf("expected false for explicit opt-out")
		}
	})

	t.Run("inherits from doc security", func(t *testing.T) {
		op := &openapi3.Operation{}
		if !opRequiresAuth(op, &openapi3.T{Security: sec}) {
			t.Errorf("expected true inheriting doc security")
		}
	})

	t.Run("no doc security returns false", func(t *testing.T) {
		op := &openapi3.Operation{}
		if opRequiresAuth(op, &openapi3.T{}) {
			t.Errorf("expected false with empty doc security")
		}
	})

	t.Run("nil doc returns false", func(t *testing.T) {
		op := &openapi3.Operation{}
		if opRequiresAuth(op, nil) {
			t.Errorf("expected false for nil doc")
		}
	})
}
