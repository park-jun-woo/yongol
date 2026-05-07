//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what resolveOAPITypeWithFormat — resolve an OpenAPI type+format pair into a precise type string

package ssac_sqlc

var oapiFormatMap = map[string]map[string]string{
	"integer": {
		"int64": "int64",
		"int32": "int32",
		"int16": "int16",
	},
	"number": {
		"float":  "float32",
		"double": "float64",
	},
	"string": {
		"uuid": "uuid",
	},
}

func resolveOAPITypeWithFormat(baseType, format string) string {
	if format == "" {
		return baseType
	}
	fmts, ok := oapiFormatMap[baseType]
	if !ok {
		return baseType
	}
	resolved, ok := fmts[format]
	if !ok {
		return baseType
	}
	return resolved
}
