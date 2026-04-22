//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what xqs18LookupDDLType — look up the Go type for a given tableName + colName

package ssac_sqlc

// xqs18LookupDDLType looks up the Go type for a table column.
func xqs18LookupDDLType(ddlColType map[string]map[string]string, tableName, colName string) (string, bool) {
	cols, ok := ddlColType[tableName]
	if !ok {
		return "", false
	}
	goType, ok := cols[colName]
	return goType, ok
}
