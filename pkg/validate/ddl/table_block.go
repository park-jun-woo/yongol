//ff:type feature=validate type=model topic=ddl-structural
//ff:what tableBlock — CREATE TABLE 블록의 테이블명·시작/끝 라인·본문 보관 구조체
package ddl

// tableBlock describes one CREATE TABLE block extracted from a SQL file.
type tableBlock struct {
	tableName string
	file      sqlFile
	startLine int    // 1-based line where CREATE TABLE appears
	endLine   int    // 1-based line of the closing );
	body      string // text from CREATE TABLE through the closing );
}
