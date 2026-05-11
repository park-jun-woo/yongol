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

	// Collect all "delete" actions from Rego policies.
	regoDeleteResources := collectRegoDeleteActions(fs)

	var diags []diagnostic.Diagnostic
	for _, dd := range docs {
		if dd.Name != "public" {
			continue
		}
		if dd.Doc.Paths == nil {
			continue
		}
		for path, item := range dd.Doc.Paths.Map() {
			if item.Delete == nil {
				continue
			}
			op := item.Delete
			if op.OperationID == "" {
				continue
			}
			// Check if there's a Rego rule covering this delete operation.
			if !hasRegoDeleteRule(op.OperationID, regoDeleteResources) {
				diags = append(diags, diagnostic.Diagnostic{
					File:    dd.Cfg.OpenAPI,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[XDS-82] public DELETE endpoint %q (%s) has no corresponding admin Rego allow rule", op.OperationID, path),
					Advice:  "Add a Rego allow rule with action \"delete\" for this resource, or move the endpoint to the admin domain",
				})
			}
		}
	}
	return diags
}

// collectRegoDeleteActions gathers resources that have "delete" actions in Rego policies.
func collectRegoDeleteActions(fs *yongol.Fullstack) map[string]struct{} {
	result := make(map[string]struct{})
	for _, policy := range fs.ParsedPolicies {
		for _, rule := range policy.Rules {
			for _, action := range rule.Actions {
				if action == "delete" {
					result[rule.Resource] = struct{}{}
				}
			}
		}
	}
	return result
}

// hasRegoDeleteRule checks if the operationId is covered by any Rego delete rule.
// It checks both exact resource match and plural/singular variations.
func hasRegoDeleteRule(operationID string, regoResources map[string]struct{}) bool {
	if len(regoResources) == 0 {
		return false
	}
	// Check if any Rego resource name is a substring of the operationId (case-insensitive match).
	// E.g., operationId "DeleteWorkflow" should match resource "workflow".
	opLower := toLower(operationID)
	for resource := range regoResources {
		if contains(opLower, toLower(resource)) {
			return true
		}
	}
	return false
}

// toLower is a simple ASCII lowercase helper to avoid importing strings.
func toLower(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'A' && c <= 'Z' {
			b[i] = c + 32
		}
	}
	return string(b)
}

// contains checks if s contains substr.
func contains(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(s) < len(substr) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

