//ff:func feature=validate type=rule control=sequence topic=manifest-infra
//ff:what Run — manifest auth ↔ DDL columns 교차 검증 실행 (XDN-01~04)

package manifest_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all manifest ↔ DDL cross-checks anchored on
// backend.auth.user_table and backend.auth.claims. The four rules
// degrade gracefully — XDN-02 only fires once XDN-01 has been satisfied,
// and XDN-03 / XDN-04 only when XDN-02 found a real DDL table — so
// authors get one actionable error at a time rather than a cascade.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xdn01UserTableRequired(fs)...)
	diags = append(diags, xdn02UserTableExists(fs)...)
	diags = append(diags, xdn03ClaimColumnExists(fs)...)
	diags = append(diags, xdn04ClaimColumnType(fs)...)
	return diags
}
