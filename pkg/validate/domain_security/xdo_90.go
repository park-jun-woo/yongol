//ff:func feature=validate type=rule control=iteration dimension=1 topic=domain-security
//ff:what XDO-90 — 동일 operationId가 두 개 이상 도메인 OpenAPI에 선언되면 ERROR
package domain_security

import (
	"fmt"
	"sort"

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
		collectDocOpIDs(dd, opDomains)
	}

	var diags []diagnostic.Diagnostic
	for opID, domains := range opDomains {
		if len(domains) < 2 {
			continue
		}
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
