//ff:func feature=cli-init type=loader control=sequence
//ff:what loadFeatures — reads features.yaml, runs featcheck, returns FeaturesFile or error

package cliinit

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/cmd/featcheck"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// loadFeatures reads and parses features.yaml from the given absolute path,
// then runs FT-* validation. Returns the parsed FeaturesFile or an error
// if the file is unreadable, unparseable, or contains ERROR-level diagnostics.
func loadFeatures(path string) (*features.FeaturesFile, error) {
	ff, diags, err := featcheck.Run(path)
	if err != nil {
		return nil, err
	}

	var errs []string
	for _, d := range diags {
		if d.Level == diagnostic.LevelError {
			errs = append(errs, d.Message)
		}
	}
	if len(errs) > 0 {
		return nil, fmt.Errorf("features validation failed:\n  %s", strings.Join(errs, "\n  "))
	}

	return ff, nil
}
