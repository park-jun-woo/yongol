//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"reflect"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestPascalCase(t *testing.T) {
	tests := []struct{ in, want string }{
		{"", ""},
		{"user_id", "UserId"},
		{"user", "User"},
		{"User", "User"},
		{"userName", "UserName"},
		{"audit_log_entry", "AuditLogEntry"},
		{"__leading", "Leading"},
	}
	for _, tt := range tests {
		if got := pascalCase(tt.in); got != tt.want {
			t.Errorf("pascalCase(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestPascalCaseSnake(t *testing.T) {
	tests := []struct{ in, want string }{
		{"a_b_c", "ABC"},
		{"user_id", "UserId"},
		{"_x_", "X"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := pascalCaseSnake(tt.in); got != tt.want {
			t.Errorf("pascalCaseSnake(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestToSnake(t *testing.T) {
	tests := []struct{ in, want string }{
		{"UserId", "user_id"},
		{"AuditLog", "audit_log"},
		{"User", "user"},
	}
	for _, tt := range tests {
		if got := toSnake(tt.in); got != tt.want {
			t.Errorf("toSnake(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestIndexCaseFold(t *testing.T) {
	if got := indexCaseFold("id as user_id", " AS "); got != 2 {
		t.Errorf("indexCaseFold lowercase AS = %d, want 2", got)
	}
	if got := indexCaseFold("id AS x", " as "); got != 2 {
		t.Errorf("indexCaseFold mixed = %d, want 2", got)
	}
	if got := indexCaseFold("nothing", "zzz"); got != -1 {
		t.Errorf("indexCaseFold no match = %d, want -1", got)
	}
}

func TestStripSQLLineComments(t *testing.T) {
	in := "SELECT id -- the id\nFROM users; -- partial\n"
	got := stripSQLLineComments(in)
	want := "SELECT id \nFROM users; \n"
	if got != want {
		t.Errorf("stripSQLLineComments = %q, want %q", got, want)
	}
	// No comments → unchanged.
	if got := stripSQLLineComments("SELECT 1"); got != "SELECT 1" {
		t.Errorf("no-comment passthrough = %q", got)
	}
}

func TestSequenceTag(t *testing.T) {
	if got := sequenceTag("call"); got != "call" {
		t.Errorf("sequenceTag = %q, want call", got)
	}
}

func TestCastToIntWidth(t *testing.T) {
	tests := []struct{ in, want string }{
		{"bigint", "int64"},
		{"int8", "int64"},
		{"int", "int32"},
		{"int4", "int32"},
		{"integer", "int32"},
		{"", "int32"},
		{"text", ""},
		{"numeric", ""},
	}
	for _, tt := range tests {
		if got := castToIntWidth(tt.in); got != tt.want {
			t.Errorf("castToIntWidth(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestWidthToPGCast(t *testing.T) {
	if got := widthToPGCast("int64"); got != "bigint" {
		t.Errorf("widthToPGCast(int64) = %q, want bigint", got)
	}
	if got := widthToPGCast("int32"); got != "int" {
		t.Errorf("widthToPGCast(int32) = %q, want int", got)
	}
	if got := widthToPGCast(""); got != "int" {
		t.Errorf("widthToPGCast(empty) = %q, want int", got)
	}
}

func TestContainsUsedBy(t *testing.T) {
	usedBy := []string{"GET", "POST"}
	if !containsUsedBy(usedBy, "POST") {
		t.Error("expected POST present")
	}
	if containsUsedBy(usedBy, "DELETE") {
		t.Error("expected DELETE absent")
	}
	if containsUsedBy(nil, "GET") {
		t.Error("nil slice should not contain anything")
	}
}

func TestSelectColsContain(t *testing.T) {
	cols := []string{"id", "email"}
	if !selectColsContain(cols, "email") {
		t.Error("expected email present")
	}
	if selectColsContain(cols, "name") {
		t.Error("expected name absent")
	}
}

func TestSplitReturningColumns(t *testing.T) {
	tests := []struct {
		name   string
		clause string
		want   map[string]bool
	}{
		{"plain", "id, email", map[string]bool{"id": true, "email": true}},
		{"alias prefix", "u.id, u.email", map[string]bool{"id": true, "email": true}},
		{"AS alias", "id AS user_id, email", map[string]bool{"user_id": true, "email": true}},
		{"uppercase folded", "ID, EMAIL", map[string]bool{"id": true, "email": true}},
		{"quoted ident", "\"Name\"", map[string]bool{"name": true}},
		{"empty entries skipped", "id, ,email", map[string]bool{"id": true, "email": true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitReturningColumns(tt.clause)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitReturningColumns(%q) = %v, want %v", tt.clause, got, tt.want)
			}
		})
	}
}

func TestModelToTableName(t *testing.T) {
	tests := []struct{ in, want string }{
		{"User", "users"},
		{"AuditLog", "audit_logs"},
		{"Workflow", "workflows"},
	}
	for _, tt := range tests {
		if got := modelToTableName(tt.in); got != tt.want {
			t.Errorf("modelToTableName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestFindCaseInsensitiveParam(t *testing.T) {
	params := rule.StringSet{"UserID": true, "email": true}
	if got := findCaseInsensitiveParam("userid", params); got != "UserID" {
		t.Errorf("findCaseInsensitiveParam(userid) = %q, want UserID", got)
	}
	if got := findCaseInsensitiveParam("EMAIL", params); got != "email" {
		t.Errorf("findCaseInsensitiveParam(EMAIL) = %q, want email", got)
	}
	if got := findCaseInsensitiveParam("missing", params); got != "" {
		t.Errorf("findCaseInsensitiveParam(missing) = %q, want empty", got)
	}
}

func TestExtractReturningClause(t *testing.T) {
	tests := []struct{ body, want string }{
		{"INSERT INTO users VALUES (1) RETURNING *;", "*"},
		{"INSERT INTO users VALUES (1) RETURNING id, email", "id, email"},
		{"SELECT * FROM users WHERE id = @id;", ""},
		{"", ""},
		{"INSERT ... RETURNING id; -- partial", "id"},
	}
	for _, tt := range tests {
		if got := extractReturningClause(tt.body); got != tt.want {
			t.Errorf("extractReturningClause(%q) = %q, want %q", tt.body, got, tt.want)
		}
	}
}

func ddlTable(cols ...string) *ddl.Table {
	m := make(map[string]ddl.Column, len(cols))
	for _, c := range cols {
		m[c] = ddl.Column{Name: c}
	}
	return &ddl.Table{Columns: m}
}

func TestReturningCoversAllColumns(t *testing.T) {
	tbl := ddlTable("id", "email")
	full := map[string]bool{"id": true, "email": true}
	if !returningCoversAllColumns(full, tbl.Columns) {
		t.Error("expected full coverage")
	}
	partial := map[string]bool{"id": true}
	if returningCoversAllColumns(partial, tbl.Columns) {
		t.Error("expected partial NOT to cover all")
	}
}

func TestClassifyReturningShape(t *testing.T) {
	tbl := ddlTable("id", "email")
	tests := []struct {
		name   string
		clause string
		table  *ddl.Table
		want   ReturningShape
	}{
		{"empty", "", tbl, ShapeNone},
		{"star", "*", tbl, ShapeFull},
		{"covers all", "id, email", tbl, ShapeFull},
		{"partial subset", "id", tbl, ShapePartial},
		{"nil table non-empty", "id", nil, ShapePartial},
		{"empty-column table", "id", &ddl.Table{}, ShapePartial},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := classifyReturningShape(tt.clause, tt.table); got != tt.want {
				t.Errorf("classifyReturningShape(%q) = %q, want %q", tt.clause, got, tt.want)
			}
		})
	}
}

func TestResolveOAPITypeWithFormat(t *testing.T) {
	tests := []struct{ base, format, want string }{
		{"integer", "int64", "int64"},
		{"integer", "int32", "int32"},
		{"number", "double", "float64"},
		{"number", "float", "float32"},
		{"string", "uuid", "uuid"},
		{"string", "", "string"},         // empty format → base
		{"integer", "weird", "integer"},  // unknown format → base
		{"object", "anything", "object"}, // unknown base → base
	}
	for _, tt := range tests {
		if got := resolveOAPITypeWithFormat(tt.base, tt.format); got != tt.want {
			t.Errorf("resolveOAPITypeWithFormat(%q,%q) = %q, want %q", tt.base, tt.format, got, tt.want)
		}
	}
}

func TestExpectedSsacReturnType(t *testing.T) {
	if got := expectedSsacReturnType(ShapePartial, "User", "GetUser"); got != "GetUserRow" {
		t.Errorf("partial: %q, want GetUserRow", got)
	}
	if got := expectedSsacReturnType(ShapeFull, "User", "GetUser"); got != "User" {
		t.Errorf("full: %q, want User", got)
	}
	if got := expectedSsacReturnType(ShapeNone, "User", "GetUser"); got != "User" {
		t.Errorf("none: %q, want User", got)
	}
	if got := expectedSsacReturnType(ShapeFull, "", "GetUser"); got != "" {
		t.Errorf("empty model: %q, want empty", got)
	}
}

func TestResolveBuiltinCall(t *testing.T) {
	p, m := resolveBuiltinCall(ssacparser.Sequence{Type: "call", Package: "session", Model: "Get"}, false)
	if p != "session" || m != "Get" {
		t.Errorf("call = (%q,%q), want (session,Get)", p, m)
	}
	p, m = resolveBuiltinCall(ssacparser.Sequence{Type: "call", Package: "", Model: "Get"}, false)
	if p != "" || m != "" {
		t.Errorf("call missing pkg = (%q,%q), want empties", p, m)
	}
	p, m = resolveBuiltinCall(ssacparser.Sequence{Type: "publish"}, false)
	if p != "queue" || m != "Publish" {
		t.Errorf("publish = (%q,%q), want (queue,Publish)", p, m)
	}
	p, m = resolveBuiltinCall(ssacparser.Sequence{Type: "get"}, false)
	if p != "" || m != "" {
		t.Errorf("get = (%q,%q), want empties", p, m)
	}
}

func TestResolveQueryName(t *testing.T) {
	if got := resolveQueryName(ssacparser.Sequence{Model: "Workflow.FindByID"}); got != "WorkflowFindByID" {
		t.Errorf("got %q, want WorkflowFindByID", got)
	}
	if got := resolveQueryName(ssacparser.Sequence{Model: "Get"}); got != "Get" {
		t.Errorf("got %q, want Get", got)
	}
}

func TestXqs73EligibleSeqType(t *testing.T) {
	for _, ty := range []string{"get", "post", "put"} {
		if !xqs73EligibleSeqType(ty) {
			t.Errorf("%q should be eligible", ty)
		}
	}
	for _, ty := range []string{"delete", "call", "empty"} {
		if xqs73EligibleSeqType(ty) {
			t.Errorf("%q should not be eligible", ty)
		}
	}
}

func TestXqs20EligibleSeqType(t *testing.T) {
	for _, ty := range []string{ssacparser.SeqGet, ssacparser.SeqPost, ssacparser.SeqPut} {
		if !xqs20EligibleSeqType(ty) {
			t.Errorf("%q should be eligible", ty)
		}
	}
	if xqs20EligibleSeqType(ssacparser.SeqDelete) {
		t.Error("delete should not be eligible")
	}
}

func TestBuildXqs19Advice(t *testing.T) {
	got := buildXqs19Advice("session", "GetUser")
	if !strings.Contains(got, "GetUser") || !strings.Contains(got, "specs/db/queries/session.sql") {
		t.Errorf("advice missing parts: %q", got)
	}
}

func TestBuildXqs20Advice(t *testing.T) {
	partial := buildXqs20Advice("get", "Model", "GetRow", ShapePartial)
	if !strings.Contains(partial, "expand RETURNING") {
		t.Errorf("partial advice: %q", partial)
	}
	full := buildXqs20Advice("get", "Row", "Model", ShapeFull)
	if !strings.Contains(full, "restrict RETURNING") {
		t.Errorf("full advice: %q", full)
	}
	none := buildXqs20Advice("get", "Row", "Model", ShapeNone)
	if !strings.Contains(none, "@get Model") {
		t.Errorf("none advice: %q", none)
	}
}

func TestFormatReturningReason(t *testing.T) {
	if got := formatReturningReason(ShapeFull, "*"); got != "RETURNING * → model" {
		t.Errorf("full star: %q", got)
	}
	if got := formatReturningReason(ShapeFull, "id, email"); got != "RETURNING <full column list> → model" {
		t.Errorf("full list: %q", got)
	}
	if got := formatReturningReason(ShapePartial, "id"); got != "RETURNING id → partial Row" {
		t.Errorf("partial: %q", got)
	}
	if got := formatReturningReason(ShapeNone, ""); got != "no RETURNING → model" {
		t.Errorf("none: %q", got)
	}
}
