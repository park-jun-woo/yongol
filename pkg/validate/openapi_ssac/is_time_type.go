//ff:func feature=validate type=util control=sequence topic=openapi-ssac
//ff:what isTimeType — report whether a Go type name is time.Time / *time.Time

package openapi_ssac

// isTimeType reports whether t is the Go time type produced by GoTypeOf for
// DDL TIMESTAMPTZ/TIMESTAMP/DATE columns. Only bare time.Time / *time.Time are
// matched: XOS-67's actual comes from GoTypeOf, which collapses every timestamp
// family to bare time.Time, so pgtype.* wrappers never appear here.
func isTimeType(t string) bool {
	return t == "time.Time" || t == "*time.Time"
}
