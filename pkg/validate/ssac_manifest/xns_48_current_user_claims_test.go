//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what XNS-48 test — currentUser 사용 시 manifest backend.auth.claims 필요 검증

package ssac_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestXns48CurrentUserClaims(t *testing.T) {
	// nil fs -> nil.
	if d := xns48CurrentUserClaims(nil); d != nil {
		t.Errorf("nil fs -> nil, got %v", d)
	}

	// Does not use currentUser -> nil.
	noUse := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})
	if d := xns48CurrentUserClaims(noUse); d != nil {
		t.Errorf("no currentUser use -> nil, got %v", d)
	}

	// Uses @auth and claims enabled in Ground -> nil.
	authFn := ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "auth"}}}
	enabled := fsWithFuncs(authFn)
	enabled.SetGround(&rule.Ground{Config: map[string]bool{"backend.auth.claims": true}})
	if d := xns48CurrentUserClaims(enabled); d != nil {
		t.Errorf("claims enabled -> nil, got %v", d)
	}

	// Uses currentUser.<field> input but claims not declared -> ERROR.
	inputFn := ssac.ServiceFunc{Sequences: []ssac.Sequence{
		{Type: "call", Inputs: map[string]string{"id": "currentUser.id"}},
	}}
	missing := fsWithFuncs(inputFn)
	d := xns48CurrentUserClaims(missing)
	if len(d) != 1 || !strings.Contains(d[0].Message, "[XNS-48]") {
		t.Fatalf("expected XNS-48 diag, got %v", d)
	}
}
