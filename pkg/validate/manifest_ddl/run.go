//ff:func feature=validate type=rule control=sequence topic=manifest-infra
//ff:what Run — manifest auth ↔ DDL columns 교차 검증 실행 (XDN-01~03, XDN-05~06)

package manifest_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all manifest ↔ DDL cross-checks anchored on
// backend.auth.user_table and backend.auth.claims. The rules degrade
// gracefully — XDN-02 only fires once XDN-01 has been satisfied, XDN-03
// and XDN-06 only when XDN-02 found a real DDL table, and XDN-05 fires
// independently on claim syntax. XDN-04 is deprecated (superseded by
// XDN-06).
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xdn01UserTableRequired(fs)...)
	diags = append(diags, xdn02UserTableExists(fs)...)
	diags = append(diags, xdn03ClaimColumnExists(fs)...)
	diags = append(diags, xdn05ClaimTypeRequired(fs)...)
	diags = append(diags, xdn06ClaimDDLType(fs)...)
	return diags
}
