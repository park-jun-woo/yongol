//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what writeEnvHelperFiles — top-level 헬퍼 func 선언을 cmd/<name>.go 로 개별 emit

package boot

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ettle/strcase"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/fffile"
)

// writeEnvHelperFiles emits each top-level helper declaration in funcs to
// its own cmd/<snake_case>.go file. Each file carries imports derived from
// the caller-supplied imports slice filtered to entries the helper body
// actually references, plus a complete //ff:func + //ff:what annotation
// so filefunc F1 / A1 / A3 pass on cmd/.
//
// Duplicate file names (theoretically impossible once env helper names are
// unique across blocks) are resolved via fffile.EnsureUnique.
func writeEnvHelperFiles(artifactsDir string, imports, funcs []string) error {
	outDir := filepath.Join(artifactsDir, "backend", "cmd")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}
	used := make(map[string]bool)
	for _, fn := range funcs {
		fn = strings.TrimSpace(fn)
		if fn == "" {
			continue
		}
		name := parseFuncName(fn)
		if name == "" {
			continue
		}
		file := fffile.EnsureUnique(strcase.ToSnake(name)+".go", used)
		src := renderEnvHelperFile(name, fn, imports)
		if err := os.WriteFile(filepath.Join(outDir, file), []byte(src), 0o644); err != nil {
			return err
		}
	}
	return nil
}
