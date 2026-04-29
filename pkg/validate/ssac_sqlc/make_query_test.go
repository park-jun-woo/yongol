//ff:func feature=validate type=test-helper control=selection topic=ssac-sqlc
//ff:what makeQuery — XQS-20 테스트용 sqlc QuerySpec 생성 헬퍼 (RowType 자동)

package ssac_sqlc

import sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"

// makeQuery wraps a sqlc QuerySpec literal with sensible defaults for the
// XQS-20 tests. RowType is computed exactly as the sqlc parser would
// (`<Name>Row` for :one/:many, empty otherwise).
func makeQuery(name, model, method, cardinality, body string) sqlcparser.QuerySpec {
	var rowType string
	switch cardinality {
	case "one", "many":
		rowType = name + "Row"
	default:
		rowType = ""
	}
	return sqlcparser.QuerySpec{
		Name:        name,
		Model:       model,
		Method:      method,
		Cardinality: cardinality,
		RowType:     rowType,
		Body:        body,
		File:        "users.sql",
		Line:        1,
	}
}
