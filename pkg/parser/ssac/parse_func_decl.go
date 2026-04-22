//ff:func feature=ssac-parse type=parser control=sequence
//ff:what AST 함수 선언에서 ServiceFunc를 추출
package ssac

import (
	"go/ast"
	"go/token"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// parseFuncDecl은 AST 함수 선언에서 ServiceFunc를 추출한다.
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
		return nil, nil
	}

	noPagination := hasNoPaginationComment(comments)

	sf := ServiceFunc{
		Name:         fn.Name.Name,
		FileName:     filepath.Base(path),
		Line:         fset.Position(fn.Pos()).Line,
		Imports:      imports,
		Structs:      structs,
		Param:        extractParamInfo(fn),
		NoPagination: noPagination,
	}

	// @subscribe 추출: 시퀀스가 아닌 함수 메타데이터
	sf.Sequences = filterSubscribe(&sf, sequences)

	return &sf, nil
}
