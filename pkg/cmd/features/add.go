//ff:func feature=features type=command control=iteration dimension=1
//ff:what RunAdd — features.yaml diff 후 신규 op 의 SSaC stub 생성 + 해시 갱신

package features

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// RunAdd compares the incoming features.yaml with the existing one in specs/,
// generates SSaC stubs for new ops, replaces specs/features.yaml, and updates
// the .yongol hash.
func RunAdd(out io.Writer, specsDir string, newFeaturesPath string) error {
	// 1. Parse new features.yaml.
	newFeats, err := loadFeaturesFile(newFeaturesPath)
	if err != nil {
		return fmt.Errorf("new features: %w", err)
	}

	// 2. Parse existing features.yaml.
	existingPath := filepath.Join(specsDir, "features.yaml")
	oldFeats, err := loadFeaturesFile(existingPath)
	if err != nil {
		return fmt.Errorf("existing features: %w", err)
	}

	// 3. Diff — extract new ops.
	diff := DiffOps(oldFeats, newFeats)
	if len(diff.Added) == 0 {
		fmt.Fprintln(out, "yongol features add: no new features found")
		return nil
	}

	// 4. Generate SSaC stubs for new ops (skip if file already exists).
	created := 0
	for _, f := range diff.Added {
		domain := extractDomain(f.Path)
		dir := filepath.Join(specsDir, "service", domain)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("mkdir service/%s: %w", domain, err)
		}
		dest := filepath.Join(dir, f.Op+".ssac")
		if _, err := os.Stat(dest); err == nil {
			fmt.Fprintf(out, "  skip %s (already exists)\n", dest)
			continue
		}
		content := fmt.Sprintf("package service\n\n// TODO: implement\nfunc %s() {}\n", f.Op)
		if err := os.WriteFile(dest, []byte(content), 0o644); err != nil {
			return fmt.Errorf("write ssac %s: %w", dest, err)
		}
		fmt.Fprintf(out, "  create %s\n", dest)
		created++
	}

	// 5. Replace specs/features.yaml with the new one.
	newData, err := os.ReadFile(newFeaturesPath)
	if err != nil {
		return fmt.Errorf("read new features: %w", err)
	}
	if err := os.WriteFile(existingPath, newData, 0o644); err != nil {
		return fmt.Errorf("write features.yaml: %w", err)
	}

	// 6. Update .yongol hash.
	if err := writeHash(specsDir, newData); err != nil {
		return err
	}

	fmt.Fprintf(out, "yongol features add: %d new feature(s), %d SSaC stub(s) created\n", len(diff.Added), created)
	return nil
}
