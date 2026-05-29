//ff:func feature=migration type=test-helper control=sequence
//ff:what normalizeTypeCases — TestNormalizeType 가 사용하는 케이스 목록
package migration

func normalizeTypeCases() []normalizeTypeCase {
	return []normalizeTypeCase{
		{"int", "INTEGER", 0, 0, 0, false, false},
		{"int4", "INTEGER", 0, 0, 0, false, false},
		{"integer", "INTEGER", 0, 0, 0, false, false},
		{"INTEGER", "INTEGER", 0, 0, 0, false, false},
		{"bigint", "BIGINT", 0, 0, 0, false, false},
		{"int8", "BIGINT", 0, 0, 0, false, false},
		{"smallint", "SMALLINT", 0, 0, 0, false, false},
		{"int2", "SMALLINT", 0, 0, 0, false, false},
		{"serial", "INTEGER", 0, 0, 0, false, true},
		{"BIGSERIAL", "BIGINT", 0, 0, 0, false, true},
		{"smallserial", "SMALLINT", 0, 0, 0, false, true},
		{"bool", "BOOLEAN", 0, 0, 0, false, false},
		{"boolean", "BOOLEAN", 0, 0, 0, false, false},
		{"varchar(255)", "VARCHAR", 255, 0, 0, false, false},
		{"character varying(64)", "VARCHAR", 64, 0, 0, false, false},
		{"VARCHAR(1024)", "VARCHAR", 1024, 0, 0, false, false},
		{"char(10)", "CHAR", 10, 0, 0, false, false},
		{"character(5)", "CHAR", 5, 0, 0, false, false},
		{"text", "TEXT", 0, 0, 0, false, false},
		{"uuid", "UUID", 0, 0, 0, false, false},
		{"jsonb", "JSONB", 0, 0, 0, false, false},
		{"json", "JSON", 0, 0, 0, false, false},
		{"bytea", "BYTEA", 0, 0, 0, false, false},
		{"timestamp", "TIMESTAMP", 0, 0, 0, false, false},
		{"timestamp without time zone", "TIMESTAMP", 0, 0, 0, false, false},
		{"timestamptz", "TIMESTAMPTZ", 0, 0, 0, false, false},
		{"timestamp with time zone", "TIMESTAMPTZ", 0, 0, 0, false, false},
		{"date", "DATE", 0, 0, 0, false, false},
		{"numeric(10,2)", "NUMERIC", 0, 10, 2, false, false},
		{"decimal(5,0)", "NUMERIC", 0, 5, 0, false, false},
		{"numeric(7)", "NUMERIC", 0, 7, 0, false, false},
		{"integer[]", "INTEGER", 0, 0, 0, true, false},
		{"text[]", "TEXT", 0, 0, 0, true, false},
		{"real", "REAL", 0, 0, 0, false, false},
		{"double precision", "DOUBLE PRECISION", 0, 0, 0, false, false},
	}
}
