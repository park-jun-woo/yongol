//ff:func feature=manifest type=parser control=sequence
//ff:what handleCreateTable — CREATE TABLE 라인 처리 + pendingArchived/pendingFuncManaged 반영

package ddl

// handleCreateTable extracts the new table name and applies pending archived
// and func-managed state. Both flags are independent and may both be set.
// Returns the new currentTable value.
func handleCreateTable(line string, tables map[string]*Table, pendingArchived bool, pendingFuncManaged bool, file string, lineNum int) string {
	name := extractTableName(line, tables, file, lineNum)
	if name == "" {
		return name
	}
	t := tables[name]
	if t == nil {
		return name
	}
	if pendingArchived {
		t.Archived = true
	}
	if pendingFuncManaged {
		t.FuncManaged = true
	}
	return name
}
