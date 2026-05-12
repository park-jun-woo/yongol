//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-design
//ff:what XVM-05 — STML inline style 하드코딩 색상이 DESIGN.md 토큰 값과 일치 → 토큰 사용 권고 (WARNING)
package stml_design

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// hexColorRe matches 3/4/6/8 digit hex colors in inline styles.
var hexColorRe = regexp.MustCompile("#(?:[0-9a-fA-F]{3,4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})\\b")

// xvm05Inline detects hardcoded color values in inline style attributes that
// match a DESIGN.md color token value, suggesting token usage instead.
func xvm05Inline(fs *yongol.Fullstack, ovr overrideSet) []diagnostic.Diagnostic {
	if len(fs.DesignSpec.Colors) == 0 {
		return nil
	}

	// Build reverse map: hex value (lowercased) → token name
	hexToToken := make(map[string]string)
	for name, val := range fs.DesignSpec.Colors {
		hexToToken[strings.ToLower(val)] = name
	}

	frontendDir := filepath.Join(fs.SpecsDir, "frontend")
	var diags []diagnostic.Diagnostic

	for _, page := range fs.STMLPages {
		path := filepath.Join(frontendDir, page.FileName)
		pageDiags := scanInlineStyles(path, page.FileName, hexToToken, ovr)
		diags = append(diags, pageDiags...)
	}
	return diags
}
