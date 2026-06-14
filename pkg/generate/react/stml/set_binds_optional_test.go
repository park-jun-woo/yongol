//ff:func feature=stml-gen type=test control=sequence
//ff:what setBindsOptional — required 미포함 route 바인드만 Optional 표시, required/비route는 유지 검증

package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSetBindsOptional(t *testing.T) {
	binds := []stmlparser.ParamBind{
		{Name: "BuildingID", Source: "route.BuildingID"}, // required → false
		{Name: "RoomID", Source: "route.RoomID"},         // not required → true
		{Name: "Q", Source: "query.q"},                   // non-route → false
	}
	required := map[string]bool{"BuildingID": true}

	setBindsOptional(binds, required)

	if binds[0].Optional {
		t.Errorf("required BuildingID should stay Optional=false")
	}
	if !binds[1].Optional {
		t.Errorf("unrequired RoomID should be Optional=true")
	}
	if binds[2].Optional {
		t.Errorf("non-route query.q should stay Optional=false")
	}

	// no route binds → nothing flagged
	plain := []stmlparser.ParamBind{{Name: "Q", Source: "query.q"}}
	setBindsOptional(plain, map[string]bool{})
	if plain[0].Optional {
		t.Errorf("non-route bind must not be flagged")
	}
}
