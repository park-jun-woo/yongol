//ff:func feature=cli-init type=generator control=iteration dimension=1
//ff:what generateRegoFromFeatures — creates authz.rego with allow rule stubs from features

package cliinit

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// generateRegoFromFeatures creates specs/policy/authz.rego with one allow
// rule stub per feature. Each rule checks input.action against the feature op
// and input.resource against the derived domain.
func generateRegoFromFeatures(targetDir string, feats []features.Feature) error {
	var b strings.Builder
	b.WriteString("package authz\n\n")
	b.WriteString("default allow := false\n")

	for _, f := range feats {
		domain := extractDomain(f.Path)
		fmt.Fprintf(&b, "\nallow if {\n")
		fmt.Fprintf(&b, "    input.action == %q\n", f.Op)
		fmt.Fprintf(&b, "    input.resource == %q\n", domain)
		b.WriteString("}\n")
	}

	dest := filepath.Join(targetDir, "specs", "policy", "authz.rego")
	if err := os.WriteFile(dest, []byte(b.String()), 0o644); err != nil {
		return fmt.Errorf("write rego: %w", err)
	}
	return nil
}
