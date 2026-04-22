//ff:func feature=validate type=rule control=sequence topic=ddl-structural
//ff:what Run — execute all DDL validation rules (D-*, XDD-*)
package ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all DDL validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, d01SqlcQueryDuplicate(fs)...)
	diags = append(diags, d02NullableColumn(fs)...)
	diags = append(diags, d03SentinelMissing(fs)...)
	diags = append(diags, d04SqlcYamlRequired(fs)...)
	diags = append(diags, d05SqlcYamlSchemaPath(fs)...)
	diags = append(diags, d06SqlcYamlQueriesPath(fs)...)
	diags = append(diags, d07SqlcPositionalParam(fs)...)
	diags = append(diags, xdd61SensitiveNoAnnotation(fs)...)
	return diags
}
