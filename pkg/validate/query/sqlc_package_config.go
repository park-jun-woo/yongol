//ff:type feature=validate type=model topic=query-structural
//ff:what sqlcPackageConfig — sqlc.yaml 의 sql[].gen.go.sql_package 만 뽑는 최소 스키마

package query

// sqlcPackageConfig is the minimal sqlc.yaml projection required by Q-11:
// only the sql[].gen.go.sql_package chain is needed. Declared as a
// named type so the rule file stays annotation-compliant and the
// anonymous struct does not clutter the func body.
type sqlcPackageConfig struct {
	SQL []struct {
		Gen struct {
			Go struct {
				SqlPackage string `yaml:"sql_package"`
			} `yaml:"go"`
		} `yaml:"gen"`
	} `yaml:"sql"`
}
