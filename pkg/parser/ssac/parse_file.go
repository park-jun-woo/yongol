//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what ParseFile — parses a single .ssac file and returns []ServiceFunc
package ssac

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseFile parses a single .ssac file and returns the extracted []ServiceFunc.
func ParseFile(path string) ([]ServiceFunc, []diagnostic.Diagnostic) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "Go parse error: " + err.Error(),
		}}
	}

	imports := collectImports(f)
	structs := collectStructs(f)
	var funcs []ServiceFunc
	var diags []diagnostic.Diagnostic

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		sf, d := parseFuncDecl(fset, fn, f, path, imports, structs)
		diags = append(diags, d...)
		if sf != nil {
			funcs = append(funcs, *sf)
		}
	}
	return funcs, diags
}
