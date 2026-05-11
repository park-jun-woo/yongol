//ff:func feature=validate type=rule control=iteration dimension=2 topic=domain-security
//ff:what XDO-90 — 동일 operationId가 두 개 이상 도메인 OpenAPI에 선언되면 ERROR
package domain_security

import (
	"fmt"
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdo90DuplicateOperationID detects operationIds that appear in more than one
// domain's OpenAPI specification.
func xdo90DuplicateOperationID(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	docs := loadDomainOpenAPIDocs(fs)
	if len(docs) < 2 {
		return nil
	}

	// opID → list of domain names that declare it.
	opDomains := make(map[string][]string)

	for _, dd := range docs {
		if dd.Doc.Paths == nil {
			continue
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

	var diags []diagnostic.Diagnostic
	for opID, domains := range opDomains {
		if len(domains) < 2 {
			continue
		}
		// Deduplicate domain names (same domain could have multiple paths with same opID — that's a separate issue).
		unique := uniqueStrings(domains)
		if len(unique) < 2 {
			continue
		}
		sort.Strings(unique)
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[XDO-90] operationId %q is declared in multiple domains: %v", opID, unique),
			Advice:  fmt.Sprintf("Ensure operationId %q is unique across all domain OpenAPI files; rename one of the duplicates", opID),
		})
	}
	return diags
}

// uniqueStrings returns deduplicated slice preserving first occurrence order.
func uniqueStrings(ss []string) []string {
	seen := make(map[string]struct{})
	var out []string
	for _, s := range ss {
		if _, ok := seen[s]; !ok {
			seen[s] = struct{}{}
			out = append(out, s)
		}
	}
	return out
}
