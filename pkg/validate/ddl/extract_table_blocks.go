//ff:func feature=validate type=util control=iteration dimension=1 topic=ddl-structural
//ff:what extractTableBlocks — SQL 파일에서 CREATE TABLE 블록들을 파싱해 반환
package ddl

import (
	"regexp"
	"strings"
)

var createTableRe = regexp.MustCompile(`(?i)^\s*CREATE\s+TABLE\s+(?:IF\s+NOT\s+EXISTS\s+)?([A-Za-z_][A-Za-z0-9_]*)`)

// columnNameRe captures the leading identifier from a column-definition line.
var columnNameRe = regexp.MustCompile(`^([A-Za-z_][A-Za-z0-9_]*)\s+`)

// referencesRe captures the table name from REFERENCES <table>.
var referencesRe = regexp.MustCompile(`(?i)REFERENCES\s+([A-Za-z_][A-Za-z0-9_]*)`)

// extractTableBlocks walks the file lines and returns each CREATE TABLE block
// terminated by a line starting with ")". Inner block-end scanning is
// delegated to findTableBlockEnd to keep the outer iteration shallow.
func extractTableBlocks(f sqlFile) []tableBlock {
	lines := strings.Split(f.content, "\n")
	var blocks []tableBlock
	i := 0
	for i < len(lines) {
		m := createTableRe.FindStringSubmatch(lines[i])
		if m == nil {
			i++
			continue
		}
		end := findTableBlockEnd(lines, i)
		blocks = append(blocks, tableBlock{
			tableName: m[1],
			file:      f,
			startLine: i + 1,
			endLine:   end + 1,
			body:      strings.Join(lines[i:end+1], "\n"),
		})
		i = end + 1
	}
	return blocks
}
