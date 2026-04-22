//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what parseDirEntry — parses the .ssac files for a single directory entry and assigns the feature
package ssac

import (
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// parseDirEntry parses a single .ssac file and assigns the feature.
func parseDirEntry(dir, path, name string) ([]ServiceFunc, []diagnostic.Diagnostic) {
	sfs, diags := ParseFile(path)
	if len(diags) > 0 {
		return nil, diags
	}
	rel, _ := filepath.Rel(dir, path)
	if filepath.Dir(rel) == "." {
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: name + " — SSaC files must not be placed directly in service/. Use a feature subfolder (e.g. service/auth/" + name + ")",
		}}
	}
	for i := range sfs {
		parts := strings.Split(filepath.Dir(rel), string(filepath.Separator))
		sfs[i].Feature = parts[0]
	}
	return sfs, nil
}
