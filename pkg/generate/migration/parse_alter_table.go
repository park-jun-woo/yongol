//ff:func feature=migration type=parser control=sequence
//ff:what parseAlterTable — ALTER TABLE ADD [CONSTRAINT ...] FOREIGN KEY 문장 파싱
package migration

import "regexp"

var reAlterAddFK = regexp.MustCompile(`(?is)^\s*ALTER\s+TABLE\s+(?:ONLY\s+)?(\S+)\s+ADD\s+(?:CONSTRAINT\s+(\S+)\s+)?FOREIGN\s+KEY\s*\(([^)]*)\)\s*REFERENCES\s+(\S+?)\s*\(([^)]*)\)(.*)$`)

// parseAlterTable handles `ALTER TABLE ... ADD [CONSTRAINT ...]
// FOREIGN KEY (...)` statements.
func parseAlterTable(s *Schema, stmt string) error {
	m := reAlterAddFK.FindStringSubmatch(stmt)
	if m == nil {
		return nil
	}
	table := canonIdent(m[1])
	name := ""
	if m[2] != "" {
		name = canonIdent(m[2])
	}
	localCols := parseColumnList(m[3])
	refTable := canonIdent(m[4])
	refCols := parseColumnList(m[5])
	fk := &ForeignKey{Columns: localCols, RefTable: refTable, RefColumns: refCols}
	if name == "" {
		name = FKName(table, localCols)
	}
	fk.Name = name
	applyAlterFKOnActions(fk, m[6])
	t := ensureTable(s, table)
	t.ForeignKeys = append(t.ForeignKeys, fk)
	return nil
}
