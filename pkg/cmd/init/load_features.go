//ff:func feature=cli-init type=loader control=sequence
//ff:what loadFeatures — reads and parses features.yaml from an arbitrary path

package cliinit

import (
	"fmt"
	"os"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"gopkg.in/yaml.v3"
)

// loadFeatures reads and parses features.yaml from the given absolute path.
// Returns the parsed features slice or an error on read/parse failure.
func loadFeatures(path string) ([]features.Feature, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read features: %w", err)
	}

	var ff features.FeaturesFile
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
