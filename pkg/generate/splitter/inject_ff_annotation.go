//ff:func feature=gen-splitter type=util control=selection
//ff:what injectFFAnnotation — 선언 종류에 맞춰 //ff:func 또는 //ff:type + //ff:what 헤더 라인 생성
package splitter

import (
	"fmt"
	"go/ast"
	"strings"
)

// injectFFAnnotation builds the filefunc annotation header lines for a
// decl about to be written into its own split file. It chooses between
// //ff:func (funcs/methods) and //ff:type (struct/alias/interface
// types), picks the control= token for funcs from a pre-computed
// detectControl() result, and derives //ff:what from the first non-empty
// doc-comment line — falling back to the declared identifier when the
// source has no doc.
//
// funcType is the codebook type= value for func decls ("handler" for
// oapi-codegen, "query" for sqlc). Type decls always get type=model.
// The returned slice is a sequence of raw comment text lines (without
// the leading "//"); callers prepend "// " when emitting.
func injectFFAnnotation(decl ast.Decl, feature, funcType, control, dimension, doc string) []string {
	name := declIdentifier(decl)
	what := summariseDoc(doc, name)
	switch d := decl.(type) {
	case *ast.FuncDecl:
		line := fmt.Sprintf("ff:func feature=%s type=%s control=%s", feature, funcType, control)
		if control == "iteration" && dimension != "" {
			line += " dimension=" + dimension
		}
		_ = d
		return []string{line, "ff:what " + what}
	case *ast.GenDecl:
		if d.Tok.String() == "type" {
			return []string{
				fmt.Sprintf("ff:type feature=%s type=model", feature),
				"ff:what " + what,
			}
		}
		// const/var group: emit //ff:func with sequence since filefunc A1
		// treats pure const/var files as exempt — the annotation still
		// keeps codebook validation happy if the file ends up containing
		// a single ValueSpec block at top.
		return []string{
			fmt.Sprintf("ff:func feature=%s type=model control=sequence", feature),
			"ff:what " + what,
		}
	}
	return []string{
		fmt.Sprintf("ff:func feature=%s type=%s control=sequence", feature, funcType),
		"ff:what " + strings.TrimSpace(name),
	}
}
