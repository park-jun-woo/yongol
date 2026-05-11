//ff:func feature=validate type=rule control=sequence topic=stml-design
//ff:what Run — STML↔DESIGN.md 교차 검증 실행 (XVM-01~06, XMV-10~12)
package stml_design

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all STML↔DESIGN.md cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.DesignSpec == nil || len(fs.STMLPages) == 0 {
		return nil
	}

	tokens := extractAllTokens(fs)
	overrides := collectOverrides(fs)

	var diags []diagnostic.Diagnostic
	diags = append(diags, xvm01Color(fs, tokens, overrides)...)
	diags = append(diags, xvm02Rounded(fs, tokens, overrides)...)
	diags = append(diags, xvm03Spacing(fs, tokens, overrides)...)
	diags = append(diags, xvm04Font(fs, tokens, overrides)...)
	diags = append(diags, xvm05Inline(fs, overrides)...)
	diags = append(diags, xvm06ComponentDesignRequired(fs, tokens)...)
	diags = append(diags, xmv10DeadColor(fs, tokens)...)
	diags = append(diags, xmv11DeadTypography(fs, tokens)...)
	diags = append(diags, xmv12DeadComponent(fs, tokens)...)
	return diags
}
