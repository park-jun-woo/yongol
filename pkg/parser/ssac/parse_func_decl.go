//ff:func feature=ssac-parse type=parser control=sequence
//ff:what parseFuncDecl — extracts a ServiceFunc from an AST function declaration
package ssac

import (
	"go/ast"
	"go/token"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// parseFuncDecl extracts a ServiceFunc from an AST function declaration.
func parseFuncDecl(fset *token.FileSet, fn *ast.FuncDecl, f *ast.File, path string, imports []string, structs []StructInfo) (*ServiceFunc, []diagnostic.Diagnostic) {
	comments := collectFuncComments(f, fn.Pos())

	sequences, err := parseComments(fset, comments)
	if err != nil {
		line := fset.Position(fn.Pos()).Line
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    line,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: filepath.Base(path) + ":" + fn.Name.Name + " — " + err.Error(),
		}}
	}
	if len(sequences) == 0 {
		line := fset.Position(fn.Pos()).Line
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    line,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "[S-74] " + filepath.Base(path) + ":" + fn.Name.Name + " — SSaC function has no annotations. Add at least one sequence (@get, @post, @call, @response, etc.)",
			Advice:  "An SSaC function without annotations is a no-op. Either add sequence annotations or remove the file.",
		}}
	}

	noPagination := hasNoPaginationComment(comments)
	stateNeutral := hasStateNeutralComment(comments)

	sf := ServiceFunc{
		Name:         fn.Name.Name,
		FileName:     filepath.Base(path),
		Line:         fset.Position(fn.Pos()).Line,
		Imports:      imports,
		Structs:      structs,
		Param:        extractParamInfo(fn),
		NoPagination: noPagination,
		StateNeutral: stateNeutral,
	}

	// @subscribe extraction: function metadata, not a sequence
	sf.Sequences = filterSubscribe(&sf, sequences)

	return &sf, nil
}
