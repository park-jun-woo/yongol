//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what sortDeleteOps — FK 역순으로 deleteOp 슬라이스 정렬
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// sortDeleteOps sorts delete operations by FK reverse order (children first).
func sortDeleteOps(ops []deleteOp, tables []ddl.Table) {
	orderMap := map[string]int{}
	for i, r := range topoSortDelete(tables) {
		orderMap[r] = i
	}
	for i := 0; i < len(ops); i++ {
		for j := i + 1; j < len(ops); j++ {
			if orderMap[ops[i].resource] > orderMap[ops[j].resource] {
				ops[i], ops[j] = ops[j], ops[i]
			}
		}
	}
}
