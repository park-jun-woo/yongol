//ff:func feature=contract type=util control=iteration dimension=1
//ff:what ComputeBodyHash — 파일의 첫 번째 func(init 제외) 소스의 sha256 앞 8 hex. filefunc A7 호환.

package contract

import (
	"crypto/sha256"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// ComputeBodyHash returns the hex-encoded first 4 bytes (8 hex chars)
// of the SHA-256 digest of the primary func declaration in src. The
// input is first normalised (CRLF→LF, BOM stripped, trailing newline
// ensured) so hashes are stable across editor / OS differences.
//
// The hashed region is the source span of the first non-`init`
// FuncDecl, matching filefunc's `parse.CalcBodyHash` exactly so a
// hash value embedded by yongol passes filefunc A7 without special
// casing.
//
// When src contains no func declaration (type-only files, package
// statement only, etc.) the empty string is returned. Callers treat
// that as "no hash to write"; filefunc A7 also short-circuits when no
// func is present, so omitting the annotation is safe.
//
// Parse errors are also reported as an empty hash — yongol generates
// Go code that is guaranteed to compile, so a parse failure there
// signals a bug upstream that should surface via go build rather than
// a corrupted hash.
func ComputeBodyHash(src []byte) string {
	normalised := NormalizeBody(src)
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", normalised, 0)
	if err != nil {
		return ""
	}
	for _, decl := range f.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok || fd.Body == nil || fd.Name.Name == "init" {
			continue
		}
		start := fset.Position(fd.Pos()).Offset
		end := fset.Position(fd.End()).Offset
		if start < 0 || end > len(normalised) || start >= end {
			return ""
		}
		sum := sha256.Sum256(normalised[start:end])
		return fmt.Sprintf("%x", sum[:4])
	}
	return ""
}
