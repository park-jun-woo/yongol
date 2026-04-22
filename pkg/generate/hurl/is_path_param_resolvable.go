//ff:func feature=gen-hurl type=util control=sequence
//ff:what isPathParamResolvable — path 세그먼트 하나가 captures로 해소 가능한지 판정
package hurl

// isPathParamResolvable returns true when the single path segment `part` at
// index `i` of `parts` is either (a) not a {param} literal, or (b) a {param}
// whose canonical snake name is captured, or (c) a plain "{id}" whose
// derived "<resource>_id" is captured.
//
// Factored out of canResolvePathParams to keep loop body at depth 1. False
// from here makes the caller abort the whole path resolution.
func isPathParamResolvable(parts []string, i int, part string, captures map[string]bool) bool {
	if !pathParamRe.MatchString(part) {
		return true
	}
	m := pathParamRe.FindStringSubmatch(part)
	if len(m) < 2 {
		return false
	}
	paramSnake := snakeHurlName(m[1])
	if captures[paramSnake] {
		return true
	}
	if paramSnake != "id" {
		return false
	}
	resource := findPrecedingResource(parts, i)
	return resource != "" && captures[resource+"_id"]
}
