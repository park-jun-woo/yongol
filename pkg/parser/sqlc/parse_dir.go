//ff:func feature=orchestrator type=parser control=iteration dimension=1
//ff:what ParseDir — 디렉토리 내 모든 .sql 파일을 sqlc 쿼리로 파싱
package sqlc

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseDir walks every .sql file directly under dir (non-recursive) and
// returns their aggregated QuerySpecs. A missing directory is not an error;
// callers that require a queries/ directory can check for an empty result.
func ParseDir(dir string) ([]QuerySpec, []diagnostic.Diagnostic) {
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
			Message: "cannot read sqlc query directory: " + err.Error(),
		}}
	}

	var all []QuerySpec
	var diags []diagnostic.Diagnostic
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}
		path := filepath.Join(dir, e.Name())
		specs, d := ParseFile(path)
		all = append(all, specs...)
		diags = append(diags, d...)
	}
	return all, diags
}
