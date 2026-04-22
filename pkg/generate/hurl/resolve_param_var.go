//ff:func feature=gen-hurl type=util control=sequence
//ff:what resolveParamVar — path segment의 {param}을 {{resource_param}} 변수로 변환 (snake 정규화)
package hurl

// resolveParamVar converts a path segment like {id} to {{resource_id}}.
// Identifiers are normalized to snake_case via snakeHurlName so PascalCase
// ({GigID}) and hyphenated resources (/audit-logs/) both yield valid hurl
// variable names without hyphens.
func resolveParamVar(parts []string, idx int) string {
	matches := pathParamRe.FindStringSubmatch(parts[idx])
	if len(matches) < 2 {
		return parts[idx]
	}
	paramSnake := snakeHurlName(matches[1])
	resource := findPrecedingResource(parts, idx)
	varName := paramSnake
	if resource != "" && paramSnake == "id" {
		varName = resource + "_id"
	}
	return "{{" + varName + "}}"
}
