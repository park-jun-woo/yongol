//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what renderComponentTSX simple Card TSX 렌더링 검증

package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

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

	if strings.Contains(src, "type Variant") {
		t.Error("simple component should not have Variant type")
	}
	if strings.Contains(src, "type Size") {
		t.Error("simple component should not have Size type")
	}
}
