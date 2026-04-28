//ff:type feature=validate type=model topic=query-structural
//ff:what sqlcOverrideEntry — sqlc.yaml gen.go.overrides 단일 entry (db_type / nullable / go_type)

package query

// sqlcOverrideEntry mirrors a single sqlc gen.go.overrides element. Only the
// fields Q-12 inspects are declared. `nullable` defaults to false when the
// key is absent (sqlc convention); the omitted key therefore matches a
// "NOT NULL" entry.
type sqlcOverrideEntry struct {
	DBType   string `yaml:"db_type"`
	Nullable bool   `yaml:"nullable"`
	GoType   struct {
		Import  string `yaml:"import"`
		Package string `yaml:"package"`
		Type    string `yaml:"type"`
	} `yaml:"go_type"`
}
