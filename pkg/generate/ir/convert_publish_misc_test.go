//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what convertPublish/resolveExposeInternal/isCountResultType/ddlTableSingularIR/DDLTableSingularIR/findDDLTable

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestConvertPublish(t *testing.T) {
	op := convertPublish(ssac.Sequence{
		Topic:   "order.completed",
		Inputs:  map[string]string{"ID": "order.ID"},
		Options: map[string]string{"delay": "1800"},
	})
	if op.Kind != OpPublish || op.Publish == nil {
		t.Fatalf("op = %+v", op)
	}
	if op.Publish.Topic != "order.completed" {
		t.Errorf("topic = %q", op.Publish.Topic)
	}
	if len(op.Publish.Payload) != 1 || len(op.Publish.Options) != 1 {
		t.Errorf("payload=%v options=%v", op.Publish.Payload, op.Publish.Options)
	}
}

func TestResolveExposeInternal(t *testing.T) {
	if resolveExposeInternal(nil) {
		t.Errorf("nil should be false")
	}
	if resolveExposeInternal(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}) {
		t.Errorf("no error config should be false")
	}
	tru := true
	on := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{
		Error: &manifest.ErrorConfig{ExposeInternalError: &tru},
	}}}
	if !resolveExposeInternal(on) {
		t.Errorf("explicit true should be true")
	}
}

func TestIsCountResultType(t *testing.T) {
	for _, ty := range []string{"int64", "int32", "int", "uint64"} {
		if !isCountResultType(ty) {
			t.Errorf("%q should be count type", ty)
		}
	}
	for _, ty := range []string{"string", "bool", "float64", "User"} {
		if isCountResultType(ty) {
			t.Errorf("%q should not be count type", ty)
		}
	}
}

func TestDDLTableSingularIR(t *testing.T) {
	// exported and unexported delegate to caseconv.TableSingular
	if got := DDLTableSingularIR("users"); got != ddlTableSingularIR("users") {
		t.Errorf("exported/unexported mismatch")
	}
	if DDLTableSingularIR("users") == "" {
		t.Errorf("expected non-empty singular")
	}
}

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
