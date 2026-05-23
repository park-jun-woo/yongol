//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what extractErrorOps — 검증 에러에서 원인 operationId 탐색

package agent

import (

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func extractErrorOps(err error, offsets []pathOffset, feats []features.Feature, yamlContent string) ([]string, map[string]int) {
	if err == nil {
		return nil, nil
	}
	msg := err.Error()
	seen := make(map[string]bool)

	if ops := matchByLine(msg, offsets); len(ops) > 0 {
		for _, op := range ops {
			seen[op] = true
		}
	}
	if ops := matchBySchema(msg, offsets, feats); len(ops) > 0 {
		for _, op := range ops {
			seen[op] = true
		}
	}
	if ops := matchByPath(msg, offsets); len(ops) > 0 {
		for _, op := range ops {
			seen[op] = true
		}
	}

	var relativeLines map[string]int
	if grepOps, rl := matchByGrep(msg, yamlContent, offsets); len(grepOps) > 0 {
		relativeLines = rl
		for _, op := range grepOps {
			seen[op] = true
		}
	}

	out := make([]string, 0, len(seen))
	for op := range seen {
		out = append(out, op)
	}
	return out, relativeLines
}
