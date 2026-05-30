//ff:func feature=agent type=test control=sequence
//ff:what TestRecordPathOps — path 블록(직접/래핑)에서 pathToOps·opToPath 매핑 기록 검증

package agent

import "testing"

func TestRecordPathOps(t *testing.T) {
	// Direct map form: top-level key is the path.
	pathToOps := map[string][]string{}
	opToPath := map[string]string{}
	recordPathOps("/users:\n  get:\n    operationId: ListUsers", "ListUsers", pathToOps, opToPath)
	if opToPath["ListUsers"] != "/users" {
		t.Errorf("opToPath = %v, want /users", opToPath)
	}
	if len(pathToOps["/users"]) != 1 || pathToOps["/users"][0] != "ListUsers" {
		t.Errorf("pathToOps = %v, want [ListUsers]", pathToOps)
	}

	// Adding another op to same path appends uniquely.
	recordPathOps("/users:\n  post:\n    operationId: CreateUser", "CreateUser", pathToOps, opToPath)
	if len(pathToOps["/users"]) != 2 {
		t.Errorf("pathToOps[/users] = %v, want 2 entries", pathToOps["/users"])
	}
	if opToPath["CreateUser"] != "/users" {
		t.Errorf("opToPath[CreateUser] = %q, want /users", opToPath["CreateUser"])
	}

	// Content that is a YAML sequence fails the direct map unmarshal, so the
	// wrapped ("paths:\n  ...") fallback path is entered. The wrapped Paths is
	// itself a sequence, so no mapping is recorded — but the fallback branch runs.
	p2 := map[string][]string{}
	o2 := map[string]string{}
	recordPathOps("- item1\n- item2", "SeqOp", p2, o2)
	if len(p2) != 0 || len(o2) != 0 {
		t.Errorf("sequence content should record no mappings, got %v %v", p2, o2)
	}
}
