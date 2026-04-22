//ff:func feature=validate type=util control=sequence topic=ddl-structural
//ff:what parseSqlcYaml — db/sqlc.yaml 를 읽어 schema/queries 경로 추출

package ddl

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// parseSqlcYaml reads db/sqlc.yaml and returns the first sql entry's schema
// and queries paths as string slices. Returns nil slices when the file does
// not exist or is unparseable (D-4 already reports the missing-file case).
func parseSqlcYaml(specsDir string) (schemas []string, queries []string) {
	sqlcPath := filepath.Join(specsDir, "db", "sqlc.yaml")
	data, err := os.ReadFile(sqlcPath)
	if err != nil {
		return nil, nil
	}
	var cfg sqlcConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil || len(cfg.SQL) == 0 {
		return nil, nil
	}
	entry := cfg.SQL[0]
	schemas = toStringSlice(entry.Schema)
	queries = toStringSlice(entry.Queries)
	return schemas, queries
}
