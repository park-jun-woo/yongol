//ff:func feature=gen-ir type=test control=sequence
//ff:what TestSplitPaginationArgs -- FieldArg를 where-clause/pagination 인자로 분리(PascalCase→snake 매칭) 검증

package ir

import "testing"

func TestSplitPaginationArgs(t *testing.T) {
	args := []FieldArg{
		{Key: "OwnerId"}, // where
		{Key: "PerPage"}, // pagination (per_page)
		{Key: "Status"},  // where
		{Key: "Cursor"},  // pagination
	}

	whereArgs, pagArgs := splitPaginationArgs(args)

	if len(whereArgs) != 2 {
		t.Fatalf("whereArgs len = %d, want 2 (%v)", len(whereArgs), whereArgs)
	}
	if whereArgs[0].Key != "OwnerId" || whereArgs[1].Key != "Status" {
		t.Errorf("whereArgs = %v, want [OwnerId Status]", whereArgs)
	}
	if len(pagArgs) != 2 {
		t.Fatalf("pagArgs len = %d, want 2 (%v)", len(pagArgs), pagArgs)
	}
	if pagArgs[0].Key != "PerPage" || pagArgs[1].Key != "Cursor" {
		t.Errorf("pagArgs = %v, want [PerPage Cursor]", pagArgs)
	}
}
