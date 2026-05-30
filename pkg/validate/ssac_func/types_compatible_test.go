//ff:func feature=validate type=test control=iteration dimension=1 topic=func-check
//ff:what TypesCompatible test — exact/pointer/int family/UUID family 호환 + pgtype.UUID 제외 검증

package ssac_func

import "testing"

func TestTypesCompatible(t *testing.T) {
	tests := []struct {
		name           string
		actual         string
		expected       string
		wantCompatible bool
	}{
		{"exact equality", "string", "string", true},
		{"pointer stripped", "*string", "string", true},
		{"int family", "int32", "int64", true},
		{"nil to pointer", "nil", "*User", true},
		{"object vs primitive", "User", "string", false},

		// UUID family (Phase018): the oapi-codegen runtime type under either
		// import alias is interchangeable, including pointer forms.
		{"openapi_types.UUID == types.UUID", "openapi_types.UUID", "types.UUID", true},
		{"types.UUID == openapi_types.UUID", "types.UUID", "openapi_types.UUID", true},
		{"*openapi_types.UUID == openapi_types.UUID", "*openapi_types.UUID", "openapi_types.UUID", true},
		{"openapi_types.UUID == openapi_types.UUID", "openapi_types.UUID", "openapi_types.UUID", true},

		// pgtype.UUID is the DB/sqlc type — must NOT be compatible with the
		// api UUID type (func→api converter does no conversion).
		{"pgtype.UUID vs openapi_types.UUID", "pgtype.UUID", "openapi_types.UUID", false},
		{"openapi_types.UUID vs pgtype.UUID", "openapi_types.UUID", "pgtype.UUID", false},
		{"string vs openapi_types.UUID", "string", "openapi_types.UUID", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TypesCompatible(tt.actual, tt.expected)
			if got != tt.wantCompatible {
				t.Errorf("TypesCompatible(%q, %q) = %v, want %v",
					tt.actual, tt.expected, got, tt.wantCompatible)
			}
		})
	}
}
