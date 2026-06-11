//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what Run — STML<->OpenAPI 교차 검증 실행 (TM-01 ~ TM-14, TM-16, TM-17, TM-19 ~ TM-22, TM-24 ~ TM-44, TM-46 ~ TM-50, XMO-10/11/12)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all STML<->OpenAPI cross-validation rules. OpenAPIDoc nil makes
// cross-validation impossible, so that branch early-returns. STML 0 pages is no
// longer an early-return when the frontend is ON: the page loop becomes a no-op
// while the coverage rules (XMO-10/11/12) still run to close the gozhip gap.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil {
		return nil
	}
	if len(fs.STMLPages) == 0 && !frontendEnabled(fs) {
		return nil
	}

	opMap := buildOperationMethodMap(fs.OpenAPIDoc)
	// item schema of each response array field — TM-30 resolves item.<Field>
	// sources against it (same extraction the react emitter keys rows with).
	raif := oapiparser.ExtractResponseArrayItemFields(fs.OpenAPIDoc)

	var diags []diagnostic.Diagnostic
	for _, page := range fs.STMLPages {
		for _, f := range page.Fetches {
			diags = append(diags, validateFetchBlock(f, page.FileName, opMap, fs)...)
		}
		modelFetchMap := buildModelFetchMap(page.Fetches, opMap)
		for _, a := range page.Actions {
			diags = append(diags, validateActionBlock(a, page.FileName, opMap, fs)...)
			diags = append(diags, tm14EnabledWhenRefNotFound(a, page.FileName, modelFetchMap)...)
			diags = append(diags, tm16InvalidatesOpNotFound(a, page.FileName, opMap)...)
			diags = append(diags, tm20CaptureFieldInResponse(a, page.FileName, opMap)...)
			diags = append(diags, tm26RedirectRouteExists(a, page.FileName, fs.STMLPages)...)
			diags = append(diags, tm29ActionOnErrorMissing(a, page.FileName, opMap)...)
			diags = append(diags, tm33RedirectParams(a, page.FileName, opMap, fs.STMLPages)...)
		}
		for _, cond := range collectStateConditions(page.Children) {
			diags = append(diags, tm17GuardSyntax(cond, page.FileName)...)
		}
		diags = append(diags, tm25FlowAttrPlacement(page)...)
		diags = append(diags, tm27RouteParamMissing(page)...)
		diags = append(diags, tm28RouteSegmentUnused(page)...)
		diags = append(diags, tm30ItemSource(page, raif)...)
		diags = append(diags, tm31LinkTargetNotFound(page, fs.STMLPages)...)
		diags = append(diags, tm32LinkParamsUnsatisfied(page, fs.STMLPages, raif)...)
	}

	diags = append(diags, tm10ClassProhibited(fs.STMLPages)...)
	diags = append(diags, tm21CaptureSinkUnused(fs, opMap)...)
	diags = append(diags, tm22ProtectedOpNoTokenSupply(fs, opMap)...)
	diags = append(diags, tm24CookieModeCaptureConflict(fs)...)
	diags = append(diags, xmo10Unconsumed(fs)...)
	diags = append(diags, xmo11NoStml(fs)...)
	diags = append(diags, xmo12NoFrontConsumed(fs)...)
	if fs.Manifest != nil {
		diags = append(diags, tm34IndexTarget(fs)...)
		diags = append(diags, tm35IndexFallback(fs, opMap)...)
	}

	// Sitemap rules (plans/stml/sitemap Phase001/002/003/005/006/007):
	// TM-39~42 validate an existing frontend/sitemap.html, TM-43 runs
	// the reachability BFS over it (listing ≠ reaching), TM-44 rejects a
	// surviving layout data-nav once the sitemap owns the menu truth,
	// TM-46/47 validate the data-roles values and their role-claim wiring,
	// TM-50 the data-crumb-field declarations against each page's first
	// fetch response; TM-49 fires on its absence. Dynamic menu groups
	// (Phase007) get TM-48 (data-entry block / structural completeness)
	// plus the sitemap extensions of TM-01/07/30/31/32 (fetch op, each
	// array field, label field, link target, link params). TM-45 retired
	// in Phase007 — the whole reserved vocabulary graduated.
	if fs.Sitemap != nil {
		diags = append(diags, sitemapDiags(fs, opMap, raif)...)
	}
	diags = append(diags, tm49SitemapAbsent(fs)...)

	if len(fs.Layouts) > 0 || (fs.Manifest != nil && fs.Manifest.Frontend.DefaultLayout != "") {
		diags = append(diags, tm11LayoutNotFound(fs.STMLPages, fs.Layouts)...)
		if fs.Manifest != nil {
			diags = append(diags, tm12DefaultLayoutNotFound(fs.Manifest.Frontend.DefaultLayout, fs.Layouts)...)
		}
		diags = append(diags, tm13UnusedLayout(fs.STMLPages, fs.Layouts, defaultLayoutFromManifest(fs), fs.Sitemap)...)
		diags = append(diags, tm36NavTarget(fs.Layouts, fs.STMLPages)...)
		diags = append(diags, tm37LogoutOp(fs.Layouts, opMap)...)
		diags = append(diags, tm38LogoutMode(fs)...)
	}

	return diags
}
