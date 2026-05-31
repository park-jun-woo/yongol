//ff:func feature=gen-nestjs type=generator control=iteration dimension=1
//ff:what buildReverseRelations — 전체 테이블 FK 를 스캔해 참조 테이블명→역관계 목록 맵 구성

package prisma

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// buildReverseRelations scans all tables' ForeignKeys and returns a map
// from referenced table name to the list of reverse relation entries.
func buildReverseRelations(tables []ddl.Table) map[string][]reverseRelation {
	rm := make(map[string][]reverseRelation)
	for _, table := range tables {
		appendReverseRelations(rm, table)
	}
	return rm
}
