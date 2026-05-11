//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-12 — manifest.frontend.defaultLayout 값이 layouts/에 없음 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm12DefaultLayoutNotFound checks that the manifest's DefaultLayout value
// references an existing layout defined in Layouts.
func tm12DefaultLayoutNotFound(defaultLayout string, layouts []stml.LayoutSpec) []diagnostic.Diagnostic {
	if defaultLayout == "" {
		return nil
	}

	for _, l := range layouts {
		if l.Name == defaultLayout {
			return nil
		}
	}

	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[TM-12] manifest.frontend.defaultLayout %q not found in layouts/", defaultLayout),
		Advice:  fmt.Sprintf("Create layouts/%s.html or fix the defaultLayout value in manifest.yaml", defaultLayout),
	}}
}
