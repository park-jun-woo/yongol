//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-09 — data-component 파일이 존재하지 않음

package stml_openapi

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm09Components checks that each data-component references an existing
// .tsx component file.
func tm09Components(comps []stml.ComponentRef, file string, fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, c := range comps {
		diags = append(diags, tm09Component(c.Name, file, fs)...)
	}
	return diags
}

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
