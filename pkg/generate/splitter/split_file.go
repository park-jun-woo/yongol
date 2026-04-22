//ff:func feature=gen-splitter type=util control=iteration dimension=1
//ff:what SplitFile — 외부 코드젠 산출 1 파일을 AST 로 읽어 여러 파일로 분할 + filefunc 어노테이션 주입
package splitter

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// SplitFile reads srcPath, groups declarations by their target file, and
// writes one Go file per group into outDir. It is idempotent: existing
// files in outDir with the same name are overwritten. SplitFile does not
// delete the original — cleanOriginal handles that after all splits in a
// directory succeed.
//
// Returns the list of written file names (relative to outDir) so the
// caller can verify the result against preserved_files.go.
func SplitFile(srcPath, outDir string, tool Tool) ([]string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, srcPath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", srcPath, err)
	}
	header := extractHeader(srcPath)
	feature := filepath.Base(filepath.Dir(srcPath))
	funcType := funcTypeFor(tool)
	isModels := strings.HasSuffix(filepath.Base(srcPath), "models.go")
	cmap := ast.NewCommentMap(fset, file, file.Comments)
	units := map[string]*splitUnit{}
	for _, decl := range file.Decls {
		if gd, ok := decl.(*ast.GenDecl); ok && gd.Tok.String() == "import" {
			continue
		}
		name := fileNameForDecl(decl, tool, isModels)
		u, ok := units[name]
		if !ok {
			u = &splitUnit{
				FileName: name,
				PkgName:  file.Name.Name,
				Header:   header,
			}
			units[name] = u
		}
		u.Decls = append(u.Decls, decl)
		u.Docs = append(u.Docs, docOf(decl))
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return nil, err
	}
	var written []string
	for _, u := range units {
		finaliseUnit(u, feature, funcType, cmap, file.Imports)
		if err := writeSplitUnit(outDir, fset, u); err != nil {
			return nil, fmt.Errorf("write %s: %w", u.FileName, err)
		}
		written = append(written, u.FileName)
	}
	return written, nil
}
