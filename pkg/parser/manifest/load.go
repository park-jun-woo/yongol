//ff:func feature=projectconfig type=loader control=sequence
//ff:what manifest.yaml 파일을 읽어 파싱하고 검증한 뒤 ProjectConfig를 반환한다
package manifest

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"gopkg.in/yaml.v3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

var reYAMLLine = regexp.MustCompile(`line (\d+)`)

// Load reads and parses manifest.yaml from the given specs directory root.
func Load(specsDir string) (*ProjectConfig, []diagnostic.Diagnostic) {
	path := filepath.Join(specsDir, "manifest.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "manifest.yaml not found: " + err.Error(),
		}}
	}

	var cfg ProjectConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		line := 0
		if m := reYAMLLine.FindStringSubmatch(err.Error()); len(m) == 2 {
			line, _ = strconv.Atoi(m[1])
		}
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    line,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "manifest.yaml parse error: " + err.Error(),
		}}
	}

	// Second pass: parse into yaml.Node to extract per-claim and per-role line numbers.
	claimLines, roleLines := extractAuthLines(data)

	// Post-process: convert RawClaims → Claims (ClaimDef) carrying line info.
	if cfg.Backend.Auth != nil && len(cfg.Backend.Auth.RawClaims) > 0 {
		cfg.Backend.Auth.Claims = parseRawClaims(cfg.Backend.Auth.RawClaims, claimLines)
	}
	if cfg.Backend.Auth != nil && len(cfg.Backend.Auth.Roles) > 0 {
		cfg.Backend.Auth.RolesLines = roleLines
	}
	if cfg.Backend.Auth != nil {
		cfg.Backend.Auth.UserTableLine = extractUserTableLine(data)
	}

	return &cfg, nil
}
