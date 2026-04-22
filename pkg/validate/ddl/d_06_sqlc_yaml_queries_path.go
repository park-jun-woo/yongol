//ff:func feature=validate type=rule control=iteration dimension=1 topic=ddl-structural
//ff:what D-6 — verify sqlc.yaml queries path includes queries/

package ddl

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d06SqlcYamlQueriesPath validates D-6: sqlc.yaml's queries path should cover
// db/queries/*.sql files. Accepted patterns: "queries", "queries/", "./queries",
// "./queries/". WARNING when no queries entry contains "queries".
func d06SqlcYamlQueriesPath(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.PresenceOf(yongol.KindDDL) == yongol.SSOTAbsent {
		return nil
	}
	_, queries := parseSqlcYaml(fs.SpecsDir)
	if queries == nil {
		return nil
	}
	for _, q := range queries {
		if strings.Contains(q, "queries") {
			return nil
		}
	}
	return []diagnostic.Diagnostic{{
		File:    "db/sqlc.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[D-6] sqlc.yaml queries path does not include \"queries/\" — db/queries/*.sql files may not be picked up by sqlc",
		Advice:  "Set queries to \"queries/\"",
	}}
}
