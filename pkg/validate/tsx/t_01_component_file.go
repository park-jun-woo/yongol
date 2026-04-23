//ff:func feature=validate type=rule control=iteration dimension=1 topic=tsx
//ff:what T-1 — verifies that local component import paths resolve to an existing file
package tsx

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// t01ComponentFile validates T-1: for each PageSpec.Imports entry (local
// component imports only — npm packages are filtered out in the parser),
// the referenced file must exist on disk. WARNING level because AI agents
// iterate on TSX frequently and a transient missing file during scaffolding
// should not block the entire validate run.
func t01ComponentFile(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.TSXPages) == 0 {
		return nil
	}
	frontendRoot := filepath.Join(fs.SpecsDir, "frontend")
	aliasRoot := resolveAliasRoot(frontendRoot)

	var diags []diagnostic.Diagnostic
	for _, page := range fs.TSXPages {
		diags = append(diags, t01CheckPage(page, aliasRoot)...)
	}
	return diags
}
