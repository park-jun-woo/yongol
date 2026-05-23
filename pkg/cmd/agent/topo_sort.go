//ff:func feature=agent type=helper control=iteration dimension=2
//ff:what topoSortTables — belongs_to 기반 위상 정렬 (순환 시 끊고 반환)

package agent

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// topoSortTables returns table names in topological order based on belongs_to
// relationships. Parents (dependencies) come before children. If a cycle is
// detected, the cycle edge is broken and the remaining order is returned.
func topoSortTables(tables map[string]features.TableDef) []string {
	// Collect all table names sorted for deterministic output
	names := make([]string, 0, len(tables))
	for name := range tables {
		names = append(names, name)
	}
	sort.Strings(names)

	// Build adjacency: child -> parents (belongs_to)
	inDegree := make(map[string]int, len(names))
	children := make(map[string][]string, len(names)) // parent -> children
	for _, name := range names {
		inDegree[name] = 0
	}
	for _, name := range names {
		td := tables[name]
		for _, parent := range td.BelongsTo {
			if _, ok := tables[parent]; !ok {
				continue // skip references to tables not in the map
			}
			inDegree[name]++
			children[parent] = append(children[parent], name)
		}
	}

	// Kahn's algorithm
	var queue []string
	for _, name := range names {
		if inDegree[name] == 0 {
			queue = append(queue, name)
		}
	}

	var result []string
	for len(queue) > 0 {
		// pop front
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, child := range children[node] {
			inDegree[child]--
			if inDegree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	// If cycle exists, remaining nodes have inDegree > 0 — append them sorted
	if len(result) < len(names) {
		var remaining []string
		inResult := make(map[string]bool, len(result))
		for _, r := range result {
			inResult[r] = true
		}
		for _, name := range names {
			if !inResult[name] {
				remaining = append(remaining, name)
			}
		}
		sort.Strings(remaining)
		result = append(result, remaining...)
	}

	return result
}
