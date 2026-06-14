//ff:type feature=validate type=model topic=openapi-ddl
//ff:what entityRepr — 한 operation 의 2xx 응답 표현 시그니처 (canonical 비교 단위)

package openapi_ddl

// entityRepr captures one operation's 2xx response representation for an entity:
// the operationId, its OpenAPI line (for diagnostics), and the shape signature
// (responseShapeKey) used to detect divergence (XDO-11) and inline drift risk
// (XDO-12).
type entityRepr struct {
	opID     string
	line     int
	shapeKey string
}
