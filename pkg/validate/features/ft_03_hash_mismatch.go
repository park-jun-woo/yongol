//ff:func feature=validate type=rule control=sequence topic=features-structural
//ff:what FT-03 — features.yaml 해시가 specs/.yongol과 일치하지 않으면 에러

package features

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"

	"gopkg.in/yaml.v3"
)

// ft03HashMismatch validates FT-03: when features.yaml exists, specs/.yongol
// must also exist and the SHA-256 hash must match. This detects unintended
// modifications to features.yaml after the baseline was established by
// `yongol init`.
func ft03HashMismatch(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.SpecsDir == "" {
		return nil
	}

	featuresPath := filepath.Join(fs.SpecsDir, "features.yaml")
	if _, err := os.Stat(featuresPath); os.IsNotExist(err) {
		// features.yaml absent → skip (no project features).
		return nil
	}

	yongolPath := filepath.Join(fs.SpecsDir, ".yongol")
	yongolData, err := os.ReadFile(yongolPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []diagnostic.Diagnostic{{
				File:    ".yongol",
				Line:    0,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[FT-03] specs/.yongol not found. Run: yongol init <project> <features.yaml>",
				Advice:  "Run yongol init to generate the hash lock file from features.yaml",
			}}
		}
		return []diagnostic.Diagnostic{{
			File:    ".yongol",
			Line:    0,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[FT-03] cannot read specs/.yongol: %v", err),
		}}
	}

	var yf yongolFile
	if err := yaml.Unmarshal(yongolData, &yf); err != nil {
		return []diagnostic.Diagnostic{{
			File:    ".yongol",
			Line:    0,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[FT-03] cannot parse specs/.yongol: %v", err),
		}}
	}

	storedHash, ok := yf.Hashes["features.yaml"]
	if !ok {
		return []diagnostic.Diagnostic{{
			File:    ".yongol",
			Line:    0,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[FT-03] specs/.yongol missing features.yaml hash entry",
			Advice:  "Run yongol init to regenerate the hash lock file",
		}}
	}

	featuresData, err := os.ReadFile(featuresPath)
	if err != nil {
		return []diagnostic.Diagnostic{{
			File:    "features.yaml",
			Line:    0,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[FT-03] cannot read features.yaml: %v", err),
		}}
	}

	hash := sha256.Sum256(featuresData)
	actual := fmt.Sprintf("sha256:%x", hash)

	if !strings.EqualFold(storedHash, actual) {
		return []diagnostic.Diagnostic{{
			File:    "features.yaml",
			Line:    0,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[FT-03] features.yaml was modified after baseline",
			Advice:  "Re-run yongol init to update the hash lock, or revert features.yaml changes",
		}}
	}

	return nil
}
