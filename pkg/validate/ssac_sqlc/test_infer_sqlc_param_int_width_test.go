//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestInferSqlcParamIntWidth — sqlc 쿼리 body 캐스트에서 정수 폭 추론 검증

package ssac_sqlc

import "testing"

func TestInferSqlcParamIntWidth(t *testing.T) {
	tests := []struct {
		name  string
		body  string
		param string
		want  string
	}{
		{"NoCast_Int32", "SELECT * FROM items LIMIT sqlc.arg(per_page) OFFSET sqlc.arg(page)", "per_page", "int32"},
		{"Bigint_Int64", "SELECT * FROM items LIMIT sqlc.arg(per_page)::bigint", "per_page", "int64"},
		{"Int8_Int64", "SELECT * FROM items LIMIT sqlc.arg(per_page)::int8", "per_page", "int64"},
		{"Int_Int32", "SELECT * FROM items LIMIT sqlc.arg(per_page)::int", "per_page", "int32"},
		{"Int4_Int32", "SELECT * FROM items LIMIT sqlc.arg(per_page)::int4", "per_page", "int32"},
		{"Integer_Int32", "SELECT * FROM items LIMIT sqlc.arg(per_page)::integer", "per_page", "int32"},
		{"SpacesAroundCast", "LIMIT sqlc.arg( per_page ) :: bigint", "per_page", "int64"},
		{"AtSyntax_NoCast", "SELECT * FROM items LIMIT @per_page", "per_page", "int32"},
		{"AtSyntax_Bigint", "SELECT * FROM items LIMIT @per_page::bigint", "per_page", "int64"},
		{"NotFound", "SELECT * FROM items WHERE id = sqlc.arg(item_id)", "per_page", ""},
		{"TextCast_Empty", "SELECT * FROM items WHERE name = sqlc.arg(name)::text", "name", ""},
	}
	for _, tt := range tests {
		got := inferSqlcParamIntWidth(tt.body, tt.param)
		if got != tt.want {
			t.Errorf("%s: want %q, got %q", tt.name, tt.want, got)
		}
	}
}
