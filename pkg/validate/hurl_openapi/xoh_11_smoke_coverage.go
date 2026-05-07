//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-11 — smoke.hurl 이 OpenAPI 의 모든 operationId 를 호출해야 함

package hurl_openapi

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func xoh11SmokeCoverage(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}

	var smokeFile string
	for _, f := range fs.HurlFiles {
		if filepath.Base(f) == "smoke.hurl" {
			smokeFile = f
			break
		}
	}
	if smokeFile == "" {
		return nil
	}

	allOps := collectAllOperationIDs(fs.OpenAPIDoc)
	if len(allOps) == 0 {
		return nil
	}

	var smokeEntries []hurl.HurlEntry
	for _, e := range fs.HurlEntries {
		if e.File == smokeFile {
			smokeEntries = append(smokeEntries, e)
		}
	}

	routes := collectOpenAPIRoutes(fs.OpenAPIDoc)
	covered := collectCoveredOps(smokeEntries, routes)

	var missing []string
	for op := range allOps {
		if !covered[op] {
			missing = append(missing, op)
		}
	}
	if len(missing) == 0 {
		return nil
	}

	sort.Strings(missing)

	return []diagnostic.Diagnostic{{
		File:    smokeFile,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[XOH-11] smoke.hurl covers %d/%d endpoints. Missing: %s", len(covered), len(allOps), joinCSV(missing)),
		Advice:  "Add smoke requests for the missing operationIds in specs/tests/smoke.hurl",
	}}
}
