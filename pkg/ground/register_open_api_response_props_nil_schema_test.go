//ff:func feature=rule type=test control=sequence
//ff:what registerOpenAPIResponseProps test — 2xx schema properties → OpenAPI.response.<op>.<field> 등록 (nil guard·맥락·skip)
package ground

import (
	"testing"
)

func TestRegisterOpenAPIResponseProps_NilSchema(t *testing.T) {
	g := newGround()
	registerOpenAPIResponseProps(g, "Op", nil) // must not panic
	if len(g.Types) != 0 {
		t.Errorf("nil schema registered %d types, want 0", len(g.Types))
	}
}
