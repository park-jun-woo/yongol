//ff:type feature=validate type=model topic=ddl-structural
//ff:what tableBlock — struct holding a CREATE TABLE block's table name, start/end lines, and body text
package ddl

// tableBlock describes one CREATE TABLE block extracted from a SQL file.
type tableBlock struct {
	tableName string
	file      sqlFile
	startLine int    // 1-based line where CREATE TABLE appears
	endLine   int    // 1-based line of the closing );
	body      string // text from CREATE TABLE through the closing );
}
