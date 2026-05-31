//ff:func feature=sqlc-parse type=test control=sequence
//ff:what TestByName_ZeroCov — sqlc 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속

package sqlc

import "testing"

func TestByNameSQLCHelpers_ZeroCov(t *testing.T) {
	// rowTypeFor across cardinalities.
	if rowTypeFor("ListItems", "many") != "ListItemsRow" {
		t.Errorf("rowTypeFor many")
	}
	if rowTypeFor("GetItem", "one") != "GetItemRow" {
		t.Errorf("rowTypeFor one")
	}
	if rowTypeFor("DeleteItem", "exec") != "" {
		t.Errorf("rowTypeFor exec should be empty")
	}

	// parseOneSelectColumn across branches.
	cases := map[string]string{
		"":                "",
		"*":               "",
		"u.name":          "name",
		"name":            "name",
		`"Name"`:          "Name",
		"COUNT(*) AS cnt": "cnt",
		"a + b":           "",
	}
	for in, want := range cases {
		if got := parseOneSelectColumn(in); got != want {
			t.Errorf("parseOneSelectColumn(%q) = %q, want %q", in, got, want)
		}
	}

	// finalizeQuerySpec.
	spec := &QuerySpec{Body: "SELECT id, name FROM users  \n"}
	finalizeQuerySpec(spec, map[string]bool{"status": true})
	if len(spec.Params) != 1 || spec.Params[0] != "status" {
		t.Errorf("finalizeQuerySpec Params = %v", spec.Params)
	}

	// processSQLCScanLine: a name comment line then a body line.
	var specs []QuerySpec
	paramSet := map[string]bool{}
	var current *QuerySpec
	current, paramSet = processSQLCScanLine(
		"-- name: ListUsers :many", "User", "queries.sql", 1, current, paramSet, &specs)
	current, paramSet = processSQLCScanLine(
		"SELECT id FROM users WHERE status = @status;", "User", "queries.sql", 2, current, paramSet, &specs)
	if current == nil {
		t.Errorf("processSQLCScanLine current is nil after scan")
	}
}
