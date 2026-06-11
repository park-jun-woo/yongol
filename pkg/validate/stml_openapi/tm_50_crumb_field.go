//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-50 — sitemap data-crumb-field 위반: 페이지 fetch 부재 / 첫 fetch 2xx 응답에 필드 없음 / 라벨 불가 타입 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm50CrumbField validates every sitemap data-crumb-field declaration
// (plans/stml/sitemap Phase006) against the page it labels: (a) the page
// must declare a data-fetch — the dynamic label is read from fetch data,
// so a fetch-less page could never fill it; (b) the field must be a
// top-level property of the first fetch operation's 2xx response schema
// (the responseFields judgment TM-20 shares — the page emitter reads
// exactly that first fetch's data variable); (c) the field must be a
// string/integer/number scalar — an object or array cannot render as a
// crumb label. An unknown page is TM-39's finding, an unknown operationId
// TM-01's, and an untyped schema ("") passes — only a contradicting type
// is an ERROR. data-crumb-field on a group <li> is rejected by TM-39.
func tm50CrumbField(fs *yongol.Fullstack, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	pages := make(map[string]int, len(fs.STMLPages))
	for i, p := range fs.STMLPages {
		pages[p.Name] = i
	}

	var diags []diagnostic.Diagnostic
	for _, e := range collectSitemapEntries(fs.Sitemap) {
		if e.Node.CrumbField == "" || e.Node.Page == "" {
			continue
		}
		idx, ok := pages[e.Node.Page]
		if !ok {
			continue // TM-39 reports the unknown page
		}
		page := fs.STMLPages[idx]
		if len(page.Fetches) == 0 {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.Sitemap.FileName,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-50] data-crumb-field %q at %s needs a data-fetch on page %q — the dynamic crumb label is read from the fetch response, and the page declares none", e.Node.CrumbField, e.Path, e.Node.Page),
				Advice:  "Add a data-fetch to the page, or remove data-crumb-field to keep the static label",
			})
			continue
		}
		opID := page.Fetches[0].OperationID
		entry, ok := opMap[opID]
		if !ok {
			continue // TM-01 reports the unknown operationId
		}
		field, ok := responseFields(entry.op)[e.Node.CrumbField]
		if !ok {
			diags = append(diags, diagnostic.Diagnostic{
				File:        fs.Sitemap.FileName,
				Phase:       diagnostic.PhaseValidate,
				Level:       diagnostic.LevelError,
				Message:     fmt.Sprintf("[TM-50] data-crumb-field %q at %s is not in the OpenAPI 2xx response schema of %q — the first fetch of page %q", e.Node.CrumbField, e.Path, opID, e.Node.Page),
				Advice:      fmt.Sprintf("Add %q to the 2xx response schema of %q, or name an existing top-level field", e.Node.CrumbField, opID),
				OperationID: opID,
			})
			continue
		}
		switch field.typ {
		case "string", "integer", "number", "":
			// label-renderable scalar (or untyped — no contradiction to report)
		default:
			diags = append(diags, diagnostic.Diagnostic{
				File:        fs.Sitemap.FileName,
				Phase:       diagnostic.PhaseValidate,
				Level:       diagnostic.LevelError,
				Message:     fmt.Sprintf("[TM-50] data-crumb-field %q at %s is %q in the 2xx response schema of %q — only string/integer/number fields can render as a crumb label", e.Node.CrumbField, e.Path, field.typ, opID),
				Advice:      "Name a string or numeric scalar field of the response",
				OperationID: opID,
			})
		}
	}
	return diags
}
