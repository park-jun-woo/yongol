//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what buildFKDepthMap — DDL FK 의존 깊이 맵 구축
package hurl

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// buildFKDepthMap computes the maximum FK dependency depth for each table.
func buildFKDepthMap(tables []ddl.Table) map[string]int {
	parentMap := map[string][]string{}
	for _, t := range tables {
		name := strings.TrimSuffix(t.Name, "s")
		for _, fk := range t.ForeignKeys {
			parentMap[name] = append(parentMap[name], strings.TrimSuffix(fk.RefTable, "s"))
		}
	}
	cache := map[string]int{}
	for _, t := range tables {
		computeDepth(strings.TrimSuffix(t.Name, "s"), parentMap, cache)
	}
	return cache
}
