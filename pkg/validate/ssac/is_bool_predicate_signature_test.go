//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what isBoolPredicateSignature — bool predicate 시그니처 검증 (true/false/multi/empty)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestIsBoolPredicateSignature(t *testing.T) {
	t.Run("single bool returns true", func(t *testing.T) {
		spec := &funcspec.FuncSpec{ReturnTypes: []string{"bool"}}
		if !isBoolPredicateSignature(spec) {
			t.Error("expected true")
		}
	})

	t.Run("single non-bool returns false", func(t *testing.T) {
		spec := &funcspec.FuncSpec{ReturnTypes: []string{"error"}}
		if isBoolPredicateSignature(spec) {
			t.Error("expected false")
		}
	})

	t.Run("multiple returns false", func(t *testing.T) {
		spec := &funcspec.FuncSpec{ReturnTypes: []string{"bool", "error"}}
		if isBoolPredicateSignature(spec) {
			t.Error("expected false")
		}
	})

	t.Run("empty returns false", func(t *testing.T) {
		spec := &funcspec.FuncSpec{}
		if isBoolPredicateSignature(spec) {
			t.Error("expected false")
		}
	})
}
