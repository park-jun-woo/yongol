//ff:type feature=validate type=model topic=ddl-structural
//ff:what sqlcConfig — sqlc v2 config 최소 서브셋

package ddl

type sqlcConfig struct {
	SQL []sqlcEntry `yaml:"sql"`
}
