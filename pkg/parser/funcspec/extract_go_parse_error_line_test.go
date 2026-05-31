//ff:func feature=funcspec type=test control=sequence
//ff:what TestFuncspecHelpers — unit tests for the pure funcspec parser helper functions
package funcspec

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestExtractGoParseErrorLine(t *testing.T) {
	// Parsing invalid Go produces a scanner.ErrorList with a line.
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "bad.go", "package p\nfunc {\n", 0)
	if err == nil {
		t.Fatal("expected a parse error")
	}
	if line := extractGoParseErrorLine(err); line <= 0 {
		t.Errorf("expected a positive error line, got %d", line)
	}
	// Non-scanner error → 0.
	if line := extractGoParseErrorLine(errPlain{}); line != 0 {
		t.Errorf("non-scanner error → %d, want 0", line)
	}
}
