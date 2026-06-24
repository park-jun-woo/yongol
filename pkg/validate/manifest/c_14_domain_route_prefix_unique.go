//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-structural
//ff:what C-14 — 도메인 route_prefix 중복 불허 ERROR

package manifest

import (
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c14DomainRoutePrefixUnique rejects two domains that mount under the same
// backend route group prefix. Distinct domains must occupy distinct URL
// namespaces; a shared prefix means their operations collide on the same Gin
// route group and the second registration shadows or panics the first. Empty
// prefixes are ignored here — their absence is unrelated to uniqueness and is
// not part of this rule's contract.
func c14DomainRoutePrefixUnique(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || len(fs.Manifest.Domains) == 0 {
		return nil
	}
	// Map each non-empty prefix to the domain names that declare it.
	byPrefix := make(map[string][]string)
	for name, cfg := range fs.Manifest.Domains {
		if cfg.RoutePrefix == "" {
			continue
		}
		byPrefix[cfg.RoutePrefix] = append(byPrefix[cfg.RoutePrefix], name)
	}

	prefixes := make([]string, 0, len(byPrefix))
	for prefix := range byPrefix {
		prefixes = append(prefixes, prefix)
	}
	sort.Strings(prefixes)

	var diags []diagnostic.Diagnostic
	for _, prefix := range prefixes {
		owners := byPrefix[prefix]
		if len(owners) < 2 {
			continue
		}
		sort.Strings(owners)
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[C-14] route_prefix \"" + prefix + "\" is declared by multiple domains: " + strings.Join(owners, ", "),
			Advice:  "Give each domain a unique route_prefix so their route groups do not collide.",
		})
	}
	return diags
}
