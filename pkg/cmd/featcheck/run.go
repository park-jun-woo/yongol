//ff:func feature=cli-featcheck type=command control=sequence
//ff:what Run — features.yaml 파싱 + 필수 필드 확인 + FT-* 내부 검증 실행

package featcheck

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
	valfeatures "github.com/park-jun-woo/yongol/pkg/validate/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run reads and parses features.yaml from featuresPath, checks required
// fields, and runs FT-* validation rules. It returns the parsed FeaturesFile,
// any diagnostics, and an error if the file cannot be read or parsed.
func Run(featuresPath string) (*features.FeaturesFile, []diagnostic.Diagnostic, error) {
	data, err := os.ReadFile(featuresPath)
	if err != nil {
		return nil, nil, fmt.Errorf("read features: %w", err)
	}

	var ff features.FeaturesFile
	if err := yaml.Unmarshal(data, &ff); err != nil {
		return nil, nil, fmt.Errorf("parse features: %w", err)
	}

	if len(ff.Features) == 0 {
		return nil, nil, fmt.Errorf("features.yaml contains no features")
	}

	// Required field check.
	for i, f := range ff.Features {
		if f.Op == "" {
			return nil, nil, fmt.Errorf("features[%d]: missing required field 'op'", i)
		}
		if f.Path == "" {
			return nil, nil, fmt.Errorf("features[%d]: missing required field 'path'", i)
		}
		if f.Desc == "" {
			return nil, nil, fmt.Errorf("features[%d]: missing required field 'desc'", i)
		}
	}

	// Build a minimal Fullstack with Features and FeatureTables only.
	// SpecsDir is left empty so FT-03 (hash mismatch) naturally skips.
	fs := &yongol.Fullstack{
		Features:      ff.Features,
		FeatureTables: ff.Tables,
	}

	diags := valfeatures.Run(fs)
	return &ff, diags, nil
}
