//ff:func feature=validate type=util control=iteration dimension=2 topic=domain-security
//ff:what collectDocOpDomains — 단일 domainDoc에서 operationId → 도메인 매핑 수집
package domain_security

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// collectDocOpDomains collects operationId-to-domain mappings from a single doc.
func collectDocOpDomains(dd domainDoc, result map[string]string) {
	for _, item := range dd.Doc.Paths.Map() {
		ops := []*openapi3.Operation{
			item.Get, item.Post, item.Put, item.Delete, item.Patch,
		}
		for _, op := range ops {
			if op == nil || op.OperationID == "" {
				continue
			}
			if _, exists := result[op.OperationID]; !exists {
				result[op.OperationID] = dd.Name
			}
		}
	}
}
