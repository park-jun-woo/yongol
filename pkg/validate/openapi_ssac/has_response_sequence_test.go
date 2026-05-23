//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what hasResponseSequence — empty/response 유무 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestHasResponseSequence(t *testing.T) {
	t.Run("no sequences returns false", func(t *testing.T) {
		fn := ssac.ServiceFunc{}
		if hasResponseSequence(fn) {
			t.Error("expected false")
		}
	})

	t.Run("no response type returns false", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Sequences: []ssac.Sequence{{Type: "get"}, {Type: "empty"}},
		}
		if hasResponseSequence(fn) {
			t.Error("expected false")
		}
	})

	t.Run("has response returns true", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Sequences: []ssac.Sequence{{Type: "get"}, {Type: "response"}},
		}
		if !hasResponseSequence(fn) {
			t.Error("expected true")
		}
	})
}
