//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-39 보조 — data-page 없는 항목(그룹 라벨/외부 링크)의 data-crumb-field 거부 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// tm39CrumbFieldMisplaced rejects a data-crumb-field on a page-less
// sitemap entry — a group label or external link (plans/stml/sitemap
// Phase006). The dynamic crumb label is read from the page's data-fetch
// response, so only a page item can carry the attribute; on anything else
// it could never resolve. Part of the TM-39 structural family (the
// placement rules of sitemap attributes); the field itself is TM-50's
// concern.
func tm39CrumbFieldMisplaced(e sitemapEntry, file string) []diagnostic.Diagnostic {
	if e.Node.CrumbField == "" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    file,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[TM-39] data-crumb-field %q at %s sits on an entry without data-page — the dynamic crumb label is read from the page's fetch response, so the attribute belongs to page items only", e.Node.CrumbField, e.Path),
		Advice:  "Move data-crumb-field onto the <li data-page=...> item, or remove it",
	}}
}
