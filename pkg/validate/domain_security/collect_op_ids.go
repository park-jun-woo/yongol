//ff:func feature=validate type=util control=iteration dimension=2 topic=domain-security
//ff:what collectDocOpIDs — 단일 domainDoc에서 operationId별 도메인 이름 수집
package domain_security

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// collectDocOpIDs collects all operationIds from a single domain document
// into the provided map (opID → list of domain names).
func collectDocOpIDs(dd domainDoc, opDomains map[string][]string) {
	if dd.Doc.Paths == nil {
		return
	}
	for _, item := range dd.Doc.Paths.Map() {
		ops := []*openapi3.Operation{
			item.Get, item.Post, item.Put, item.Delete, item.Patch,
		}
		for _, op := range ops {
			if op == nil || op.OperationID == "" {
				continue
			}
			opDomains[op.OperationID] = append(opDomains[op.OperationID], dd.Name)
		}
	}
}
