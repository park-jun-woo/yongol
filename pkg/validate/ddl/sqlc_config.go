//ff:type feature=validate type=model topic=ddl-structural
//ff:what sqlcConfig — minimal subset of the sqlc v2 config

package ddl

type sqlcConfig struct {
	SQL []sqlcEntry `yaml:"sql"`
}
