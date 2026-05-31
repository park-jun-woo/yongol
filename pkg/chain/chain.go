//ff:func feature=chain type=command control=iteration dimension=1
//ff:what Chain — operationId 기준 SSOT 노드 추적
package chain

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Chain traces all SSOT nodes connected to the given operationId.
func Chain(fs *yongol.Fullstack, operationID string) ([]Link, error) {
	if fs.OpenAPIDoc == nil {
		return nil, fmt.Errorf("openapi not parsed")
	}
	specsDir := fs.SpecsDir

	var links []Link

	openapiLink := traceOpenAPI(fs.OpenAPIDoc, operationID, specsDir)
	if openapiLink == nil {
		return nil, fmt.Errorf("operationId %q not found in OpenAPI", operationID)
	}
	links = append(links, *openapiLink)

	var matched *ssac.ServiceFunc
	for i := range fs.ServiceFuncs {
		if fs.ServiceFuncs[i].Name == operationID {
			matched = &fs.ServiceFuncs[i]
			break
		}
	}

	if matched != nil {
		links = append(links, traceSSaC(matched, specsDir))
		links = append(links, traceDDL(matched, fs.DDLTables, specsDir)...)
		links = append(links, tracePolicy(matched, fs.ParsedPolicies, specsDir)...)
		links = append(links, traceStates(matched, fs.StateDiagrams, specsDir)...)
		allFuncs := append([]funcspec.FuncSpec{}, fs.ProjectFuncSpecs...)
		allFuncs = append(allFuncs, fs.YongolPkgSpecs...)
		links = append(links, traceFuncSpecs(matched, allFuncs, specsDir)...)
	}

	links = append(links, traceHurl(operationID, fs.OpenAPIDoc, filepath.Join(specsDir, "tests"), specsDir)...)

	return links, nil
}
