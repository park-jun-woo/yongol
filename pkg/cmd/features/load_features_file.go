//ff:func feature=features type=loader control=iteration dimension=1
//ff:what loadFeaturesFile — 임의 경로의 features.yaml 읽기 + 파싱 + 필수 필드 검증

package features

import (
	"fmt"
	"os"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"gopkg.in/yaml.v3"
)

// loadFeaturesFile reads and parses a features.yaml from an arbitrary path.
func loadFeaturesFile(path string) ([]featparser.Feature, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read features: %w", err)
	}

	var ff featparser.FeaturesFile
	if err := yaml.Unmarshal(data, &ff); err != nil {
		return nil, fmt.Errorf("parse features: %w", err)
	}

	if len(ff.Features) == 0 {
		return nil, fmt.Errorf("features.yaml contains no features")
	}

	for i, f := range ff.Features {
		if f.Op == "" {
			return nil, fmt.Errorf("features[%d]: missing required field 'op'", i)
		}
		if f.Path == "" {
			return nil, fmt.Errorf("features[%d]: missing required field 'path'", i)
		}
		if f.Desc == "" {
			return nil, fmt.Errorf("features[%d]: missing required field 'desc'", i)
		}
	}

	return ff.Features, nil
}
