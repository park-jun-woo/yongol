//ff:func feature=cli-init type=generator control=iteration dimension=1
//ff:what generateHurlFromFeatures — creates tests/smoke.hurl with request stubs from features

package cliinit

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// generateHurlFromFeatures creates specs/tests/smoke.hurl with one smoke
// request stub per feature.
func generateHurlFromFeatures(targetDir string, feats []features.Feature) error {
	var b strings.Builder

	for i, f := range feats {
		route, err := parseHTTPPath(f.Path)
		if err != nil {
			return err
		}
		if i > 0 {
			b.WriteString("\n")
		}
		fmt.Fprintf(&b, "# %s\n", f.Op)
		fmt.Fprintf(&b, "%s {{host}}%s\n", strings.ToUpper(route.Method), route.URI)
		b.WriteString("Authorization: Bearer {{token}}\n")
		b.WriteString("HTTP 200\n")
	}

	dest := filepath.Join(targetDir, "specs", "tests", "smoke.hurl")
	if err := os.WriteFile(dest, []byte(b.String()), 0o644); err != nil {
		return fmt.Errorf("write hurl: %w", err)
	}
	return nil
}
