//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what tm09Component — 단일 data-component 참조의 .tsx 파일 존재 확인

package stml_openapi

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm09Component checks a single component reference.
func tm09Component(name, file string, fs *yongol.Fullstack) []diagnostic.Diagnostic {
	compPath := filepath.Join(fs.SpecsDir, "frontend", "components", name+".tsx")
	if _, err := os.Stat(compPath); os.IsNotExist(err) {
		relPath := filepath.Join("frontend", "components", name+".tsx")
		return []diagnostic.Diagnostic{{
			File:    file,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-09] data-component %q references %s which does not exist", name, relPath),
			Advice:  fmt.Sprintf("Create the component file at %s, or remove the data-component attribute from the STML file", relPath),
		}}
	}
	return nil
}
