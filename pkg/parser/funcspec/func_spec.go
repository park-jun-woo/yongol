//ff:type feature=funcspec type=model
//ff:what 파싱된 func spec 파일의 데이터 타입
package funcspec

// FuncSpec holds a parsed func spec file.
type FuncSpec struct {
	Package        string   // "auth"
	Name           string   // "hashPassword"
	Description    string   // @description value
	ErrStatus      int      // @error HTTP status code (0 = unspecified)
	Line           int      // 1-based line number of the @func annotation (fallback: func declaration). 0 = unknown.
	RequestFields  []Field  // FuncNameRequest struct fields
	ResponseFields []Field  // FuncNameResponse struct fields
	HasBody        bool     // true if function body is not just "// TODO: implement"
	Imports        []string // import paths (e.g. "database/sql", "net/http")
	// ReturnTypes captures the return types declared on the matching FuncDecl
	// (e.g. ["bool"], ["FooResponse", "error"], ["error"]). Used by @eval
	// validation (S-67) to detect predicate funcs whose sole return is `bool`.
	ReturnTypes []string
}
