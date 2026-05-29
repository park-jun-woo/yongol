//ff:func feature=rule type=test control=iteration dimension=1
//ff:what resolveOAPIParamGoType test — OpenAPI type+format → Go 타입 매핑 검증

package ground

import "testing"

func TestResolveOAPIParamGoType(t *testing.T) {
	tests := []struct {
		baseType string
		format   string
		want     string
	}{
		{"string", "uuid", "openapi_types.UUID"},
		{"string", "email", "openapi_types.Email"},
		{"string", "", "string"},
		{"string", "date-time", "string"},
		{"integer", "int32", "int32"},
		{"integer", "int64", "int64"},
		{"integer", "", "int32"},
		{"boolean", "", "bool"},
		{"number", "float", ""},
		{"object", "", ""},
		{"", "", ""},
	}
	for _, tt := range tests {
		got := resolveOAPIParamGoType(tt.baseType, tt.format)
		if got != tt.want {
			t.Errorf("resolveOAPIParamGoType(%q, %q) = %q, want %q",
				tt.baseType, tt.format, got, tt.want)
		}
	}
}
