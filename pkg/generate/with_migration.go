//ff:func feature=generate type=util control=sequence
//ff:what WithMigration — DDL 마이그레이션 단계 설정을 GenerateOption 으로 래핑
package generate

// WithMigration attaches the DDL migration step configuration.
func WithMigration(h MigrationHook) GenerateOption {
	return func(c *generateConfig) { c.migration = h }
}
