//ff:func feature=gen-splitter type=util control=sequence
//ff:what funcIdentifier — FuncDecl 식별자(Receiver.Name 또는 Name) 포맷
package splitter

import "go/ast"

// funcIdentifier returns a human-readable identifier for a FuncDecl.
// Methods include the receiver ("T.Method"); plain functions return
// their Name as is. Used as a //ff:what fallback when no doc is present.
func funcIdentifier(d *ast.FuncDecl) string {
	recv := methodReceiver(d)
	if recv == "" {
		return d.Name.Name
	}
	return recv + "." + d.Name.Name
}
