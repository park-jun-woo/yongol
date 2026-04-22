//ff:func feature=chain type=util control=iteration dimension=1
//ff:what traceSSaC locates the SSaC service function and summarizes its sequence types.
package chain

import (
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func traceSSaC(sf *ssac.ServiceFunc, specsDir string) Link {
	var seqTypes []string
	seen := map[string]bool{}
	for _, seq := range sf.Sequences {
		tag := "@" + seq.Type
		if !seen[tag] {
			seqTypes = append(seqTypes, tag)
			seen[tag] = true
		}
	}

	relPath := findSSaCFile(sf, specsDir)
	line := 0
	if relPath != "" {
		line = grepLine(filepath.Join(specsDir, relPath), "func "+sf.Name)
	}

	return Link{
		Kind:    "SSaC",
		File:    relPath,
		Line:    line,
		Summary: strings.Join(seqTypes, " "),
	}
}
