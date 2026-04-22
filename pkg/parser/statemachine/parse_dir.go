//ff:func feature=statemachine type=parser control=iteration dimension=1 topic=states
//ff:what 디렉토리 내 모든 .md 파일을 파싱하여 StateDiagram 슬라이스를 반환한다
package statemachine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseDir parses all *.md files in the given directory.
func ParseDir(dir string) ([]*StateDiagram, []diagnostic.Diagnostic) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, []diagnostic.Diagnostic{{
			File:    dir,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("read states dir: %v", err),
		}}
	}

	var diagrams []*StateDiagram
	var allDiags []diagnostic.Diagnostic
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		d, diags := ParseFile(filepath.Join(dir, e.Name()))
		if len(diags) > 0 {
			allDiags = append(allDiags, diags...)
			continue
		}
		diagrams = append(diagrams, d)
	}
	return diagrams, allDiags
}
