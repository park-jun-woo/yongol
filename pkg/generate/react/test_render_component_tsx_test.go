//ff:func feature=gen-react type=test
//ff:what renderComponentTSX — variant/simple 컴포넌트 TSX 렌더링 검증

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

func TestRenderComponentTSX_SimpleCard(t *testing.T) {
	tok := design.ComponentToken{
		Base: "rounded-lg border border-border bg-background shadow-sm",
	}
	src := renderComponentTSX("Card", tok)

	checks := []string{
		"import * as React from 'react'",
		"import { cn } from '@/lib/utils'",
		"React.HTMLAttributes<HTMLDivElement>",
		"cn('rounded-lg border border-border bg-background shadow-sm', className)",
		"<div ref={ref}",
		"Card.displayName = 'Card'",
	}
	for _, want := range checks {
		if !strings.Contains(src, want) {
			t.Errorf("missing expected substring: %q\n\ngot:\n%s", want, src)
		}
	}

	// Should NOT contain variant/size types
	if strings.Contains(src, "type Variant") {
		t.Error("simple component should not have Variant type")
	}
	if strings.Contains(src, "type Size") {
		t.Error("simple component should not have Size type")
	}
}

func TestRenderComponentTSX_InputTag(t *testing.T) {
	tok := design.ComponentToken{
		Base: "flex h-10 w-full rounded-md border",
	}
	src := renderComponentTSX("Input", tok)
	if !strings.Contains(src, "<input ref={ref}") {
		t.Errorf("Input component should use <input> tag\n\ngot:\n%s", src)
	}
	if !strings.Contains(src, "HTMLInputElement") {
		t.Errorf("Input component should reference HTMLInputElement\n\ngot:\n%s", src)
	}
}

func TestRenderComponentTSX_VariantsOnlyNoSizes(t *testing.T) {
	tok := design.ComponentToken{
		Base: "inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium",
		Variants: map[string]string{
			"default":     "bg-primary text-primary-foreground",
			"destructive": "bg-destructive text-destructive-foreground",
		},
		DefaultVariant: "default",
	}
	src := renderComponentTSX("Badge", tok)

	if !strings.Contains(src, "type Variant") {
		t.Error("should have Variant type")
	}
	if strings.Contains(src, "type Size") {
		t.Error("should not have Size type when no sizes defined")
	}
	if !strings.Contains(src, "variant = 'default'") {
		t.Errorf("should default to 'default' variant\n\ngot:\n%s", src)
	}
}

func TestRenderComponentTSX_NoBase(t *testing.T) {
	tok := design.ComponentToken{}
	src := renderComponentTSX("Panel", tok)
	if !strings.Contains(src, "className={className}") {
		t.Errorf("no-base component should just pass className through\n\ngot:\n%s", src)
	}
}
