//ff:func feature=gen-gogin type=util control=selection
//ff:what integerGoType — OpenAPI integer format (int32/int64) to internal GoType string

package ssac

func integerGoType(format string) string {
	switch format {
	case "int64":
		return "integer64"
	case "int32":
		return "integer32"
	default:
		return "integer"
	}
}
