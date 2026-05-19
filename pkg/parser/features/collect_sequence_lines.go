//ff:func feature=features type=util control=iteration dimension=1
//ff:what collectSequenceLines — SequenceNode 자식 노드의 라인 번호 슬라이스 반환
package features

import "gopkg.in/yaml.v3"

// collectSequenceLines converts a yaml.SequenceNode's child items into a
// slice of 1-based line numbers.
func collectSequenceLines(seq *yaml.Node) []int {
	lines := make([]int, len(seq.Content))
	for j, item := range seq.Content {
		lines[j] = item.Line
	}
	return lines
}
