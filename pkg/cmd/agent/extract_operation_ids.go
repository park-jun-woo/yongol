//ff:func feature=agent type=helper control=iteration dimension=3
//ff:what extractOperationIDs — 진단 메시지에서 고유 operationId 추출

package agent

import (
	"regexp"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// reOperationID matches operationId references in diagnostic messages.
var reOperationID = regexp.MustCompile(
	`(?:operationId[:\s]+\"?([A-Z][A-Za-z0-9]+)\"?` +
		`|SSaC func ([A-Z][A-Za-z0-9]+)` +
		`|SSaC authorize \(([A-Z][A-Za-z0-9]+)` +
		`|input\.action == "([A-Z][A-Za-z0-9]+)"` +
		`|# ([A-Z][A-Za-z0-9]+))`,
)

// reMissingList matches "Missing: X, Y, Z" lists from XOH-11 style diagnostics.
var reMissingList = regexp.MustCompile(`Missing:\s*([A-Z][A-Za-z0-9]+(?:\s*,\s*[A-Z][A-Za-z0-9]+)*)`)

// reMissingItem extracts individual operationIds from a comma-separated list.
var reMissingItem = regexp.MustCompile(`[A-Z][A-Za-z0-9]+`)

// extractOperationIDs extracts unique operationIds from diagnostic messages.
func extractOperationIDs(diags []diagnostic.Diagnostic) []string {
	seen := map[string]struct{}{}
	var result []string
	addUnique := func(id string) {
		if _, ok := seen[id]; !ok {
			seen[id] = struct{}{}
			result = append(result, id)
		}
	}
	for _, d := range diags {
		matches := reOperationID.FindAllStringSubmatch(d.Message, -1)
		for _, m := range matches {
			for _, g := range m[1:] {
				if g != "" {
					addUnique(g)
				}
			}
		}
		if ml := reMissingList.FindStringSubmatch(d.Message); len(ml) > 1 {
			items := reMissingItem.FindAllString(ml[1], -1)
			for _, item := range items {
				addUnique(item)
			}
		}
	}
	return result
}
