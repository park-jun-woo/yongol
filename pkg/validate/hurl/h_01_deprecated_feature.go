//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-structural
//ff:what H-1 — tests/*.feature 파일 존재 감지 (deprecated)

package hurl

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// h01DeprecatedFeature flags any *.feature file under tests/ or scenario/.
// .feature is deprecated — scenarios must be authored in .hurl format.
func h01DeprecatedFeature(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, sub := range []string{"tests", "scenario"} {
		matches, _ := filepath.Glob(filepath.Join(fs.SpecsDir, sub, "*.feature"))
		for _, m := range matches {
			diags = append(diags, diagnostic.Diagnostic{
				File:    m,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[H-1] .feature files are deprecated",
				Advice:  ".feature 파일을 tests/scenario-*.hurl 형식으로 재작성하세요",
			})
		}
	}
	return diags
}
