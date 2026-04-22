//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-structural
//ff:what collectPathVars — path 템플릿의 {var} 변수 이름 집합 추출

package openapi

// collectPathVars returns the set of variable names declared in an OpenAPI
// path template. e.g. `/workflows/{id}/actions/{aid}` → {"id", "aid"}.
func collectPathVars(path string) map[string]bool {
	set := map[string]bool{}
	for _, m := range rePathVar.FindAllStringSubmatch(path, -1) {
		set[m[1]] = true
	}
	return set
}
