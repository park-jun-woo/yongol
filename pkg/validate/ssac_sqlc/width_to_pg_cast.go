//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what widthToPGCast — OpenAPI int width 를 PostgreSQL 캐스트 토큰으로 매핑

package ssac_sqlc

// widthToPGCast maps an OpenAPI int width to the corresponding PG cast token.
func widthToPGCast(width string) string {
	if width == "int64" {
		return "bigint"
	}
	return "int"
}
