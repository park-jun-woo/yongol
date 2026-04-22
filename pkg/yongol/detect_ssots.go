//ff:func feature=orchestrator type=command control=iteration dimension=1
//ff:what specs 루트에서 SSOT 파일/디렉토리를 Presence 3-상태로 탐지
package yongol

import (
	"fmt"
	"os"
	"path/filepath"
)

// DetectSSOTs scans root for known SSOT directories/files. Each detected
// entry carries a Presence: SSOTDeclared (directory exists but no matching
// content) or SSOTPopulated (content file found). SSOTAbsent entries are
// omitted. Single-file SSOTs (manifest.yaml, api/openapi.yaml) are only
// emitted when the file is present (equivalent to SSOTPopulated).
func DetectSSOTs(root string) ([]DetectedSSOT, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(abs)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("not a directory: %s", abs)
	}

	var found []DetectedSSOT

	configPath := filepath.Join(abs, "manifest.yaml")
	if _, err := os.Stat(configPath); err == nil {
		found = append(found, DetectedSSOT{Kind: KindConfig, Path: configPath, Presence: SSOTPopulated})
	}
	openapiPath := filepath.Join(abs, "api", "openapi.yaml")
	if _, err := os.Stat(openapiPath); err == nil {
		found = append(found, DetectedSSOT{Kind: KindOpenAPI, Path: openapiPath, Presence: SSOTPopulated})
	}

	type dirSSOT struct {
		kind  SSOTKind
		dir   string
		globs []string
	}
	dirs := []dirSSOT{
		{KindDDL, filepath.Join(abs, "db"), []string{"*.sql"}},
		{KindSSaC, filepath.Join(abs, "service"), []string{"*.ssac", "*/*.ssac"}},
		{KindStates, filepath.Join(abs, "states"), []string{"*.md"}},
		{KindPolicy, filepath.Join(abs, "policy"), []string{"*.rego"}},
		{KindScenario, filepath.Join(abs, "tests"), []string{"scenario-*.hurl", "invariant-*.hurl"}},
		{KindFunc, filepath.Join(abs, "func"), []string{"*/*.go"}},
		{KindTSX, filepath.Join(abs, "frontend"), []string{"*.tsx", "*/*.tsx", "*/*/*.tsx", "*/*/*/*.tsx"}},
	}
	for _, d := range dirs {
		count := 0
		for _, g := range d.globs {
			pattern := filepath.Join(d.dir, g)
			matches, err := filepath.Glob(pattern)
			if err != nil {
				// filepath.Glob only returns ErrBadPattern (syntax error).
				// Patterns are hard-coded so this is effectively unreachable,
				// but surface it as a diagnostic to prevent silent pass.
				return nil, fmt.Errorf("detect SSOTs glob failed for %s: %w", pattern, err)
			}
			count += len(matches)
		}
		p := dirPresence(d.dir, count)
		if p == SSOTAbsent {
			continue
		}
		found = append(found, DetectedSSOT{Kind: d.kind, Path: d.dir, Presence: p})
	}

	return found, nil
}
