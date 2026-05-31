//ff:func feature=gen-nestjs type=generator control=iteration dimension=1
//ff:what appendReverseRelations — 단일 테이블의 FK 들을 참조 테이블 키 아래 역관계로 추가

package prisma

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// appendReverseRelations appends one reverse relation entry per foreign key of
// table into rm, keyed by the referenced table name.
func appendReverseRelations(rm map[string][]reverseRelation, table ddl.Table) {
	sourceModelName := pascalCase(singularize(table.Name))
	for _, fk := range table.ForeignKeys {
		// Reverse field name: pluralized source table name.
		rm[fk.RefTable] = append(rm[fk.RefTable], reverseRelation{
			FieldName: table.Name,
			ModelName: sourceModelName,
		})
	}
}
