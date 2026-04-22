//ff:func feature=validate type=rule control=iteration dimension=1 topic=ddl-structural
//ff:what D-5 — sqlc.yaml 의 schema 경로가 DDL 디렉토리를 포함하는지 검증

package ddl

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d05SqlcYamlSchemaPath validates D-5: sqlc.yaml's schema path should cover
// db/*.sql DDL files. Accepted patterns: ".", "./", any path containing "."
// as a component (meaning the db/ directory itself since sqlc.yaml lives in
// db/). WARNING when none of the schema entries look like they point at the
// DDL directory.
func d05SqlcYamlSchemaPath(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.PresenceOf(yongol.KindDDL) == yongol.SSOTAbsent {
		return nil
	}
	schemas, _ := parseSqlcYaml(fs.SpecsDir)
	if schemas == nil {
		return nil
	}
	for _, s := range schemas {
		s = strings.TrimRight(s, "/")
		if s == "." || s == "./" || s == "" {
			return nil
		}
	}
	return []diagnostic.Diagnostic{{
		File:    "db/sqlc.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[D-5] sqlc.yaml schema path does not include current directory (\".\") — db/*.sql DDL files may not be picked up by sqlc",
		Advice:  "schema 를 \".\" 으로 설정하세요",
	}}
}
