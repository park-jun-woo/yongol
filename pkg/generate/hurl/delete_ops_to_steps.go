//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what deleteOpsToSteps — deleteOp 목록을 hurl step 목록으로 변환
package hurl

// deleteOpsToSteps converts sorted deleteOps into hurl steps.
func deleteOpsToSteps(ops []deleteOp, ctx *scenarioCtx) []step {
	var steps []step
	for _, d := range ops {
		s, ok := buildDeleteStep(ctx, d, len(steps) == 0)
		if !ok {
			continue
		}
		steps = append(steps, s)
	}
	return steps
}
