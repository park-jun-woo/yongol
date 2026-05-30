//ff:func feature=validate type=test control=selection topic=states
//ff:what TestXsm27StateIntentDeclarationNilGuards — xsm27StateIntentDeclaration 가드 early-return 분기 검증

package ssac_statemachine

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXsm27StateIntentDeclarationNilGuards(t *testing.T) {
	t.Run("nil fullstack", func(t *testing.T) {
		if d := xsm27StateIntentDeclaration(nil); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("nil openapi doc", func(t *testing.T) {
		if d := xsm27StateIntentDeclaration(&yongol.Fullstack{}); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("nil ground", func(t *testing.T) {
		fs := &yongol.Fullstack{OpenAPIDoc: &openapi3.T{Paths: openapi3.NewPaths()}}
		// no SetGround call -> Ground() returns nil
		if d := xsm27StateIntentDeclaration(fs); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})
}
