//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what renderComponentTSX variant Button TSX 렌더링 검증

package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

func TestRenderComponentTSX_VariantButton(t *testing.T) {
	tok := design.ComponentToken{
		Base: "inline-flex items-center justify-center rounded-md font-medium",
		Variants: map[string]string{
			"primary":   "bg-primary text-primary-foreground hover:bg-primary/90",
			"secondary": "bg-secondary text-secondary-foreground hover:bg-secondary/80",
		},
		Sizes: map[string]string{
			"sm": "h-8 px-3 text-sm",
			"md": "h-10 px-4 text-sm",
		},
		DefaultVariant: "primary",
		DefaultSize:    "md",
	}
	src := renderComponentTSX("Button", tok)

	checks := []string{
		"import * as React from 'react'",
		"import { cn } from '@/lib/utils'",
		"type Variant = 'primary' | 'secondary'",
		"type Size = 'md' | 'sm'",
		"export type ButtonProps = React.ButtonHTMLAttributes<HTMLButtonElement>",
		"variant?: Variant",
		"size?: Size",
		"const variants: Record<Variant, string>",
		"primary: 'bg-primary text-primary-foreground hover:bg-primary/90'",
		"const sizes: Record<Size, string>",
		"sm: 'h-8 px-3 text-sm'",
		"variant = 'primary'",
		"size = 'md'",
		"cn('inline-flex items-center justify-center rounded-md font-medium'",
		"variants[variant]",
		"sizes[size]",
		"<button",
		"Button.displayName = 'Button'",
	}
	for _, want := range checks {
		if !strings.Contains(src, want) {
			t.Errorf("missing expected substring: %q\n\ngot:\n%s", want, src)
		}
	}
}
