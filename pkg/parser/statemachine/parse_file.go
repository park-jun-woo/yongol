//ff:func feature=statemachine type=parser control=sequence topic=states
//ff:what Mermaid stateDiagram 마크다운 파일 하나를 파싱한다
package statemachine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseFile parses a single Mermaid stateDiagram markdown file.
// The diagram ID is derived from the filename (without extension).
func ParseFile(path string) (*StateDiagram, []diagnostic.Diagnostic) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("read state file %s: %v", path, err),
		}}
	}

	id := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	return Parse(id, string(data), path)
}
