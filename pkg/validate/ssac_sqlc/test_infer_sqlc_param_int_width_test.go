//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestInferSqlcParamIntWidth — infer int width from sqlc query body casts

package ssac_sqlc

import "testing"

func TestInferSqlcParamIntWidth_NoCast_Int32(t *testing.T) {
	body := "SELECT * FROM items LIMIT sqlc.arg(per_page) OFFSET sqlc.arg(page)"
	if w := inferSqlcParamIntWidth(body, "per_page"); w != "int32" {
		t.Errorf("no cast: want int32, got %q", w)
	}
}

func TestInferSqlcParamIntWidth_Bigint_Int64(t *testing.T) {
	body := "SELECT * FROM items LIMIT sqlc.arg(per_page)::bigint"
	if w := inferSqlcParamIntWidth(body, "per_page"); w != "int64" {
		t.Errorf("::bigint: want int64, got %q", w)
	}
}

func TestInferSqlcParamIntWidth_Int8_Int64(t *testing.T) {
	body := "SELECT * FROM items LIMIT sqlc.arg(per_page)::int8"
	if w := inferSqlcParamIntWidth(body, "per_page"); w != "int64" {
		t.Errorf("::int8: want int64, got %q", w)
	}
}

func TestInferSqlcParamIntWidth_Int_Int32(t *testing.T) {
	body := "SELECT * FROM items LIMIT sqlc.arg(per_page)::int"
	if w := inferSqlcParamIntWidth(body, "per_page"); w != "int32" {
		t.Errorf("::int: want int32, got %q", w)
	}
}

func TestInferSqlcParamIntWidth_Int4_Int32(t *testing.T) {
	body := "SELECT * FROM items LIMIT sqlc.arg(per_page)::int4"
	if w := inferSqlcParamIntWidth(body, "per_page"); w != "int32" {
		t.Errorf("::int4: want int32, got %q", w)
	}
}

func TestInferSqlcParamIntWidth_Integer_Int32(t *testing.T) {
	body := "SELECT * FROM items LIMIT sqlc.arg(per_page)::integer"
	if w := inferSqlcParamIntWidth(body, "per_page"); w != "int32" {
		t.Errorf("::integer: want int32, got %q", w)
	}
}

func TestInferSqlcParamIntWidth_SpacesAroundCast(t *testing.T) {
	body := "LIMIT sqlc.arg( per_page ) :: bigint"
	if w := inferSqlcParamIntWidth(body, "per_page"); w != "int64" {
		t.Errorf("spaced cast: want int64, got %q", w)
	}
}

func TestInferSqlcParamIntWidth_AtSyntax_NoCast(t *testing.T) {
	body := "SELECT * FROM items LIMIT @per_page"
	if w := inferSqlcParamIntWidth(body, "per_page"); w != "int32" {
		t.Errorf("@name no cast: want int32, got %q", w)
	}
}

func TestInferSqlcParamIntWidth_AtSyntax_Bigint(t *testing.T) {
	body := "SELECT * FROM items LIMIT @per_page::bigint"
	if w := inferSqlcParamIntWidth(body, "per_page"); w != "int64" {
		t.Errorf("@name::bigint: want int64, got %q", w)
	}
}

func TestInferSqlcParamIntWidth_NotFound(t *testing.T) {
	body := "SELECT * FROM items WHERE id = sqlc.arg(item_id)"
	if w := inferSqlcParamIntWidth(body, "per_page"); w != "" {
		t.Errorf("param not found: want empty, got %q", w)
	}
}

func TestInferSqlcParamIntWidth_TextCast_Empty(t *testing.T) {
	body := "SELECT * FROM items WHERE name = sqlc.arg(name)::text"
	if w := inferSqlcParamIntWidth(body, "name"); w != "" {
		t.Errorf("::text: want empty, got %q", w)
	}
}
