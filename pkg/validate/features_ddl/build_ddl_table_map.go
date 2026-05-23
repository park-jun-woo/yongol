//ff:func feature=validate type=helper control=iteration dimension=1 topic=features-ddl
//ff:what buildDDLTableMap — DDL 테이블 슬라이스를 이름 → *ddl.Table 맵으로 변환

package features_ddl

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// buildDDLTableMap creates a lookup map from table name to ddl.Table.
func buildDDLTableMap(tables []ddl.Table) map[string]*ddl.Table {
	m := make(map[string]*ddl.Table, len(tables))
	for i := range tables {
		m[tables[i].Name] = &tables[i]
	}
	return m
}
