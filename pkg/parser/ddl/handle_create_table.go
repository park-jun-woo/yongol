//ff:func feature=manifest type=parser control=sequence
//ff:what handleCreateTable — CREATE TABLE 라인 처리 + pendingArchived 반영

package ddl

// handleCreateTable extracts the new table name and applies pending archived
// state. Returns the new currentTable value.
func handleCreateTable(line string, tables map[string]*Table, pendingArchived bool, file string, lineNum int) string {
	name := extractTableName(line, tables, file, lineNum)
	if !pendingArchived || name == "" {
		return name
	}
	if t := tables[name]; t != nil {
		t.Archived = true
	}
	return name
}
