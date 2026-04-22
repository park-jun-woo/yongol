//ff:func feature=policy type=parser control=sequence
//ff:what strictParse — OPA AST 기반 Rego 문법 엄격 검증 (R-1 parse ERROR)

package rego

import (
	"github.com/open-policy-agent/opa/ast"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// strictParse runs the OPA AST parser on raw .rego content and surfaces any
// parse error as an R-1 parse ERROR. Uses RegoV1 options to match yongol's
// required syntax. Error enumeration is delegated to collectOpaErrors.
func strictParse(path, content string) []diagnostic.Diagnostic {
	_, err := ast.ParseModuleWithOpts(path, content, ast.ParserOptions{RegoVersion: ast.RegoV1})
	if err == nil {
		return nil
	}
	if errs, ok := err.(ast.Errors); ok {
		return collectOpaErrors(path, errs)
	}
	return []diagnostic.Diagnostic{{
		File:    path,
		Phase:   diagnostic.PhaseParse,
		Level:   diagnostic.LevelError,
		Message: "R-1: Rego parse error: " + err.Error(),
	}}
}
