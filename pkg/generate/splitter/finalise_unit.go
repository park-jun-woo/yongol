//ff:func feature=gen-splitter type=util control=sequence
//ff:what finaliseUnit — splitUnit 에 imports 선별 + filefunc 어노테이션 라인 채워 넣기
package splitter

import "go/ast"

// finaliseUnit fills the Imports and Annotations slots of u based on its
// Decls. Annotation selection uses the primary (first) decl — when
// multiple decls share a file (a type plus its methods, or grouped
// const/var blocks) the first decl is the representative one for
// //ff:what. funcType picks the codebook type= label (handler vs query).
func finaliseUnit(u *splitUnit, feature, funcType string, cmap ast.CommentMap, allImports []*ast.ImportSpec) {
	u.Imports = collectImportsForDecl(u.Decls, allImports)
	_ = cmap
	primary := u.Decls[0]
	doc := ""
	if len(u.Docs) > 0 {
		doc = u.Docs[0]
	}
	control, dimension := controlFor(primary)
	u.Annotations = injectFFAnnotation(primary, feature, funcType, control, dimension, doc)
}
