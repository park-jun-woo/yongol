//ff:func feature=validate type=rule control=sequence topic=init-check
//ff:what INI-01 — specs/.yongol 미존재 시 yongol init 안내 WARN 출력

package initcheck

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run checks whether the .yongol file exists in the specs directory.
// If absent, it returns a single WARNING diagnostic advising the user
// to run yongol init.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.SpecsDir == "" {
		return nil
	}

	yongolPath := filepath.Join(fs.SpecsDir, ".yongol")
	if _, err := os.Stat(yongolPath); os.IsNotExist(err) {
		return []diagnostic.Diagnostic{{
			File:    ".yongol",
			Line:    0,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[INI-01] Project not initialized. Run: yongol init <project-id> <features.yaml>",
			Advice:  "See docs/features.md for features.yaml format and examples.",
		}}
	}

	return nil
}
