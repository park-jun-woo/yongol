//ff:func feature=validate type=rule control=iteration dimension=2 topic=domain-security
//ff:what XDS-82 — public OpenAPI에 DELETE operation이 admin Rego 룰 없이 존재하면 ERROR
package domain_security

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xds82PublicDeleteNoRego detects public domain DELETE operations that lack
// a corresponding admin Rego allow rule with "delete" action.
func xds82PublicDeleteNoRego(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	docs := loadDomainOpenAPIDocs(fs)
	regoDeleteResources := collectRegoDeleteActions(fs)

	var diags []diagnostic.Diagnostic
	for _, dd := range docs {
		if dd.Name != "public" || dd.Doc.Paths == nil {
			continue
		}
		for path, item := range dd.Doc.Paths.Map() {
			if item.Delete == nil || item.Delete.OperationID == "" {
				continue
			}
			if !hasRegoDeleteRule(item.Delete.OperationID, regoDeleteResources) {
				diags = append(diags, diagnostic.Diagnostic{
					File:    dd.Cfg.OpenAPI,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[XDS-82] public DELETE endpoint %q (%s) has no corresponding admin Rego allow rule", item.Delete.OperationID, path),
					Advice:  "Add a Rego allow rule with action \"delete\" for this resource, or move the endpoint to the admin domain",
				})
			}
		}
	}
	return diags
}
