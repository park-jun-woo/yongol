//ff:func feature=gen-gogin type=command control=iteration dimension=2
//ff:what Generate — SSaC → StrictServerInterface Server struct + method files (1파일 1func 전면)

package ssac

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate produces all service-layer artifacts from SSaC:
//   - internal/service/server.go (Server struct only)
//   - internal/service/{ptr_of,deref_*}.go (pointer helpers, 1 file 1 func)
//   - internal/service/convert_<name>.go / convert_<name>_list.go per 200-response schema
//   - internal/service/<func_name>.go per SSaC function (StrictServerInterface method)
//   - Subscribe methods (not part of StrictServerInterface, registered via queue.Subscribe)
//
// Every emitted file holds exactly one top-level func (Phase004: 1-file-1-func
// applied uniformly, replacing the Phase003 POC gating that limited the split
// to ActivateWorkflow's dependency tree).
//
// apiSuffix / funcPrefix carry the per-domain emission context (Phase007). In
// single-site mode both are "" — the api import stays "<module>/internal/api"
// and converters keep their plain convert<Name> name. In domain mode the caller
// passes sanitizeDomainName(name) / domainTitle(name) and invokes Generate once
// per domain over fs.DomainView(name); the operationId membership filter below
// then emits each method exactly once, by its owning domain.
func Generate(fs *yongol.Fullstack, artifactsDir, apiSuffix, funcPrefix string) error {
	if len(fs.ServiceFuncs) == 0 {
		return nil
	}
	dg := domainGen{ApiSuffix: apiSuffix, FuncPrefix: funcPrefix}
	modulePath := ""
	if fs.Manifest != nil {
		modulePath = fs.Manifest.Backend.Module
	}
	serviceDir := filepath.Join(artifactsDir, "backend", "internal", "service")
	if err := os.MkdirAll(serviceDir, 0o755); err != nil {
		return err
	}
	if err := generateServerGo(fs, artifactsDir, modulePath); err != nil {
		return fmt.Errorf("server.go: %w", err)
	}
	if err := generateServerHelpers(artifactsDir); err != nil {
		return fmt.Errorf("server helpers: %w", err)
	}

	// All converters (db row → api DTO) are emitted as individual
	// 1-file-1-func files. The legacy bundled converters.go is no longer
	// written; a stale converters.go from a previous run is swept below.
	//
	// `needed` is the set of OpenAPI schemas referenced from success
	// responses. We intersect it with `sqlcModelNames` (DDL tables mapped
	// to sqlc struct names) so we only emit convert<X> for schemas that
	// actually have a sqlc row type — otherwise an api-only wrapper like
	// ExecuteWorkflowResponse becomes a phantom convertExecuteWorkflowResponse
	// that references the non-existent db.ExecuteWorkflowResponse.
	usedNames := make(map[string]bool)
	needed := collectResponseSchemas(fs.OpenAPIDoc)
	sqlcModelNames := sqlcModelNameSet(fs)
	filtered := make(map[string]bool, len(needed))
	for name := range needed {
		if sqlcModelNames[name] {
			filtered[name] = true
		}
	}
	if err := emitAllConverterFiles(fs.OpenAPIDoc, serviceDir, modulePath, filtered, fs.DDLTables, usedNames, dg); err != nil {
		return fmt.Errorf("converters: %w", err)
	}

	// Func Response converters — same pattern as DB model converters but
	// for @call result types that have no sqlc backing row. Intersect
	// needed with funcRespNames (excluding sqlcModelNames) so we only
	// emit convert<X> for schemas produced by Func Response types.
	funcRespNames := collectFuncResponseNames(fs.ServiceFuncs)
	funcFiltered := make(map[string]funcRespInfo)
	for name := range needed {
		if info, ok := funcRespNames[name]; ok && !sqlcModelNames[name] {
			funcFiltered[name] = info
		}
	}
	// Inner types: complex struct fields within @call result types (e.g.
	// ChatMessage inside LoadMessagesResponse). These are in needed (from
	// OpenAPI property $refs) but not in funcRespNames or sqlcModelNames,
	// so they fell through both converter paths before BUG-149 / Phase006.
	funcInnerNames := collectFuncInnerTypeNames(funcRespNames, fs.ProjectFuncSpecs, needed, sqlcModelNames)
	for name, info := range funcInnerNames {
		funcFiltered[name] = info
	}
	if err := emitFuncResponseConverterFiles(fs.OpenAPIDoc, serviceDir, modulePath, funcFiltered, fs.ProjectFuncSpecs, fs.FuncPackageTypes, usedNames, dg); err != nil {
		return fmt.Errorf("func response converters: %w", err)
	}

	// pgtype bridge calls now route through ssac/pkg/pgtypex — no
	// internal/service helper emit needed.

	// Remove a converters.go bundle left behind by older generations so
	// duplicate convert<Name> declarations cannot reach the compiler.
	stale := filepath.Join(serviceDir, "converters.go")
	if _, err := os.Stat(stale); err == nil {
		if err := os.Remove(stale); err != nil {
			return fmt.Errorf("remove stale converters.go: %w", err)
		}
	}

	// Classify the oapi-codegen response wrappers (alias vs embedded struct)
	// from the already-emitted internal/api sources. The map plumbs into each
	// methodGen so error-response literals pick the correct shape for shared
	// `components/responses` references (BUG-106 / Phase012). A nil map (api
	// dir absent / unreadable) triggers the alias fallback in
	// errorResponseLiteral, preserving prior behaviour.
	// In domain mode the oapi-codegen types live in internal/api_<domain>,
	// so classification reads that per-domain directory (Phase007).
	apiDirName := "api"
	if dg.ApiSuffix != "" {
		apiDirName = "api_" + dg.ApiSuffix
	}
	apiDir := filepath.Join(artifactsDir, "backend", "internal", apiDirName)
	respShapes := classifyResponseShapes(apiDir)

	// operationId membership filter (Phase007 linchpin). In domain mode the
	// shared internal/service package is emitted once per domain; without this
	// gate every ServiceFunc's method would be re-emitted for every domain,
	// duplicating declarations in the shared directory. Skipping any sf whose
	// Name is absent from this domain's OpenAPI doc makes each operationId emit
	// exactly once, by its owning domain, with that domain's alias. Single-site
	// (dg.ApiSuffix == "") never filters — behaviour is unchanged.
	var ownedOps map[string]bool
	if dg.ApiSuffix != "" {
		ownedOps = opIDsInDoc(fs.OpenAPIDoc)
	}
	for _, sf := range fs.ServiceFuncs {
		if ownedOps != nil && !ownedOps[sf.Name] {
			continue
		}
		if sf.Subscribe != nil {
			if err := generateSubscribeMethod(sf, fs, serviceDir, modulePath, respShapes); err != nil {
				return fmt.Errorf("subscribe %s: %w", sf.Name, err)
			}
		} else {
			if err := generateHTTPMethod(sf, fs, serviceDir, modulePath, respShapes, dg); err != nil {
				return fmt.Errorf("method %s: %w", sf.Name, err)
			}
		}
	}
	return nil
}
