//ff:func feature=migration type=test control=sequence
//ff:what io_helpers_unit_test — LoadSnapshot/WriteSnapshot/NextSequenceNumber/LoadDataMigrationSQL/listSQLFiles/loadPrevSnapshot 단위 테스트
package migration

func sampleSchema() *Schema {
	s := NewSchema()
	t := ensureTable(s, "users")
	t.Columns = []*Column{
		{Name: "id", Type: CanonicalType{Base: "BIGINT"}, Nullable: false},
		{Name: "email", Type: CanonicalType{Base: "TEXT"}, Nullable: true},
	}
	t.PrimaryKey = []string{"id"}
	return s
}
