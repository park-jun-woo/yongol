//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what loadFeatureLookup — features.yaml에서 operationId → Feature 맵 구성

package agent

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"gopkg.in/yaml.v3"
)

// loadFeatureLookup builds a map from operationId to Feature.
func loadFeatureLookup(specsDir string) map[string]features.Feature {
	path := filepath.Join(specsDir, "features.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var ff features.FeaturesFile
	if err := yaml.Unmarshal(data, &ff); err != nil {
		return nil
	}
	m := make(map[string]features.Feature, len(ff.Features))
	for _, f := range ff.Features {
		m[f.Op] = f
	}
	return m
}
