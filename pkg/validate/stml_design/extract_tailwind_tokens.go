//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what STML 전체 페이지에서 class 속성의 Tailwind 커스텀 토큰을 추출·분류
package stml_design

import (
	"regexp"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tokenRef records where a potential design token name appears in STML.
type tokenRef struct {
	File  string // STML filename
	Class string // full class attribute string
	Name  string // extracted token name (e.g. "primary", "sm")
}

// pageTokenRefs holds all extracted custom token references from all pages.
type pageTokenRefs struct {
	Colors     []tokenRef // color-prefix tokens (bg-X, text-X, ...)
	Rounded    []tokenRef // rounded-X tokens
	Spacing    []tokenRef // spacing-prefix tokens (p-X, m-X, ...)
	Fonts      []tokenRef // font-X tokens
	Components []tokenRef // data-component references
}

// colorPrefixes lists Tailwind utility prefixes that reference color tokens.
var colorPrefixes = []string{
	"bg-", "text-", "border-", "ring-", "shadow-", "outline-",
	"accent-", "fill-", "stroke-", "divide-",
	"from-", "via-", "to-",
	"decoration-", "placeholder-",
}

// spacingPrefixes lists Tailwind utility prefixes that reference spacing tokens.
var spacingPrefixes = []string{
	"p-", "px-", "py-", "pt-", "pr-", "pb-", "pl-", "ps-", "pe-",
	"m-", "mx-", "my-", "mt-", "mr-", "mb-", "ml-", "ms-", "me-",
	"gap-", "gap-x-", "gap-y-",
	"space-x-", "space-y-",
	"w-", "h-", "min-w-", "min-h-", "max-w-", "max-h-",
	"inset-", "top-", "right-", "bottom-", "left-",
	"basis-", "size-",
}

// tailwindPaletteRe matches standard Tailwind color palette names (e.g. "gray-500", "red-200").
var tailwindPaletteRe = regexp.MustCompile(`^[a-z]+-\d+$`)

// extractAllTokens collects custom token references from all STML pages.
func extractAllTokens(fs *yongol.Fullstack) pageTokenRefs {
	var result pageTokenRefs
	for _, page := range fs.STMLPages {
		extractPageTokens(page, &result)
	}
	return result
}

// extractPageTokens processes a single page.
func extractPageTokens(page stml.PageSpec, out *pageTokenRefs) {
	for _, f := range page.Fetches {
		extractFetchTokens(f, page.FileName, out)
	}
	for _, a := range page.Actions {
		extractActionTokens(a, page.FileName, out)
	}
	for _, c := range page.Children {
		extractChildTokens(c, page.FileName, out)
	}
}

// extractFetchTokens processes a FetchBlock.
func extractFetchTokens(fb stml.FetchBlock, file string, out *pageTokenRefs) {
	classifyTokens(fb.ClassName, file, out)
	for _, b := range fb.Binds {
		classifyTokens(b.ClassName, file, out)
	}
	for _, e := range fb.Eaches {
		extractEachTokens(e, file, out)
	}
	for _, comp := range fb.Components {
		recordComponent(comp, file, out)
		classifyTokens(comp.ClassName, file, out)
	}
	for _, c := range fb.Children {
		extractChildTokens(c, file, out)
	}
	for _, nf := range fb.NestedFetches {
		extractFetchTokens(nf, file, out)
	}
}

// extractActionTokens processes an ActionBlock.
func extractActionTokens(ab stml.ActionBlock, file string, out *pageTokenRefs) {
	classifyTokens(ab.ClassName, file, out)
	for _, f := range ab.Fields {
		classifyTokens(f.ClassName, file, out)
	}
	for _, c := range ab.Children {
		extractChildTokens(c, file, out)
	}
}

// extractEachTokens processes an EachBlock.
func extractEachTokens(eb stml.EachBlock, file string, out *pageTokenRefs) {
	classifyTokens(eb.ClassName, file, out)
	classifyTokens(eb.ItemClassName, file, out)
	for _, b := range eb.Binds {
		classifyTokens(b.ClassName, file, out)
	}
	for _, comp := range eb.Components {
		recordComponent(comp, file, out)
		classifyTokens(comp.ClassName, file, out)
	}
	for _, c := range eb.Children {
		extractChildTokens(c, file, out)
	}
}

// extractChildTokens processes a ChildNode recursively.
func extractChildTokens(cn stml.ChildNode, file string, out *pageTokenRefs) {
	switch cn.Kind {
	case "static":
		if cn.Static != nil {
			extractStaticTokens(*cn.Static, file, out)
		}
	case "fetch":
		if cn.Fetch != nil {
			extractFetchTokens(*cn.Fetch, file, out)
		}
	case "action":
		if cn.Action != nil {
			extractActionTokens(*cn.Action, file, out)
		}
	case "each":
		if cn.Each != nil {
			extractEachTokens(*cn.Each, file, out)
		}
	case "component":
		if cn.Component != nil {
			recordComponent(*cn.Component, file, out)
			classifyTokens(cn.Component.ClassName, file, out)
		}
	case "bind":
		if cn.Bind != nil {
			classifyTokens(cn.Bind.ClassName, file, out)
		}
	}
}

// extractStaticTokens processes a StaticElement recursively.
func extractStaticTokens(se stml.StaticElement, file string, out *pageTokenRefs) {
	classifyTokens(se.ClassName, file, out)
	for _, c := range se.Children {
		extractChildTokens(c, file, out)
	}
}

// recordComponent records a data-component reference.
func recordComponent(comp stml.ComponentRef, file string, out *pageTokenRefs) {
	if comp.Name != "" {
		out.Components = append(out.Components, tokenRef{
			File: file,
			Name: comp.Name,
		})
	}
}

// classifyTokens splits a class string and classifies each part.
// Only non-numeric, non-builtin token names are recorded.
func classifyTokens(class, file string, out *pageTokenRefs) {
	if class == "" {
		return
	}
	for _, part := range strings.Fields(class) {
		classifySingleToken(part, class, file, out)
	}
}

// classifySingleToken classifies a single Tailwind class into a token category.
func classifySingleToken(part, fullClass, file string, out *pageTokenRefs) {
	// Handle responsive/state prefixes (e.g. "sm:bg-primary", "hover:text-accent")
	if idx := strings.LastIndex(part, ":"); idx >= 0 {
		part = part[idx+1:]
	}

	// Negative prefix (e.g. "-mt-xs")
	stripped := part
	if strings.HasPrefix(stripped, "-") {
		stripped = stripped[1:]
	}

	// Check color prefixes
	for _, prefix := range colorPrefixes {
		if strings.HasPrefix(stripped, prefix) {
			name := stripped[len(prefix):]
			if name == "" || isSkippable(name) {
				continue
			}
			out.Colors = append(out.Colors, tokenRef{File: file, Class: fullClass, Name: name})
			return
		}
	}

	// Check rounded prefix
	if strings.HasPrefix(stripped, "rounded-") {
		name := stripped[len("rounded-"):]
		if name == "" || isSkippable(name) {
			return
		}
		out.Rounded = append(out.Rounded, tokenRef{File: file, Class: fullClass, Name: name})
		return
	}

	// Check spacing prefixes
	for _, prefix := range spacingPrefixes {
		if strings.HasPrefix(stripped, prefix) {
			name := stripped[len(prefix):]
			if name == "" || isSkippable(name) {
				continue
			}
			out.Spacing = append(out.Spacing, tokenRef{File: file, Class: fullClass, Name: name})
			return
		}
	}

	// Check font prefix
	if strings.HasPrefix(stripped, "font-") {
		name := stripped[len("font-"):]
		if name == "" || isSkippable(name) {
			return
		}
		out.Fonts = append(out.Fonts, tokenRef{File: file, Class: fullClass, Name: name})
		return
	}
}

// isSkippable returns true for values that cannot be custom DESIGN.md tokens:
// numeric values, Tailwind builtins, palette patterns (gray-500), and arbitrary values.
func isSkippable(s string) bool {
	if s == "" {
		return false
	}
	// Arbitrary value [...]
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		return true
	}
	// Standard Tailwind palette pattern (e.g. "gray-500", "red-200")
	if tailwindPaletteRe.MatchString(s) {
		return true
	}
	// Common Tailwind built-in keywords
	switch s {
	case "full", "none", "auto", "px", "white", "black", "inherit",
		"current", "transparent", "thin", "extralight", "light",
		"normal", "medium", "semibold", "bold", "extrabold",
		"screen", "fit", "min", "max", "prose",
		"sans", "serif", "mono":
		return true
	}
	// Pure numeric (integer, decimal, or fraction: "4", "0.5", "1/2")
	for _, c := range s {
		if (c >= '0' && c <= '9') || c == '.' || c == '/' {
			continue
		}
		return false
	}
	return true
}
