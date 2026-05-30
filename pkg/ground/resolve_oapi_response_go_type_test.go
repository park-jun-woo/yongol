//ff:func feature=rule type=test control=iteration dimension=1
//ff:what resolveOAPIResponseGoType test — 응답 본문 type+format → Go 타입 매핑 검증 (date-time→time.Time)

package ground

import "testing"

func TestResolveOAPIResponseGoType(t *testing.T) {
	tests := []struct {
		baseType string
		format   string
		want     string
	}{
		{"string", "uuid", "openapi_types.UUID"},
		{"string", "email", "openapi_types.Email"},
		// Differs from resolveOAPIParamGoType: response-body date-time → time.Time.
		{"string", "date-time", "time.Time"},
		{"string", "", "string"},
		{"string", "byte", "string"},
		// Non-string base types: caller falls back to resolveSchemaType.
		{"integer", "int64", ""},
		{"boolean", "", ""},
		{"", "", ""},
	}
	for _, tt := range tests {
		got := resolveOAPIResponseGoType(tt.baseType, tt.format)
		if got != tt.want {
			t.Errorf("resolveOAPIResponseGoType(%q, %q) = %q, want %q",
				tt.baseType, tt.format, got, tt.want)
		}
	}
}
