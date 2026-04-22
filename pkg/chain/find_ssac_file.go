//ff:func feature=chain type=util control=sequence
//ff:what findSSaCFile locates the SSaC source file for a service function.
package chain

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func findSSaCFile(sf *ssac.ServiceFunc, specsDir string) string {
	// Try feature-folder structure first.
	if sf.Feature != "" {
		rel := filepath.Join("service", sf.Feature, sf.FileName)
		if _, err := os.Stat(filepath.Join(specsDir, rel)); err == nil {
			return rel
		}
	}
	// Try flat structure.
	rel := filepath.Join("service", sf.FileName)
	if _, err := os.Stat(filepath.Join(specsDir, rel)); err == nil {
		return rel
	}
	return "service/" + sf.FileName
}
