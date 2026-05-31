//ff:func feature=gen-splitter type=test-helper control=iteration dimension=1
//ff:what findFuncDecl — 파일 선언 목록에서 이름이 일치하는 FuncDecl 조회 헬퍼
package splitter

import "go/ast"

// findFuncDecl returns the FuncDecl named name in file, or nil.
func findFuncDecl(file *ast.File, name string) ast.Decl {
	for _, d := range file.Decls {
		if fd, ok := d.(*ast.FuncDecl); ok && fd.Name.Name == name {
			return d
		}
	}
	return nil
}
