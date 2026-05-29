//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestParseEvalRejectsResultCapture — @eval 는 결과 캡처(Type var =) 거절

package ssac

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseEvalRejectsResultCapture(t *testing.T) {
	src := `package service

// @eval bool ok = pkg.IsThing({}) "msg" 402
func Anything(c *gin.Context) {}
`
	dir := t.TempDir()
	path := filepath.Join(dir, "test.go")
	if err := os.WriteFile(path, []byte(src), 0644); err != nil {
		t.Fatal(err)
	}
	_, diags := ParseFile(path)
	if len(diags) == 0 {
		t.Fatal("expected diagnostic for @eval with result capture, got none")
	}
}
