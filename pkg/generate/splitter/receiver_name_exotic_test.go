//ff:func feature=gen-splitter type=test control=sequence
//ff:what methodReceiver / receiverName / funcIdentifier / genDeclIdentifier / declIdentifier / funcFileName / genDeclFileName / fileNameForDecl
package splitter

import (
	"go/ast"
	"testing"
)

func TestReceiverNameExotic(t *testing.T) {
	if got := receiverName(&ast.IndexExpr{X: &ast.Ident{Name: "G"}}); got != "" {
		t.Errorf("generic receiver = %q, want empty", got)
	}
}
