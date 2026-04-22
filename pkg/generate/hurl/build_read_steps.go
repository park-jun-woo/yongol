//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what buildReads — Phase4: 모든 GET endpoint를 PathDepth/Path 순으로 단일 출력
package hurl

import (
	"sort"
	"strings"
)

// buildReads collects every GET endpoint (excluding auth ops) and emits them
// once, sorted by PathDepth ascending then Path alphabetically. This replaces
// the earlier split between buildCRUDReadSteps (all GETs) and buildReadSteps
// (non-mutation GETs) which caused duplicate output.
func buildReads(ctx *scenarioCtx) []step {
	if ctx.fs.OpenAPIDoc == nil {
		return nil
	}

	type pathedOp struct {
		path   string
		opID   string
		method string
	}
	var collected []pathedOp
	for path, pathItem := range ctx.fs.OpenAPIDoc.Paths.Map() {
		for method, op := range pathItem.Operations() {
			if strings.ToUpper(method) != "GET" {
				continue
			}
			if op.OperationID == "" || isAuthOpID(op.OperationID) {
				continue
			}
			collected = append(collected, pathedOp{path: path, opID: op.OperationID, method: "GET"})
		}
	}

	sort.SliceStable(collected, func(i, j int) bool {
		di := strings.Count(collected[i].path, "/")
		dj := strings.Count(collected[j].path, "/")
		if di != dj {
			return di < dj
		}
		return collected[i].path < collected[j].path
	})

	var steps []step
	for _, p := range collected {
		s, ok := buildReadStep(ctx, p.path, len(steps) == 0)
		if !ok {
			continue
		}
		steps = append(steps, s)
	}
	return steps
}
