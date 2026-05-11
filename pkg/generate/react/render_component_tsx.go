//ff:func feature=gen-react type=generator control=sequence
//ff:what renderComponentTSX — ComponentToken → TSX 소스 문자열 렌더링 (variants/sizes 유무에 따라 분기)

package react

import (
	"fmt"
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

// renderComponentTSX produces a complete TSX source for a single component.
// Components with variants and/or sizes get a variant-aware forwardRef
// component. Simple components (base only) get a minimal forwardRef wrapper.
func renderComponentTSX(name string, tok design.ComponentToken) string {
	hasVariants := len(tok.Variants) > 0
	hasSizes := len(tok.Sizes) > 0

	if hasVariants || hasSizes {
		return renderVariantComponent(name, tok)
	}
	return renderSimpleComponent(name, tok)
}

// renderVariantComponent produces a forwardRef component with variant/size support.
func renderVariantComponent(name string, tok design.ComponentToken) string {
	var b strings.Builder

	b.WriteString("import * as React from 'react'\n")
	b.WriteString("import { cn } from '@/lib/utils'\n\n")

	htmlTag := inferHTMLTag(name)
	htmlElement := inferHTMLElement(htmlTag)

	variantKeys := sortedKeys(tok.Variants)
	sizeKeys := sortedKeys(tok.Sizes)

	// Variant type
	if len(variantKeys) > 0 {
		b.WriteString("type Variant = ")
		for i, k := range variantKeys {
			if i > 0 {
				b.WriteString(" | ")
			}
			b.WriteString(fmt.Sprintf("'%s'", k))
		}
		b.WriteString("\n")
	}

	// Size type
	if len(sizeKeys) > 0 {
		b.WriteString("type Size = ")
		for i, k := range sizeKeys {
			if i > 0 {
				b.WriteString(" | ")
			}
			b.WriteString(fmt.Sprintf("'%s'", k))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")

	// Props type
	b.WriteString(fmt.Sprintf("export type %sProps = React.%s & {\n", name, htmlAttrsType(htmlTag)))
	if len(variantKeys) > 0 {
		b.WriteString("  variant?: Variant\n")
	}
	if len(sizeKeys) > 0 {
		b.WriteString("  size?: Size\n")
	}
	b.WriteString("}\n\n")

	// Variants record
	if len(variantKeys) > 0 {
		b.WriteString("const variants: Record<Variant, string> = {\n")
		for _, k := range variantKeys {
			b.WriteString(fmt.Sprintf("  %s: '%s',\n", k, tok.Variants[k]))
		}
		b.WriteString("}\n\n")
	}

	// Sizes record
	if len(sizeKeys) > 0 {
		b.WriteString("const sizes: Record<Size, string> = {\n")
		for _, k := range sizeKeys {
			b.WriteString(fmt.Sprintf("  %s: '%s',\n", k, tok.Sizes[k]))
		}
		b.WriteString("}\n\n")
	}

	// Determine defaults
	defaultVariant := tok.DefaultVariant
	if defaultVariant == "" && len(variantKeys) > 0 {
		defaultVariant = variantKeys[0]
	}
	defaultSize := tok.DefaultSize
	if defaultSize == "" && len(sizeKeys) > 0 {
		defaultSize = sizeKeys[0]
	}

	// Component
	b.WriteString(fmt.Sprintf("export const %s = React.forwardRef<%s, %sProps>(\n", name, htmlElement, name))

	// Destructure params
	var destructParams []string
	if len(variantKeys) > 0 {
		destructParams = append(destructParams, fmt.Sprintf("variant = '%s'", defaultVariant))
	}
	if len(sizeKeys) > 0 {
		destructParams = append(destructParams, fmt.Sprintf("size = '%s'", defaultSize))
	}
	destructParams = append(destructParams, "className", "...props")
	b.WriteString(fmt.Sprintf("  ({ %s }, ref) => (\n", strings.Join(destructParams, ", ")))

	// cn() arguments
	var cnArgs []string
	if tok.Base != "" {
		cnArgs = append(cnArgs, fmt.Sprintf("'%s'", tok.Base))
	}
	if len(variantKeys) > 0 {
		cnArgs = append(cnArgs, "variants[variant]")
	}
	if len(sizeKeys) > 0 {
		cnArgs = append(cnArgs, "sizes[size]")
	}
	cnArgs = append(cnArgs, "className")

	b.WriteString(fmt.Sprintf("    <%s\n", htmlTag))
	b.WriteString("      ref={ref}\n")
	b.WriteString(fmt.Sprintf("      className={cn(%s)}\n", strings.Join(cnArgs, ", ")))
	b.WriteString("      {...props}\n")
	b.WriteString("    />\n")
	b.WriteString("  ),\n")
	b.WriteString(")\n")
	b.WriteString(fmt.Sprintf("%s.displayName = '%s'\n", name, name))

	return b.String()
}

// renderSimpleComponent produces a base-only forwardRef component.
func renderSimpleComponent(name string, tok design.ComponentToken) string {
	var b strings.Builder

	b.WriteString("import * as React from 'react'\n")
	b.WriteString("import { cn } from '@/lib/utils'\n\n")

	htmlTag := inferHTMLTag(name)
	htmlElement := inferHTMLElement(htmlTag)

	b.WriteString(fmt.Sprintf("export const %s = React.forwardRef<%s, React.%s>(\n", name, htmlElement, htmlAttrsType(htmlTag)))
	if tok.Base != "" {
		b.WriteString(fmt.Sprintf("  ({ className, ...props }, ref) => (\n"))
		b.WriteString(fmt.Sprintf("    <%s ref={ref} className={cn('%s', className)} {...props} />\n", htmlTag, tok.Base))
	} else {
		b.WriteString(fmt.Sprintf("  ({ className, ...props }, ref) => (\n"))
		b.WriteString(fmt.Sprintf("    <%s ref={ref} className={className} {...props} />\n", htmlTag))
	}
	b.WriteString("  ),\n")
	b.WriteString(")\n")
	b.WriteString(fmt.Sprintf("%s.displayName = '%s'\n", name, name))

	return b.String()
}

// inferHTMLTag maps component names to HTML tags.
func inferHTMLTag(name string) string {
	lower := strings.ToLower(name)
	switch {
	case lower == "button":
		return "button"
	case lower == "input":
		return "input"
	case lower == "select":
		return "select"
	case lower == "textarea":
		return "textarea"
	case lower == "form":
		return "form"
	case lower == "table":
		return "table"
	case lower == "label":
		return "label"
	case lower == "a" || lower == "link":
		return "a"
	default:
		return "div"
	}
}

// inferHTMLElement maps HTML tags to their TypeScript element types.
func inferHTMLElement(tag string) string {
	switch tag {
	case "button":
		return "HTMLButtonElement"
	case "input":
		return "HTMLInputElement"
	case "select":
		return "HTMLSelectElement"
	case "textarea":
		return "HTMLTextAreaElement"
	case "form":
		return "HTMLFormElement"
	case "table":
		return "HTMLTableElement"
	case "label":
		return "HTMLLabelElement"
	case "a":
		return "HTMLAnchorElement"
	default:
		return "HTMLDivElement"
	}
}

// htmlAttrsType returns the React HTML attributes type for a given tag.
func htmlAttrsType(tag string) string {
	switch tag {
	case "button":
		return "ButtonHTMLAttributes<HTMLButtonElement>"
	case "input":
		return "InputHTMLAttributes<HTMLInputElement>"
	case "select":
		return "SelectHTMLAttributes<HTMLSelectElement>"
	case "textarea":
		return "TextareaHTMLAttributes<HTMLTextAreaElement>"
	case "form":
		return "FormHTMLAttributes<HTMLFormElement>"
	case "table":
		return "TableHTMLAttributes<HTMLTableElement>"
	case "label":
		return "LabelHTMLAttributes<HTMLLabelElement>"
	case "a":
		return "AnchorHTMLAttributes<HTMLAnchorElement>"
	default:
		return "HTMLAttributes<HTMLDivElement>"
	}
}

// sortedKeys returns map keys in sorted order for deterministic output.
func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
