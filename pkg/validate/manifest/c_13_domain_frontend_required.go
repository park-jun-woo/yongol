//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-structural
//ff:what C-13 — domains 선언 시 각 도메인에 frontend 필드 필수 ERROR

package manifest

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c13DomainFrontendRequired enforces that every entry under the top-level
// `domains` key carries a `frontend` path. Each domain is an independent app
// whose STML source location is set by its own frontend directory; without it
// the generator has no pages to render for the domain and the domain's index
// route cannot be resolved. Failing fast at validate time avoids an empty
// frontend build downstream.
func c13DomainFrontendRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
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
		if fs.Manifest.Domains[name].Frontend != "" {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[C-13] domains." + name + " is missing the required frontend field",
			Advice:  "Add a `frontend: <dir>` to domains." + name + " pointing at the domain's STML source directory.",
		})
	}
	return diags
}
