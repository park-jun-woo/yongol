//ff:func feature=gen-react type=generator control=sequence
//ff:what React + Vite + shadcn-like scaffold 생성 — 모든 파일 한 번의 sequence 로 방출

package react

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// generateFrontendSetup orchestrates the full frontend emit. Sequence is
// significant only for openapi-typescript: api.ts imports from types/api,
// so types/api.d.ts must exist (even as a stub) before Vite's type-check
// runs. All other files are independent.
func generateFrontendSetup(fs *yongol.Fullstack, artifactsDir string) error {
	frontendDir := filepath.Join(artifactsDir, "frontend")
	srcDir := filepath.Join(frontendDir, "src")
	typesDir := filepath.Join(srcDir, "types")
	if err := os.MkdirAll(typesDir, 0o755); err != nil {
		return err
	}

	var theme *manifestTheme
	if fs != nil && fs.Manifest != nil {
		theme = resolveTheme(fs)
	}

	var dspec *design.DesignSpec
	if fs != nil {
		dspec = fs.DesignSpec
	}

	if err := writePackageJSON(frontendDir); err != nil {
		return err
	}
	if err := writeViteConfig(frontendDir); err != nil {
		return err
	}
	if err := writeTSConfig(frontendDir); err != nil {
		return err
	}
	if err := writeIndexHTML(frontendDir); err != nil {
		return err
	}
	if err := writeTailwindConfig(frontendDir, theme, dspec); err != nil {
		return err
	}
	if err := writeMainTSX(srcDir); err != nil {
		return err
	}
	var stmlPages []stml.PageSpec
	var stmlLayouts []stml.LayoutSpec
	var defaultLayout string
	if fs != nil {
		stmlPages = fs.STMLPages
		stmlLayouts = fs.Layouts
		if fs.Manifest != nil {
			defaultLayout = fs.Manifest.Frontend.DefaultLayout
		}
	}
	if err := writeLayoutsTSX(srcDir, stmlLayouts); err != nil {
		return err
	}
	if err := writeAppTSX(srcDir, stmlPages, stmlLayouts, defaultLayout); err != nil {
		return err
	}
	if err := writeLibUtils(srcDir); err != nil {
		return err
	}
	if err := writeComponentsUI(srcDir, dspec); err != nil {
		return err
	}

	// types/api.d.ts must exist before api.ts is written (import resolution).
	// Best-effort: if openapi-typescript fails, a stub is written and the
	// error is propagated so the caller knows.
	// openapi-typescript must run before api.ts emission because api.ts
	// imports `../types/api` at the type level. If the tool is missing, a
	// stub is still written (inside runOpenAPITypescript's error path) so
	// the TypeScript compiler has something to resolve against. The
	// stub-then-error contract means api.ts is always emitted, and the
	// tool error is only returned after the tree is complete.
	typesDest := filepath.Join(typesDir, "api.d.ts")
	specPath := findOpenAPISpec(fs)
	var deferredErr error
	if specPath != "" {
		if err := runOpenAPITypescript(specPath, typesDest); err != nil {
			deferredErr = err
		}
	} else {
		_ = os.WriteFile(typesDest, []byte("export type paths = Record<string, any>\n"), 0o644)
	}

	if err := writeAPIClient(srcDir, fsOpenAPIDoc(fs)); err != nil {
		return err
	}
	return deferredErr
}
