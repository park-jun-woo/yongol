//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what collectComponentApiOps — 컴포넌트 .tsx에서 api.<Op>( 호출을 스캔해 실재 operationId만 소비로 수집

package stml_openapi

import (
	"os"
	"path/filepath"
	"regexp"
)

// componentApiCallRe matches api.<identifier>( calls. The capture is taken as
// an operationId candidate; case is not forced because OpenAPI authors choose
// operationId freely. Candidates are filtered against real operationIds by the
// caller via the ops set.
var componentApiCallRe = regexp.MustCompile(`\bapi\.([A-Za-z_][A-Za-z0-9_]*)\s*\(`)

// collectComponentApiOps reads each component's <specsDir>/frontend/components/
// <Name>.tsx file, extracts api.<operationId>( calls, and adds those that exist
// in ops to out. Missing or unreadable files are skipped (zero consumption).
func collectComponentApiOps(names map[string]struct{}, specsDir string, ops map[string]struct{}, out map[string]struct{}) {
	if specsDir == "" {
		return
	}
	for name := range names {
		path := filepath.Join(specsDir, "frontend", "components", name+".tsx")
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		for _, m := range componentApiCallRe.FindAllStringSubmatch(string(data), -1) {
			if _, ok := ops[m[1]]; ok {
				out[m[1]] = struct{}{}
			}
		}
	}
}
