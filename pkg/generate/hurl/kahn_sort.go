//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what kahnSort — Kahn 알고리즘으로 위상 정렬 실행 (부모 먼저)
package hurl

// kahnSort performs topological sort using Kahn's algorithm (parents first).
func kahnSort(children map[string][]string, indegree map[string]int, allNames map[string]bool) []string {
	var queue []string
	for name := range allNames {
		if indegree[name] == 0 {
			queue = append(queue, name)
		}
	}
	var result []string
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		result = append(result, curr)
		for _, child := range children[curr] {
			indegree[child]--
			if indegree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}
	return result
}
