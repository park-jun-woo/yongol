//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-structural
//ff:what C-12 — domains 선언 시 각 도메인에 openapi 필드 필수 ERROR

package manifest

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c12DomainOpenAPIRequired enforces that every entry under the top-level
// `domains` key carries an `openapi` path. A domain without an OpenAPI spec
// has no API contract to generate routes / handlers from, so the domain would
// silently vanish from the generated backend. Failing fast at validate time
// keeps multi-site authors from discovering the empty domain at generate time.
func c12DomainOpenAPIRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || len(fs.Manifest.Domains) == 0 {
		return nil
	}
	names := make([]string, 0, len(fs.Manifest.Domains))
	for name := range fs.Manifest.Domains {
		names = append(names, name)
	}
	sort.Strings(names)

	var diags []diagnostic.Diagnostic
	for _, name := range names {
		if fs.Manifest.Domains[name].OpenAPI != "" {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[C-12] domains." + name + " is missing the required openapi field",
			Advice:  "Add an `openapi: <path>` to domains." + name + " pointing at the domain's OpenAPI spec.",
		})
	}
	return diags
}
