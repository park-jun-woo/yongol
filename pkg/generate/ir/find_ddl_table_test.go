//ff:func feature=gen-ir type=test control=sequence
//ff:what convertPublish/resolveExposeInternal/isCountResultType/ddlTableSingularIR/DDLTableSingularIR/findDDLTable
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestFindDDLTable(t *testing.T) {
	tables := []ddl.Table{{Name: "users"}, {Name: "courses"}}
	// model name "User" -> singular match against "users"
	if tb := findDDLTable(tables, "User"); tb == nil || tb.Name != "users" {
		t.Errorf("findDDLTable(User) = %v", tb)
	}
	if tb := findDDLTable(tables, "Course"); tb == nil || tb.Name != "courses" {
		t.Errorf("findDDLTable(Course) = %v", tb)
	}
	if tb := findDDLTable(tables, "Nonexistent"); tb != nil {
		t.Errorf("expected nil for unknown model, got %v", tb)
	}
}
