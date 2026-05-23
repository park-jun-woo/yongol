//ff:func feature=agent type=helper control=selection
//ff:what classifyFile — 파일 경로로 SSOT 레이어 판별

package agent

import (
	"path/filepath"
	"strings"
)

// classifyFile returns the SSOT layer for a file path (relative to specs-dir).
func classifyFile(relPath string) layer {
	dir := filepath.Dir(relPath)
	base := filepath.Base(relPath)
	ext := filepath.Ext(base)

	switch {
	case strings.HasPrefix(dir, "service") && ext == ".ssac":
		return layerSSaC
	case dir == "db" && ext == ".sql" && !strings.HasPrefix(dir, "db/queries"):
		return layerDDL
	case strings.HasPrefix(dir, "db/queries") && ext == ".sql":
		return layerSQLcQuery
	case relPath == "api/openapi.yaml" || relPath == filepath.Join("api", "openapi.yaml"):
		return layerOpenAPI
	case base == "manifest.yaml" && (dir == "." || dir == ""):
		return layerManifest
	case strings.HasPrefix(dir, "policy") && ext == ".rego":
		return layerRego
	case strings.HasPrefix(dir, "states") && ext == ".md":
		return layerStateDiagram
	case strings.HasPrefix(dir, "func") && ext == ".go":
		return layerFuncSpec
	case strings.HasPrefix(dir, "tests") && ext == ".hurl":
		return layerHurl
	default:
		return layerUnknown
	}
}
