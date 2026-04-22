//ff:type feature=validate type=model topic=ddl-structural
//ff:what sqlFile — 단일 .sql 파일의 경로/이름/내용 보관 구조체
package ddl

// sqlFile holds a single .sql file's path and content.
type sqlFile struct {
	path    string
	name    string
	content string
}
