//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what firstProtectedOpPage — security 보호 op을 호출하는 첫 STML 페이지 파일명 반환

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/yongol"

// firstProtectedOpPage scans the STML pages in order and returns the file
// name of the first page whose data-fetch or data-action consumes a
// security-protected OpenAPI operation (OpRequiresAuth), and whether any
// such page exists. TM-21/22 use it to decide if the captured token has a
// consumer and where to point the diagnostic.
func firstProtectedOpPage(fs *yongol.Fullstack, opMap map[string]operationEntry) (string, bool) {
	for _, page := range fs.STMLPages {
		ops := make(map[string]struct{})
		for _, f := range page.Fetches {
			collectFetchOps(f, ops)
		}
		for _, a := range page.Actions {
			ops[a.OperationID] = struct{}{}
		}
		for opID := range ops {
			entry, ok := opMap[opID]
			if !ok {
				continue
			}
			if OpRequiresAuth(entry.op, fs.OpenAPIDoc) {
				return page.FileName, true
			}
		}
	}
	return "", false
}
