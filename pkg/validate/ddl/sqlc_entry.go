//ff:type feature=validate type=model topic=ddl-structural
//ff:what sqlcEntry — sqlc v2 sql[] entry (schema/queries paths)

package ddl

type sqlcEntry struct {
	Schema  interface{} `yaml:"schema"`
	Queries interface{} `yaml:"queries"`
}
