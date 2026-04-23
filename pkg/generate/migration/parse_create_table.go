//ff:func feature=migration type=parser control=sequence
//ff:what parseCreateTable — CREATE TABLE 문장을 파싱해 Schema 에 Table 등록
package migration

import (
	"fmt"
	"regexp"
)

var reCreateTable = regexp.MustCompile(`(?is)^\s*CREATE\s+TABLE\s+(?:IF\s+NOT\s+EXISTS\s+)?([^\s(]+)\s*\((.*)\)\s*$`)

// parseCreateTable parses one CREATE TABLE statement and attaches its
// columns / constraints to the matching Table in s.
func parseCreateTable(s *Schema, stmt string) error {
	m := reCreateTable.FindStringSubmatch(stmt)
	if m == nil {
		return fmt.Errorf("unparseable CREATE TABLE statement: %q", trimForErr(stmt))
	}
	name := canonIdent(m[1])
	body := m[2]
	t := ensureTable(s, name)
	parseTableItems(t, body)
	return nil
}
