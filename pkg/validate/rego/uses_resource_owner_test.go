//ff:func feature=validate type=test control=sequence topic=rego-structural
//ff:what usesResourceOwner — UsesOwner 존재 여부 검증 (true/false/empty)

package rego

import (
	"testing"

	regoparser "github.com/park-jun-woo/yongol/pkg/parser/rego"
)

func TestUsesResourceOwner(t *testing.T) {
	t.Run("empty rules returns false", func(t *testing.T) {
		if usesResourceOwner(nil) {
			t.Error("expected false")
		}
	})

	t.Run("no owner rules returns false", func(t *testing.T) {
		rules := []regoparser.AllowRule{{UsesOwner: false}}
		if usesResourceOwner(rules) {
			t.Error("expected false")
		}
	})

	t.Run("has owner rule returns true", func(t *testing.T) {
		rules := []regoparser.AllowRule{{UsesOwner: false}, {UsesOwner: true}}
		if !usesResourceOwner(rules) {
			t.Error("expected true")
		}
	})
}
