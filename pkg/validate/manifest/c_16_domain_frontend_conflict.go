//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-structural
//ff:what C-16 — 도메인 frontend 경로가 단일 사이트 STML 루트("frontend")와 충돌 WARNING

package manifest

import (
	"path"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// singleSiteFrontendDir is the conventional STML source directory of a
// single-site project (specs/frontend). When `domains:` is declared the
// top-level frontend: block only supplies common settings; the STML pages move
// under each domain's own frontend directory. A domain whose frontend path
// still points at this legacy root would overlap the single-site location.
const singleSiteFrontendDir = "frontend"

// c16DomainFrontendConflict warns when a domain's frontend directory resolves
// to the same path as the legacy single-site STML root. It is a WARNING, not
// an ERROR: the project still builds, but two sources rendering from the same
// directory is almost always an authoring mistake (a domain frontend that was
// never moved out of the single-site root).
func c16DomainFrontendConflict(fs *yongol.Fullstack) []diagnostic.Diagnostic {
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
		dir := fs.Manifest.Domains[name].Frontend
		if dir == "" {
			continue // missing frontend is C-13's concern
		}
		if path.Clean(dir) != singleSiteFrontendDir {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[C-16] domains." + name + ".frontend=\"" + dir + "\" collides with the single-site STML root \"" + singleSiteFrontendDir + "\"",
			Advice:  "Move the domain's STML pages into a dedicated subdirectory (e.g. frontend/" + name + ") so domains do not share the legacy single-site root.",
		})
	}
	return diags
}
