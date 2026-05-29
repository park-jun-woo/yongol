//ff:func feature=rule type=test control=sequence
//ff:what Ground — Lookup/Types/Pairs/Config/Vars/Flags/Schemas 맵 기본 조회 동작 검증

package rule

import "testing"

func TestGround_FieldsLookup(t *testing.T) {
	g := &Ground{
		Lookup:  map[string]StringSet{"target.kind": {"a": true}},
		Types:   map[string]string{"target.kind.name": "string"},
		Pairs:   map[string]StringSet{"target.pairKind": {"k:v": true}},
		Config:  map[string]bool{"feature.on": true},
		Vars:    StringSet{"user_id": true},
		Flags:   StringSet{"archived": true},
		Schemas: map[string][]string{"target.schema": {"id", "name"}},
	}
	if !g.Lookup["target.kind"]["a"] {
		t.Errorf("Ground.Lookup miss")
	}
	if g.Types["target.kind.name"] != "string" {
		t.Errorf("Ground.Types miss")
	}
	if !g.Pairs["target.pairKind"]["k:v"] {
		t.Errorf("Ground.Pairs miss")
	}
	if !g.Config["feature.on"] {
		t.Errorf("Ground.Config miss")
	}
	if !g.Vars["user_id"] {
		t.Errorf("Ground.Vars miss")
	}
	if !g.Flags["archived"] {
		t.Errorf("Ground.Flags miss")
	}
	if got := g.Schemas["target.schema"]; len(got) != 2 || got[0] != "id" || got[1] != "name" {
		t.Errorf("Ground.Schemas = %v; want [id name]", got)
	}
}
