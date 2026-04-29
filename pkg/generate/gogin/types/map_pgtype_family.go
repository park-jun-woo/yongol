//ff:func feature=gen-gogin type=util control=selection
//ff:what mapPgtypeFamily — head 토큰이 pgtype wrapper family 면 binding 반환

package types

// mapPgtypeFamily routes UUID / NUMERIC / TIMESTAMP / INET / INTERVAL /
// JSONB / BYTEA tokens to their pgtype-backed bindings. Returns ok=false
// when head is not in the pgtype family — caller falls through to the
// native family dispatcher.
func mapPgtypeFamily(head string, notNull bool, defaultLiteral string) (GoTypeBinding, bool) {
	switch head {
	case "UUID":
		return pgtypeUUID(notNull, defaultLiteral), true
	case "NUMERIC", "DECIMAL":
		return pgtypeNumeric(notNull, defaultLiteral), true
	case "TIMESTAMPTZ", "TIMESTAMP", "DATE":
		return pgtypeTimestamp(head, notNull, defaultLiteral), true
	case "INET", "CIDR":
		return pgtypeInet(notNull, defaultLiteral), true
	case "INTERVAL":
		return pgtypeInterval(notNull, defaultLiteral), true
	case "JSONB", "JSON":
		return jsonbBinding(notNull, defaultLiteral), true
	case "BYTEA":
		return byteaBinding(notNull, defaultLiteral), true
	}
	return GoTypeBinding{}, false
}
