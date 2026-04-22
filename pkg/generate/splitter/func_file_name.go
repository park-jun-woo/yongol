//ff:func feature=gen-splitter type=util control=sequence
//ff:what funcFileName — FuncDecl 의 파일명 계산 (method 는 receiver_method 접두)
package splitter

import "go/ast"

// funcFileName renders the split file name for a top-level func or
// method. For methods the result is receiver_method + suffix; plain
// funcs are just method + suffix. Exotic receivers (generic
// instantiations) fall back to the plain form.
func funcFileName(d *ast.FuncDecl, suffix string) string {
	name := snake(d.Name.Name)
	recv := methodReceiver(d)
	if recv == "" {
		return name + suffix
	}
	return snake(recv) + "_" + name + suffix
}
