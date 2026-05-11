//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-design
//ff:what XVM-05 — STML inline style 하드코딩 색상이 DESIGN.md 토큰 값과 일치 → 토큰 사용 권고 (WARNING)
package stml_design

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// hexColorRe matches 3/4/6/8 digit hex colors in inline styles.
var hexColorRe = regexp.MustCompile(`#(?:[0-9a-fA-F]{3,4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})\b`)

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

// scanInlineStyles parses an HTML file and checks style attributes for hardcoded colors.
func scanInlineStyles(path, filename string, hexToToken map[string]string, ovr overrideSet) []diagnostic.Diagnostic {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	doc, err := html.Parse(f)
	if err != nil {
		return nil
	}

	var diags []diagnostic.Diagnostic
	walkInlineStyles(doc, filename, hexToToken, ovr, &diags)
	return diags
}

// walkInlineStyles recursively walks the DOM checking for inline style hardcoded colors.
func walkInlineStyles(n *html.Node, filename string, hexToToken map[string]string, ovr overrideSet, diags *[]diagnostic.Diagnostic) {
	if n.Type == html.ElementNode {
		style := getNodeAttr(n, "style")
		if style != "" {
			if !isPrecededByOverride(n) {
				checkStyleColors(style, filename, hexToToken, diags)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkInlineStyles(c, filename, hexToToken, ovr, diags)
	}
}

// isPrecededByOverride checks if the given element node has a preceding
// sibling that is an @override comment (skipping whitespace text nodes).
func isPrecededByOverride(n *html.Node) bool {
	for prev := n.PrevSibling; prev != nil; prev = prev.PrevSibling {
		if prev.Type == html.CommentNode && isOverrideComment(prev.Data) {
			return true
		}
		// Skip whitespace text nodes
		if prev.Type == html.TextNode && strings.TrimSpace(prev.Data) == "" {
			continue
		}
		break
	}
	return false
}

// checkStyleColors checks a style attribute value for hardcoded hex colors that
// match DESIGN.md token values.
func checkStyleColors(style, filename string, hexToToken map[string]string, diags *[]diagnostic.Diagnostic) {
	matches := hexColorRe.FindAllString(style, -1)
	seen := make(map[string]bool)
	for _, m := range matches {
		lower := strings.ToLower(m)
		if seen[lower] {
			continue
		}
		seen[lower] = true
		if tokenName, ok := hexToToken[lower]; ok {
			*diags = append(*diags, diagnostic.Diagnostic{
				File:    filename,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XVM-05] inline style contains hardcoded color %s which matches DESIGN.md token %q", m, tokenName),
				Advice:  fmt.Sprintf("Use Tailwind class with token %q (e.g. bg-%s, text-%s) instead of inline style", tokenName, tokenName, tokenName),
			})
		}
	}
}
