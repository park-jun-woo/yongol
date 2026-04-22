//ff:func feature=validate type=rule control=sequence topic=manifest-structural
//ff:what C-4 — manifest metadata.name 필수 값 검증

package manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c04MetadataName validates that manifest.metadata.name is non-empty.
func c04MetadataName(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Manifest == nil {
		return nil
	}
	if strings.TrimSpace(fs.Manifest.Metadata.Name) != "" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[C-4] metadata.name is empty",
		Advice:  "metadata.name 에 프로젝트 이름을 지정하세요",
	}}
}
