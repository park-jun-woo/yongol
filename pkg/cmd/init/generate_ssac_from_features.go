//ff:func feature=cli-init type=generator control=iteration dimension=1
//ff:what generateSSaCFromFeatures — creates SSaC stub files grouped by domain from features

package cliinit

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// generateSSaCFromFeatures creates one SSaC stub file per feature under
// <targetDir>/specs/service/{domain}/{op}.ssac. Domains are derived from the
// first URI path segment.
func generateSSaCFromFeatures(targetDir string, feats []features.Feature) error {
	for _, f := range feats {
		domain := extractDomain(f.Path)
		dir := filepath.Join(targetDir, "specs", "service", domain)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("mkdir service/%s: %w", domain, err)
		}
		dest := filepath.Join(dir, f.Op+".ssac")
		content := fmt.Sprintf("package service\n\n// TODO: implement\nfunc %s() {}\n", f.Op)
		if err := os.WriteFile(dest, []byte(content), 0o644); err != nil {
			return fmt.Errorf("write ssac %s: %w", dest, err)
		}
	}
	return nil
}
