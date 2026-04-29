//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what checkPgtypeOverride — sqlc.yaml 의 pgtype override 누락 검사 공유 헬퍼 (Q-12 + per-type Q-NN)

package query

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// checkPgtypeOverride executes one per-type override rule. Skips
// silently when the SpecsDir is empty, no matching DDL column exists,
// or sqlc.yaml is unreadable (D-4 already reports the latter).
func checkPgtypeOverride(fs *yongol.Fullstack, rule pgtypeOverrideRule) []diagnostic.Diagnostic {
	if fs == nil || fs.SpecsDir == "" {
		return nil
	}
	if !ddlHasMatchingColumn(fs.DDLTables, rule.Filter) {
		return nil
	}
	sqlcPath := filepath.Join(fs.SpecsDir, "db", "sqlc.yaml")
	data, err := os.ReadFile(sqlcPath)
	if err != nil {
		return nil
	}
	var cfg sqlcOverridesConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil
	}
	hasNotNull, hasNullable := scanOverridesFor(cfg, rule.DBType, rule.PgPackage, rule.PgType)
	return diagnoseOverrideGaps(rule, hasNotNull, hasNullable)
}
