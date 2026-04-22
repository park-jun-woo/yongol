//ff:func feature=gen-gogin type=command control=sequence
//ff:what Generate — Block Builder 로 artifacts/backend/cmd/main.go + env helper files 생성

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// Generate produces artifacts/<project>/backend/cmd/main.go by collecting
// active blocks based on Fullstack state, deduplicating imports, and
// assembling body lines. Top-level helpers (envInt, envDuration, …) are
// emitted as sibling 1-file-1-func files in cmd/ via writeEnvHelperFiles
// so filefunc F1 passes on the boot surface.
func Generate(fs *yongol.Fullstack, artifactsDir string) error {
	modulePath := ""
	if fs.Manifest != nil {
		modulePath = fs.Manifest.Backend.Module
	}
	blocks := collectActiveBlocks(fs, modulePath)
	imports := deduplicateImports(blocks)
	body := assembleBody(blocks)
	funcs := collectFuncs(blocks)
	if err := writeMainGo(artifactsDir, imports, body); err != nil {
		return err
	}
	if err := writeEnvHelperFiles(artifactsDir, imports, funcs); err != nil {
		return err
	}
	// Phase004 — requestIDHandler (slog wrapper for ctx.request_id).
	return writeRequestIDHandlerFile(artifactsDir, modulePath)
}
