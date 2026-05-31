//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestParseInputValue -- parseInputValue 숫자/따옴표/변수/dotted 리터럴 분류 검증
package ir

import (
	"testing"
)

func TestParseInputValueNumeric(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		value   string
		wantLit string
		wantSrc string
	}{
		{"Integer", "CreditsSpent", "1", "1", ""},
		{"NegativeInt", "Rate", "-3", "-3", ""},
		{"Decimal", "Amount", "1.5", "1.5", ""},
		{"NegativeDecimal", "Discount", "-0.25", "-0.25", ""},
		{"Zero", "Count", "0", "0", ""},
		{"VarNotNumeric", "Status", "wf", "", "wf"},
		{"QuotedString", "Status", `"active"`, "active", ""},
		{"DottedRef", "ID", "wf.OrgID", "", "wf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fa := parseInputValue(tt.key, tt.value)
			if fa.Literal != tt.wantLit {
				t.Errorf("Literal = %q, want %q", fa.Literal, tt.wantLit)
			}
			if fa.Source != tt.wantSrc {
				t.Errorf("Source = %q, want %q", fa.Source, tt.wantSrc)
			}
		})
	}
}
