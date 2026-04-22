//ff:func feature=gen-hurl type=util control=sequence
//ff:what topoSortDelete — DDL FK 기반 위상 정렬로 삭제 순서 결정 (자식 먼저)
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// topoSortDelete returns table names in FK-reverse topological order (children first).
func topoSortDelete(tables []ddl.Table) []string {
	children, indegree, allNames := buildFKGraph(tables)
	parentFirst := kahnSort(children, indegree, allNames)
	return reverseStrings(parentFirst)
}
