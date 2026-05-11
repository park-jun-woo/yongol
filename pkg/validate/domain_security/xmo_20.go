//ff:func feature=validate type=rule control=iteration dimension=2 topic=domain-security
//ff:what XMO-20 — public OpenAPI operationId가 public STML에서 미소비되면 ERROR
package domain_security

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xmo20PublicUnconsumed detects public domain OpenAPI operationIds that are
// not consumed by any STML page in the public frontend directory.
func xmo20PublicUnconsumed(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	docs := loadDomainOpenAPIDocs(fs)

	var publicDoc *domainDoc
	for i, dd := range docs {
		if dd.Name == "public" {
			publicDoc = &docs[i]
			break
		}
	}
	if publicDoc == nil || publicDoc.Doc.Paths == nil {
		return nil
	}
	if len(fs.STMLPages) == 0 {
		return nil
	}

	// Filter STML pages belonging to the public frontend directory.
	publicPages := filterPagesByDomain(fs.STMLPages, publicDoc.Cfg.Frontend)
	consumed := collectConsumedOpsFromPages(publicPages)

	var diags []diagnostic.Diagnostic
	for _, item := range publicDoc.Doc.Paths.Map() {
		ops := []*openapi3.Operation{
			item.Get, item.Post, item.Put, item.Delete, item.Patch,
		}
		for _, op := range ops {
			if op == nil || op.OperationID == "" {
				continue
			}
			// Skip auth endpoints.
			if hasEmptySecurity(op) {
				continue
			}
			if _, ok := consumed[op.OperationID]; !ok {
				diags = append(diags, diagnostic.Diagnostic{
					File:    publicDoc.Cfg.OpenAPI,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[XMO-20] public operationId %q is not consumed by any STML page in the public frontend", op.OperationID),
					Advice:  fmt.Sprintf("Add a data-fetch or data-action referencing %q in a public STML page, or remove the endpoint from the public OpenAPI", op.OperationID),
				})
			}
		}
	}
	return diags
}

// filterPagesByDomain returns STML pages whose FileName starts with or is
// relative to the given frontend directory path.
func filterPagesByDomain(pages []stml.PageSpec, frontendDir string) []stml.PageSpec {
	if frontendDir == "" {
		return nil
	}
	// Normalize: ensure no trailing slash for prefix matching.
	prefix := strings.TrimSuffix(frontendDir, "/")

	var result []stml.PageSpec
	for _, p := range pages {
		// FileName may be relative or absolute; match by prefix or base dir.
		dir := filepath.Dir(p.FileName)
		if strings.HasPrefix(p.FileName, prefix) || strings.HasPrefix(dir, prefix) || dir == prefix {
			result = append(result, p)
		}
	}
	return result
}

// collectConsumedOpsFromPages collects operationIds from fetch and action blocks.
func collectConsumedOpsFromPages(pages []stml.PageSpec) map[string]struct{} {
	out := make(map[string]struct{})
	for _, page := range pages {
		for _, f := range page.Fetches {
			collectFetchOpsRecursive(f, out)
		}
		for _, a := range page.Actions {
			out[a.OperationID] = struct{}{}
		}
	}
	return out
}

// collectFetchOpsRecursive recursively collects operationIds from a FetchBlock.
func collectFetchOpsRecursive(f stml.FetchBlock, out map[string]struct{}) {
	out[f.OperationID] = struct{}{}
	for _, child := range f.NestedFetches {
		collectFetchOpsRecursive(child, out)
	}
}
