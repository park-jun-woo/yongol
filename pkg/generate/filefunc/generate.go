//ff:func feature=gen-filefunc type=generator control=sequence
//ff:what Generate — 백엔드 아티팩트 디렉토리에 codebook.yaml 자동 생성
package filefunc

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate writes arts/backend/codebook.yaml for the given Fullstack. The
// file contains a required section (feature + type) populated from the
// backend's on-disk package layout and SSOT metadata, plus an optional
// section with the baseline topic / ssot / pattern keys. Must be called
// after all backend artifacts have been emitted so internal/ is final.
func Generate(fs *yongol.Fullstack, artifactsDir string) error {
	backendDir := filepath.Join(artifactsDir, "backend")
	book := Codebook{
		Required: Required{
			Feature: buildFeatureEntries(fs, backendDir),
			Type:    buildTypeEntries(),
		},
		Optional: Optional{
			Topic:   buildTopicEntries(fs),
			SSOT:    defaultSSOTEntries(),
			Pattern: defaultPatternEntries(),
		},
	}
	return writeCodebook(filepath.Join(backendDir, "codebook.yaml"), book)
}
