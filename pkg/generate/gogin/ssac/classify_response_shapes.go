//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what classifyResponseShapes — internal/api AST 스캔으로 <Op><Status>JSONResponse 래퍼 형태(embedded/alias) 분류

package ssac

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// classifyResponseShapes parses every *.go file under apiDir with go/parser
// and classifies each `type <Name> ...JSONResponse` declaration as embedded
// or alias. The returned map keys on the wrapper type name. Errors reading or
// parsing individual files are skipped (best-effort: a missing entry triggers
// the alias fallback at emit time, preserving prior behaviour). A nil map is
// returned when the directory is absent.
func classifyResponseShapes(apiDir string) map[string]respShape {
	entries, err := os.ReadDir(apiDir)
	if err != nil {
		return nil
	}
	shapes := make(map[string]respShape)
	fset := token.NewFileSet()
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
			continue
		}
		path := filepath.Join(apiDir, e.Name())
		file, perr := parser.ParseFile(fset, path, nil, 0)
		if perr != nil {
			continue
		}
		collectShapesFromFile(file, shapes)
	}
	return shapes
}
