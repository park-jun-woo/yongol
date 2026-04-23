//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what parseCreateTableName — CREATE TABLE 헤더 라인에서 table name 추출

package ddl

import "strings"

// parseCreateTableName extracts the bare table name from a CREATE TABLE
// header line. Handles `IF NOT EXISTS` prefix and quoted identifiers.
// Returns lowercase name, "" if no name can be parsed.
func parseCreateTableName(trim string) string {
	// "CREATE TABLE [IF NOT EXISTS] <name> (..."
	up := strings.ToUpper(trim)
	idx := strings.Index(up, "CREATE TABLE")
	if idx < 0 {
		return ""
	}
	rest := strings.TrimSpace(trim[idx+len("CREATE TABLE"):])
	if strings.HasPrefix(strings.ToUpper(rest), "IF NOT EXISTS") {
		rest = strings.TrimSpace(rest[len("IF NOT EXISTS"):])
	}
	// Stop at "(", whitespace, or ";".
	for i, c := range rest {
		if c == '(' || c == ' ' || c == '\t' || c == ';' {
			return strings.ToLower(strings.Trim(rest[:i], `"`))
		}
	}
	return strings.ToLower(strings.Trim(rest, `"`))
}
