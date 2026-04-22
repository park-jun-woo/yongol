//ff:type feature=validate type=model topic=ddl-structural
//ff:what sqlcEntry — sqlc v2 sql[] 항목 (schema/queries 경로)

package ddl

type sqlcEntry struct {
	Schema  interface{} `yaml:"schema"`
	Queries interface{} `yaml:"queries"`
}
