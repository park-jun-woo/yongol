//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what recordCaptureVars — scenarioCtx.captures 에 각 capture VarName 을 등록
package hurl

// recordCaptureVars writes every capture's VarName into ctx.captures so
// downstream phases can see them when resolving {path_params} and token
// variables. Pulled out of buildPostStep to keep that function control=sequence.
func recordCaptureVars(ctx *scenarioCtx, captures []capture) {
	for _, c := range captures {
		ctx.captures[c.VarName] = true
	}
}
