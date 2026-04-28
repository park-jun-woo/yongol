//ff:type feature=validate type=model topic=query-structural
//ff:what sqlcOverridesConfig — sqlc.yaml 의 sql[].gen.go.overrides 만 뽑는 최소 스키마 (Q-12)

package query

// sqlcOverridesConfig is the minimal sqlc.yaml projection required by Q-12:
// only the sql[].gen.go.overrides chain is needed. Declared as a named type
// so each rule file stays annotation-compliant (one type per file under F2)
// and the rule body does not embed an anonymous nested struct.
type sqlcOverridesConfig struct {
	SQL []struct {
		Gen struct {
			Go struct {
				Overrides []sqlcOverrideEntry `yaml:"overrides"`
			} `yaml:"go"`
		} `yaml:"gen"`
	} `yaml:"sql"`
}
