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

	dirs := directorySSOTs(abs)
	for _, d := range dirs {
		entry, err := detectDirSSOT(d)
		if err != nil {
			return nil, err
		}
		if entry.Presence == SSOTAbsent {
			continue
		}
		found = append(found, entry)
	}

	return found, nil
}
