//ff:func feature=gen-ir type=test control=sequence
//ff:what TestSetSourceColumns -- Field 접근자(.Prefix 제거 후 snake_case) → SourceColumn 파생, 빈 Field 무변경 검증

package ir

import "testing"

func TestSetSourceColumns(t *testing.T) {
	args := []FieldArg{
		{Field: ".OrgID"}, // -> org_id
		{Field: ".Status"},
		{Field: ""}, // empty -> untouched
	}
	setSourceColumns(args)

	if args[0].SourceColumn != "org_id" {
		t.Errorf("args[0].SourceColumn = %q, want org_id", args[0].SourceColumn)
	}
	if args[1].SourceColumn != "status" {
		t.Errorf("args[1].SourceColumn = %q, want status", args[1].SourceColumn)
	}
	if args[2].SourceColumn != "" {
		t.Errorf("args[2].SourceColumn = %q, want empty for blank Field", args[2].SourceColumn)
	}
}
