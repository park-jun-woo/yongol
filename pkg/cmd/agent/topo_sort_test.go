//ff:func feature=agent type=test control=iteration dimension=2
//ff:what TestTopoSortTables — belongs_to 위상 정렬(부모 우선)과 순환 시 깨고 반환 검증

package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func indexOf(s []string, v string) int {
	for i, x := range s {
		if x == v {
			return i
		}
	}
	return -1
}

func TestTopoSortTables(t *testing.T) {
	// orders belongs_to users -> users must come before orders.
	tables := map[string]features.TableDef{
		"users":  {},
		"orders": {BelongsTo: []string{"users"}},
	}
	got := topoSortTables(tables)
	if indexOf(got, "users") > indexOf(got, "orders") {
		t.Errorf("users must precede orders: %v", got)
	}
	if len(got) != 2 {
		t.Errorf("expected 2 tables, got %v", got)
	}

	// Cycle: a<->b. Must still return both (cycle broken).
	cyc := map[string]features.TableDef{
		"a": {BelongsTo: []string{"b"}},
		"b": {BelongsTo: []string{"a"}},
	}
	gotCyc := topoSortTables(cyc)
	if len(gotCyc) != 2 {
		t.Errorf("cycle: expected 2 tables returned, got %v", gotCyc)
	}

	// belongs_to references a table not in the map -> that edge is skipped,
	// so orders has inDegree 0 and is returned normally.
	dangling := map[string]features.TableDef{
		"orders": {BelongsTo: []string{"ghost"}},
	}
	gotDangling := topoSortTables(dangling)
	if len(gotDangling) != 1 || gotDangling[0] != "orders" {
		t.Errorf("dangling ref: expected [orders], got %v", gotDangling)
	}

	// Empty map -> empty result.
	if got := topoSortTables(map[string]features.TableDef{}); len(got) != 0 {
		t.Errorf("empty: expected no tables, got %v", got)
	}
}
