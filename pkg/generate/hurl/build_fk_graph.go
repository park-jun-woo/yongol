//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what buildFKGraph — DDL tables에서 FK 부모-자식 그래프 구축
package hurl

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// buildFKGraph builds parent->children adjacency, indegree map, and all table names.
func buildFKGraph(tables []ddl.Table) (map[string][]string, map[string]int, map[string]bool) {
	children := map[string][]string{}
	indegree := map[string]int{}
	all := map[string]bool{}
	for _, t := range tables {
		name := strings.TrimSuffix(t.Name, "s")
		all[name] = true
		for _, fk := range t.ForeignKeys {
			parent := strings.TrimSuffix(fk.RefTable, "s")
			children[parent] = append(children[parent], name)
			indegree[name]++
		}
	}
	return children, indegree, all
}
