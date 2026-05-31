//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what enrichFieldArgDDL -- Pass 1: 모든 Op 의 FieldArg.SourceColumn 세팅 / Pass 2: CRUD Op DDL 테이블 매칭으로 ColumnName/IsPK 세팅

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// enrichFieldArgDDL runs two passes over all ops:
//
// Pass 1 sets SourceColumn on every FieldArg that has a Field accessor,
// regardless of Op kind. This ensures @call/@eval/@auth ops get proper
// snake_case source column names even without DDL table matching.
//
// Pass 2 enriches CRUD ops (get/post/put/delete) with DDL table metadata,
// setting ColumnName and IsPK by matching the argument Key against DDL
// table columns.
func enrichFieldArgDDL(ops []Op, fs *yongol.Fullstack) {
	for i := range ops {
		enrichOpSourceColumns(&ops[i])
	}
	if fs == nil || len(fs.DDLTables) == 0 {
		return
	}
	for i := range ops {
		enrichOpDDLColumns(&ops[i], fs)
	}
}
