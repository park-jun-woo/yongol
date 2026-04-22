//ff:type feature=validate type=util control=selection topic=ssac-structural
//ff:what SSaC 검증 공통 예약어·시퀀스 집합 (goReservedWords, codegenReservedVars, reservedSourceNames, knownSeqTypes)

package ssac

// goReservedWords lists Go language reserved keywords that must not appear
// as variable names or result type names.
var goReservedWords = map[string]bool{
	"break": true, "case": true, "chan": true, "const": true, "continue": true,
	"default": true, "defer": true, "else": true, "fallthrough": true, "for": true,
	"func": true, "go": true, "goto": true, "if": true, "import": true,
	"interface": true, "map": true, "package": true, "range": true, "return": true,
	"select": true, "struct": true, "switch": true, "type": true, "var": true,
}

// codegenReservedVars are variable names used by the generated Go code.
// Using these as SSaC result variable names causes compilation errors.
var codegenReservedVars = map[string]bool{
	"server": true, "ctx": true, "err": true, "tx": true, "qtx": true,
	"db": true, "api": true, "conn": true, "r": true,
}

// reservedSourceNames are sources that must not be reused as result variable names.
var reservedSourceNames = map[string]bool{
	"request": true, "currentUser": true, "config": true, "message": true,
}

// goPrimitiveTypes are Go built-in types that start with lowercase — exempt from S-46.
var goPrimitiveTypes = map[string]bool{
	"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
	"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true,
	"float32": true, "float64": true, "string": true, "bool": true, "byte": true, "rune": true,
	"error": true,
}

// knownSeqTypes lists every sequence directive recognised by the parser.
var knownSeqTypes = map[string]bool{
	"get": true, "post": true, "put": true, "delete": true,
	"empty": true, "exists": true, "state": true, "auth": true,
	"call": true, "publish": true, "response": true, "subscribe": true,
	"verify-password": true,
}
