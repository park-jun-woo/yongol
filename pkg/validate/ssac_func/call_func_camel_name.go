//ff:func feature=validate type=util control=sequence topic=func-check
//ff:what callFuncCamelName — convert a @call model to the camelCase func annotation key

package ssac_func

// callFuncCamelName returns the camelCase @func annotation form for a SSaC
// @call model ("billing.CheckCredits" -> "checkCredits"). Returns "" if the
// model is not a qualified pkg.Func form. Used to bridge PascalCase @call
// references against camelCase Func.spec / Func.request keys in Ground.
func callFuncCamelName(model string) string {
	name := callFuncName(model)
	if name == "" {
		return ""
	}
	return toCamelKey(name)
}
