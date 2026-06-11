//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-30 사이트맵 확장 — 동적 메뉴 그룹의 data-label-field 누락/each 항목 스키마 미존재/라벨 불가 타입 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm30SitemapGroupLabelField validates the data-label-field of every
// sitemap dynamic menu group (plans/stml/sitemap Phase007) — the TM-30
// item-schema judgment applied to the menu's label source. One rule, three
// findings: (a) the attribute is **required** on a dynamic group — without
// it the emitted items have no label at all; (b) the field must exist in
// the item schema of the group's data-each array (the same
// ExtractResponseArrayItemFields infrastructure TM-30 resolves item.*
// sources against); (c) the field must be a string/integer/number scalar —
// an object or array cannot render as a menu label (the TM-50 type
// contract). An unresolvable item schema stays silent: TM-01/TM-07 own
// those diagnostics, and an untyped field ("") passes — only a
// contradicting type is an ERROR.
func tm30SitemapGroupLabelField(fs *yongol.Fullstack, raif map[string]map[string]map[string]bool) []diagnostic.Diagnostic {
	rait := oapiparser.ExtractResponseArrayItemTypes(fs.OpenAPIDoc)
	var diags []diagnostic.Diagnostic
	for _, e := range sitemapDynamicGroupEntries(fs.Sitemap) {
		if e.Node.LabelField == "" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.Sitemap.FileName,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-30] dynamic menu group at %s requires data-label-field — the fetched items have no label source without it", e.Path),
				Advice:  "Add data-label-field naming a string field of the data-each item schema (e.g. data-label-field=\"building_name\")",
			})
			continue
		}
		itemFields := raif[e.Node.Fetch][e.Node.Each]
		if itemFields == nil {
			continue // unknown op / non-array each — TM-01/TM-07 own those
		}
		if !itemFields[e.Node.LabelField] {
			diags = append(diags, diagnostic.Diagnostic{
				File:        fs.Sitemap.FileName,
				Phase:       diagnostic.PhaseValidate,
				Level:       diagnostic.LevelError,
				Message:     fmt.Sprintf("[TM-30] data-label-field %q at %s is not in the item schema of data-each %q (operationId %q)", e.Node.LabelField, e.Path, e.Node.Each, e.Node.Fetch),
				Advice:      fmt.Sprintf("Add %q to the array item schema in the OpenAPI response, or name an existing item field", e.Node.LabelField),
				OperationID: e.Node.Fetch,
			})
			continue
		}
		switch typ := rait[e.Node.Fetch][e.Node.Each][e.Node.LabelField]; typ {
		case "", "string", "integer", "number":
			// label-renderable scalar (or untyped — no contradiction to report)
		default:
			diags = append(diags, diagnostic.Diagnostic{
				File:        fs.Sitemap.FileName,
				Phase:       diagnostic.PhaseValidate,
				Level:       diagnostic.LevelError,
				Message:     fmt.Sprintf("[TM-30] data-label-field %q at %s is %q in the item schema of data-each %q (operationId %q) — only string/integer/number fields can render as a menu label", e.Node.LabelField, e.Path, typ, e.Node.Each, e.Node.Fetch),
				Advice:      "Name a string or numeric scalar field of the item schema",
				OperationID: e.Node.Fetch,
			})
		}
	}
	return diags
}
