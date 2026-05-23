//ff:func feature=agent type=helper control=sequence
//ff:what resolveFeature — 파일 경로에서 operationId/feature desc 추출

package agent

import (
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// resolveFeature extracts operationId and feature desc for a file.
// For SSaC files, the operationId is the filename stem.
func resolveFeature(relPath string, l layer, lookup map[string]features.Feature) (desc, path string) {
	if l == layerSSaC {
		base := filepath.Base(relPath)
		op := strings.TrimSuffix(base, ".ssac")
		if f, ok := lookup[op]; ok {
			return f.Desc, f.Path
		}
		return op + " (no desc)", ""
	}
	return "", ""
}
