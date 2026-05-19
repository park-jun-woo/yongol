//ff:func feature=features type=command control=sequence
//ff:what RunRemove — features.yaml 에서 operationId 삭제 + SSaC 파일 삭제 + 해시 갱신

package features

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"gopkg.in/yaml.v3"
)

// RunRemove removes the specified operationIds from specs/features.yaml,
// deletes corresponding SSaC stub files, and updates the .yongol hash.
// When yes is false, it prints the deletion plan and asks for confirmation.
func RunRemove(out io.Writer, in io.Reader, specsDir string, ops []string, yes bool) error {
	if len(ops) == 0 {
		return fmt.Errorf("at least one operationId is required")
	}

	// 1. Parse existing features.yaml.
	existingPath := filepath.Join(specsDir, "features.yaml")
	oldFeats, err := loadFeaturesFile(existingPath)
	if err != nil {
		return fmt.Errorf("existing features: %w", err)
	}

	// Build lookup of ops to remove.
	removeSet := make(map[string]bool, len(ops))
	for _, op := range ops {
		removeSet[op] = true
	}

	// Validate all ops exist.
	existingOps := make(map[string]bool, len(oldFeats))
	for _, f := range oldFeats {
		existingOps[f.Op] = true
	}
	for _, op := range ops {
		if !existingOps[op] {
			return fmt.Errorf("operationId %q not found in features.yaml", op)
		}
	}

	// Collect features to remove (for display and SSaC deletion).
	var toRemove []featparser.Feature
	for _, f := range oldFeats {
		if removeSet[f.Op] {
			toRemove = append(toRemove, f)
		}
	}

	// 2. Confirmation (unless --yes).
	if !yes {
		fmt.Fprintln(out, "The following features will be removed:")
		for _, f := range toRemove {
			domain := extractDomain(f.Path)
			ssacPath := filepath.Join("service", domain, f.Op+".ssac")
			fmt.Fprintf(out, "  - %s (%s) → %s\n", f.Op, f.Path, ssacPath)
		}
		fmt.Fprint(out, "\nContinue? [y/N] ")
		scanner := bufio.NewScanner(in)
		scanner.Scan()
		answer := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if answer != "y" && answer != "yes" {
			fmt.Fprintln(out, "aborted")
			return nil
		}
	}

	// 3. Delete SSaC files.
	deleted := 0
	for _, f := range toRemove {
		domain := extractDomain(f.Path)
		ssacPath := filepath.Join(specsDir, "service", domain, f.Op+".ssac")
		if err := os.Remove(ssacPath); err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintf(out, "  skip %s (not found)\n", ssacPath)
				continue
			}
			return fmt.Errorf("remove %s: %w", ssacPath, err)
		}
		fmt.Fprintf(out, "  delete %s\n", ssacPath)
		deleted++
	}

	// 4. Rebuild features.yaml without removed ops.
	var remaining []featparser.Feature
	for _, f := range oldFeats {
		if !removeSet[f.Op] {
			remaining = append(remaining, f)
		}
	}

	ff := featparser.FeaturesFile{Features: remaining}
	data, err := yaml.Marshal(&ff)
	if err != nil {
		return fmt.Errorf("marshal features.yaml: %w", err)
	}
	if err := os.WriteFile(existingPath, data, 0o644); err != nil {
		return fmt.Errorf("write features.yaml: %w", err)
	}

	// 5. Update .yongol hash.
	if err := writeHash(specsDir, data); err != nil {
		return err
	}

	fmt.Fprintf(out, "yongol features remove: %d feature(s) removed, %d SSaC file(s) deleted\n", len(toRemove), deleted)
	return nil
}
