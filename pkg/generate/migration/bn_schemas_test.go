//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

func bnSchemas() (prev, curr *Table) {
	prev = &Table{
		Name: "users",
		Columns: []*Column{
			{Name: "id", Type: CanonicalType{Base: "BIGINT"}, Nullable: false},
			{Name: "email", Type: CanonicalType{Base: "VARCHAR", Length: 255}, Nullable: false},
			{Name: "legacy", Type: CanonicalType{Base: "TEXT"}, Nullable: true},
		},
		Indexes: []*Index{
			{Name: "idx_old", Columns: []string{"email"}},
		},
		ForeignKeys: []*ForeignKey{
			{Name: "fk_old", Columns: []string{"org_id"}, RefTable: "orgs", RefColumns: []string{"id"}},
		},
		Checks: []*CheckConstraint{
			{Name: "users_chk_old", Expression: "id > 0"},
		},
	}
	curr = &Table{
		Name: "users",
		Columns: []*Column{
			{Name: "id", Type: CanonicalType{Base: "BIGINT"}, Nullable: false},
			{Name: "email", Type: CanonicalType{Base: "VARCHAR", Length: 320}, Nullable: false, Default: "''"},
			{Name: "age", Type: CanonicalType{Base: "INTEGER"}, Nullable: false},
		},
		Indexes: []*Index{
			{Name: "idx_new", Columns: []string{"age"}},
		},
		ForeignKeys: []*ForeignKey{
			{Name: "fk_new", Columns: []string{"org_id"}, RefTable: "orgs", RefColumns: []string{"id"}, OnDelete: "CASCADE"},
		},
		Checks: []*CheckConstraint{
			{Name: "users_chk_new", Expression: "age >= 0"},
		},
	}
	return prev, curr
}
