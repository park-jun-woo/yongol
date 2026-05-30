//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestSSaCFuncHelpers — unit tests for the pure ssac_func helper functions
package ssac_func

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestItoa(t *testing.T) {
	if got := itoa(42); got != "42" {
		t.Errorf("itoa(42) = %q", got)
	}
	if got := itoa(0); got != "0" {
		t.Errorf("itoa(0) = %q", got)
	}
}

func TestUcFirst(t *testing.T) {
	tests := []struct{ in, want string }{
		{"hash", "Hash"},
		{"Hash", "Hash"},
		{"", ""},
		{"1abc", "1abc"},
	}
	for _, tt := range tests {
		if got := ucFirst(tt.in); got != tt.want {
			t.Errorf("ucFirst(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestToCamelKey(t *testing.T) {
	tests := []struct{ in, want string }{
		{"HashPassword", "hashPassword"},
		{"hashPassword", "hashPassword"},
		{"", ""},
		{"X", "x"},
	}
	for _, tt := range tests {
		if got := toCamelKey(tt.in); got != tt.want {
			t.Errorf("toCamelKey(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestIsIntType(t *testing.T) {
	for _, ty := range []string{"int", "int8", "int16", "int32", "int64", "uint", "uint64", "byte", "rune"} {
		if !isIntType(ty) {
			t.Errorf("%q should be int type", ty)
		}
	}
	for _, ty := range []string{"string", "float64", "bool", ""} {
		if isIntType(ty) {
			t.Errorf("%q should not be int type", ty)
		}
	}
}

func TestJoinReturnTypes(t *testing.T) {
	if got := joinReturnTypes([]string{"int", "error"}); got != "int, error" {
		t.Errorf("got %q", got)
	}
	if got := joinReturnTypes([]string{"string"}); got != "string" {
		t.Errorf("got %q", got)
	}
	if got := joinReturnTypes(nil); got != "" {
		t.Errorf("got %q", got)
	}
}

func TestInferLiteralType(t *testing.T) {
	tests := []struct{ in, want string }{
		{`"foo"`, "string"},
		{"42", "int64"},
		{"3.14", "float64"},
		{"true", "bool"},
		{"false", "bool"},
		{"nil", "nil"},
		{"someVar", ""},
		{"course.ID", ""},
		{"", ""},
		{`  "spaced"  `, "string"},
	}
	for _, tt := range tests {
		if got := inferLiteralType(tt.in); got != tt.want {
			t.Errorf("inferLiteralType(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCallFuncName(t *testing.T) {
	tests := []struct{ in, want string }{
		{"auth.VerifyPassword", "VerifyPassword"},
		{"plainname", ""},
		{".trailing", ""}, // idx <= 0
		{"pkg.", ""},      // idx >= len-1
	}
	for _, tt := range tests {
		if got := callFuncName(tt.in); got != tt.want {
			t.Errorf("callFuncName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCallFuncCamelName(t *testing.T) {
	if got := callFuncCamelName("billing.CheckCredits"); got != "checkCredits" {
		t.Errorf("got %q, want checkCredits", got)
	}
	if got := callFuncCamelName("noqualifier"); got != "" {
		t.Errorf("got %q, want empty", got)
	}
}

func TestNormalizedCallKey(t *testing.T) {
	if got := normalizedCallKey("billing.DeductCredit"); got != "billing.deductCredit" {
		t.Errorf("got %q, want billing.deductCredit", got)
	}
	// no dot → unchanged.
	if got := normalizedCallKey("plain"); got != "plain" {
		t.Errorf("got %q, want plain", got)
	}
}

func TestBuiltinFuncNames(t *testing.T) {
	specs := []funcspec.FuncSpec{
		{Package: "auth", Name: "hashPassword"},
		{Package: "auth", Name: "verifyPassword"},
		{Package: "cache", Name: "get"},
	}
	got := builtinFuncNames("auth", specs)
	want := []string{"HashPassword", "VerifyPassword"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("builtinFuncNames(auth) = %v, want %v", got, want)
	}
	if got := builtinFuncNames("missing", specs); got != nil {
		t.Errorf("missing pkg → %v, want nil", got)
	}
}
