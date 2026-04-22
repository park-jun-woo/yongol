//ff:func feature=gen-filefunc type=util control=sequence
//ff:what buildFeatureEntries — SSOT 후보 + internal 디렉토리 후보를 병합하여 feature 맵 반환
package filefunc

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildFeatureEntries returns the feature catalogue for the generated
// codebook. Keys are the subdirectory names under arts/backend/internal/
// merged with SSOT-derived feature names (so that SSaC @func packages that
// may not yet exist on disk are still captured). Descriptions are resolved
// in priority: SSOT → infra baseline → fallback.
func buildFeatureEntries(fs *yongol.Fullstack, backendDir string) map[string]string {
	out := map[string]string{}
	mergeSSOTFeatures(out, collectSSOTFeatures(fs))
	mergeInternalDirs(out, filepath.Join(backendDir, "internal"))
	ensureGenFilefuncEntry(out)
	return out
}
