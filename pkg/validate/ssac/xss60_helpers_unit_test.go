//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what TestSSaCXss60Helpers — unit tests for the pure ssac validate helper functions
package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXss60IsIntLiteral(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"0", true},
		{"42", true},
		{"-7", true},
		{"", false},
		{"-", false},
		{"3.14", false},
		{"abc", false},
		{"12a", false},
	}
	for _, tt := range tests {
		if got := xss60IsIntLiteral(tt.in); got != tt.want {
			t.Errorf("xss60IsIntLiteral(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestXss60ModelToTableName(t *testing.T) {
	tests := []struct{ in, want string }{
		{"User", "users"},
		{"RefreshToken", "refresh_tokens"},
		{"AuditLog", "audit_logs"},
	}
	for _, tt := range tests {
		if got := xss60ModelToTableName(tt.in); got != tt.want {
			t.Errorf("xss60ModelToTableName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestExpectedQueryFile(t *testing.T) {
	tests := []struct{ in, want string }{
		{"RefreshToken", "db/queries/refresh_tokens.sql"},
		{"User", "db/queries/users.sql"},
	}
	for _, tt := range tests {
		if got := expectedQueryFile(tt.in); got != tt.want {
			t.Errorf("expectedQueryFile(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestXss60InferType(t *testing.T) {
	fn := parsessac.ServiceFunc{}
	tests := []struct{ expr, want string }{
		{`"completed"`, "string"},
		{"0", "int64"},
		{"42", "int64"},
		{"", ""},
		{"unknownvar", ""}, // no dot, not literal
	}
	for _, tt := range tests {
		if got := xss60InferType(tt.expr, fn, nil); got != tt.want {
			t.Errorf("xss60InferType(%q) = %q, want %q", tt.expr, got, tt.want)
		}
	}
}

func TestXss60CompareFieldType(t *testing.T) {
	fn := parsessac.ServiceFunc{
		FileName:  "order.ssac",
		Line:      10,
		Subscribe: &parsessac.SubscribeInfo{Topic: "order.completed"},
	}
	// Type mismatch → diagnostic returned.
	pub := map[string]string{"OrderID": "int64"}
	diag, ok := xss60CompareFieldType(fn, parsessac.StructField{Name: "OrderID", Type: "string"}, pub)
	if !ok {
		t.Fatal("expected mismatch diagnostic")
	}
	if diag.Line != 10 || diag.File != "order.ssac" {
		t.Errorf("diag location = %s:%d", diag.File, diag.Line)
	}

	// Matching type → no diagnostic.
	if _, ok := xss60CompareFieldType(fn, parsessac.StructField{Name: "OrderID", Type: "int64"}, pub); ok {
		t.Error("matching type should not produce a diagnostic")
	}
	// Field not in publish map → no diagnostic.
	if _, ok := xss60CompareFieldType(fn, parsessac.StructField{Name: "Other", Type: "string"}, pub); ok {
		t.Error("absent field should not produce a diagnostic")
	}
	// Empty publish type → no diagnostic.
	if _, ok := xss60CompareFieldType(fn, parsessac.StructField{Name: "Empty", Type: "string"}, map[string]string{"Empty": ""}); ok {
		t.Error("empty publish type should not produce a diagnostic")
	}
}
