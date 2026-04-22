//ff:type feature=validate type=model topic=ddl-structural
//ff:what sqlFile — struct holding a single .sql file's path, name, and content
package ddl

// sqlFile holds a single .sql file's path and content.
type sqlFile struct {
	path    string
	name    string
	content string
}
