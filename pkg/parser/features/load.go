//ff:func feature=features type=loader control=sequence
//ff:what features.yaml 파일을 읽어 파싱하고 Feature 슬라이스를 반환한다
package features

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"gopkg.in/yaml.v3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

var reYAMLLine = regexp.MustCompile(`line (\d+)`)

// Load reads and parses features.yaml from the given specs directory root.
func Load(specsDir string) ([]Feature, []diagnostic.Diagnostic) {
	path := filepath.Join(specsDir, "features.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "features.yaml not found: " + err.Error(),
		}}
	}

	var ff FeaturesFile
	if err := yaml.Unmarshal(data, &ff); err != nil {
		line := 0
		if m := reYAMLLine.FindStringSubmatch(err.Error()); len(m) == 2 {
			line, _ = strconv.Atoi(m[1])
		}
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    line,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "features.yaml parse error: " + err.Error(),
		}}
	}

	// Second pass: extract per-feature line numbers from yaml.Node tree.
	lines := extractFeatureLines(data)

	var diags []diagnostic.Diagnostic
	for i := range ff.Features {
		if i < len(lines) {
			ff.Features[i].Line = lines[i]
		}
		if ff.Features[i].Op == "" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    path,
				Line:    ff.Features[i].Line,
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: "features.yaml: missing required field 'op'",
			})
		}
		if ff.Features[i].Path == "" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    path,
				Line:    ff.Features[i].Line,
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: "features.yaml: missing required field 'path'",
			})
		}
		if ff.Features[i].Desc == "" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    path,
				Line:    ff.Features[i].Line,
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: "features.yaml: missing required field 'desc'",
			})
		}
	}

	if len(diags) > 0 {
		return nil, diags
	}
	return ff.Features, nil
}
