//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-49 — the Method part of Model.Method must exist in the symbol table

package ssac

import (
	"fmt"

	"github.com/jinzhu/inflection"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s49SymbolTableMethod validates S-49: Method portion of Model.Method must
// appear under Ground.Lookup["SymbolTable.method.<Model>"].
func s49SymbolTableMethod(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if !crudType(seq) {
				continue
			}
			model := extractModel(seq)
			method := extractMethod(seq)
			if model == "" || method == "" {
				continue
			}
			methods, ok := g.Lookup["SymbolTable.method."+model]
			if !ok {
				continue
			}
			if methods[method] {
				continue
			}
			expected := expectedQueryFile(model)
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-49] method %s not found on model %s", method, model),
				Advice: fmt.Sprintf(
					"sqlc queries for %s must be in %s (model name is derived from query filename: %s → %s). "+
						"If the query is in a different file, move it to %s.",
					model, expected, expected, model, expected),
			})
		}
	}
	return diags
}

// expectedQueryFile derives the expected sqlc query file path from a model
// name. This is the reverse of modelFromFilename: PascalCase → snake_case →
// plural → append .sql. Example: "RefreshToken" → "db/queries/refresh_tokens.sql".
//
//ff:func feature=validate type=util control=sequence dimension=1 topic=ssac-structural
//ff:what expectedQueryFile — 모델명에서 기대 쿼리 파일 경로 역산 (RefreshToken → db/queries/refresh_tokens.sql)
func expectedQueryFile(model string) string {
	snake := caseconv.PascalToSnake(model)
	plural := inflection.Plural(snake)
	return "db/queries/" + plural + ".sql"
}
