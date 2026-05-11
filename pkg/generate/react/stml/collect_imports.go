//ff:func feature=stml-gen type=util control=iteration dimension=1 topic=import-collect
//ff:what PageSpec을 분석하여 필요한 임포트 목록을 결정한다
package stml

import (
	"os"
	"path/filepath"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// collectImports analyzes a PageSpec and determines required imports.
func collectImports(page stmlparser.PageSpec, specsDir string) importSet {
	is := importSet{react: true}
	compSet := map[string]bool{}

	if len(page.Fetches) > 0 {
		is.useQuery = true
	}
	if len(page.Actions) > 0 {
		is.useMutation = true
		is.useQueryClient = true
		is.useForm = true
	}

	for _, f := range page.Fetches {
		collectFetchImports(f, &is, compSet)
	}
	for _, a := range page.Actions {
		collectActionImports(a, &is, compSet)
	}

	for comp := range compSet {
		is.components = append(is.components, comp)
	}

	// Check for custom.ts
	if specsDir != "" {
		customPath := filepath.Join(specsDir, page.Name+".custom.ts")
		if _, err := os.Stat(customPath); err == nil {
			is.customFile = page.Name + ".custom"
		}
	}

	return is
}
