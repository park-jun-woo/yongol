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
//   - internal/service/{str_ptr,ptr_of,deref_*}.go (pointer helpers, 1 file 1 func)
//   - internal/service/convert_<name>.go / convert_<name>_list.go per 200-response schema
//   - internal/service/<func_name>.go per SSaC function (StrictServerInterface method)
//   - Subscribe methods (not part of StrictServerInterface, registered via queue.Subscribe)
//
// Every emitted file holds exactly one top-level func (Phase004: 1-file-1-func
// applied uniformly, replacing the Phase003 POC gating that limited the split
// to ActivateWorkflow's dependency tree).
func Generate(fs *yongol.Fullstack, artifactsDir string) error {
	if len(fs.ServiceFuncs) == 0 {
		return nil
	}
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
	usedNames := make(map[string]bool)
	needed := collectResponseSchemas(fs.OpenAPIDoc)
	if err := emitAllConverterFiles(fs.OpenAPIDoc, serviceDir, modulePath, needed, usedNames); err != nil {
		return fmt.Errorf("converters: %w", err)
	}

	// Remove a converters.go bundle left behind by older generations so
	// duplicate convert<Name> declarations cannot reach the compiler.
	stale := filepath.Join(serviceDir, "converters.go")
	if _, err := os.Stat(stale); err == nil {
		if err := os.Remove(stale); err != nil {
			return fmt.Errorf("remove stale converters.go: %w", err)
		}
	}

	for _, sf := range fs.ServiceFuncs {
		if sf.Subscribe != nil {
			if err := generateSubscribeMethod(sf, fs, serviceDir, modulePath); err != nil {
				return fmt.Errorf("subscribe %s: %w", sf.Name, err)
			}
		} else {
			if err := generateHTTPMethod(sf, fs, serviceDir, modulePath); err != nil {
				return fmt.Errorf("method %s: %w", sf.Name, err)
			}
		}
	}
	return nil
}
