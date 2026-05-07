//ff:func feature=gen-gogin type=util control=selection
//ff:what queryAccessExpr — queryParam 메타를 기반으로 Go 접근 표현식 생성

package ssac

// queryAccessExpr emits Go code that reads a query parameter from
// request.Params, accounting for enum alias types (string(*p)), int32/int64
// format, required vs optional. Enum accessors go through derefEnum to stay
// nil-safe.
func queryAccessExpr(qp queryParam, accessor string) string {
	if qp.IsEnum {
		if qp.IsRequired {
			// required enum → oapi-codegen still emits a named alias type.
			// Cast to string via the underlying value.
			return "string(" + accessor + ")"
		}
		return "derefEnum(" + accessor + ")"
	}
	switch qp.GoType {
	case "integer":
		return primitiveQueryAccess(qp.IsRequired, accessor, "derefInt")
	case "integer32":
		return primitiveQueryAccess(qp.IsRequired, accessor, "derefInt32")
	case "integer64":
		return primitiveQueryAccess(qp.IsRequired, accessor, "derefInt64")
	case "string":
		return primitiveQueryAccess(qp.IsRequired, accessor, "derefStr")
	case "boolean":
		return primitiveQueryAccess(qp.IsRequired, accessor, "derefBool")
	default:
		return primitiveQueryAccess(qp.IsRequired, accessor, "derefStr")
	}
}
