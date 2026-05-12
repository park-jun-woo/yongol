//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what Run — STML<->OpenAPI 교차 검증 실행 (TM-01 ~ TM-13, XMO-10)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all STML<->OpenAPI cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil || len(fs.STMLPages) == 0 {
		return nil
	}

	opMap := buildOperationMethodMap(fs.OpenAPIDoc)

	var diags []diagnostic.Diagnostic
	for _, page := range fs.STMLPages {
		for _, f := range page.Fetches {
			diags = append(diags, validateFetchBlock(f, page.FileName, opMap, fs)...)
		}
		for _, a := range page.Actions {
			diags = append(diags, validateActionBlock(a, page.FileName, opMap, fs)...)
		}
	}

	diags = append(diags, tm10ClassProhibited(fs.STMLPages)...)
	diags = append(diags, xmo10Unconsumed(fs.STMLPages, fs.OpenAPIDoc)...)

	if len(fs.Layouts) > 0 || (fs.Manifest != nil && fs.Manifest.Frontend.DefaultLayout != "") {
		diags = append(diags, tm11LayoutNotFound(fs.STMLPages, fs.Layouts)...)
		if fs.Manifest != nil {
			diags = append(diags, tm12DefaultLayoutNotFound(fs.Manifest.Frontend.DefaultLayout, fs.Layouts)...)
		}
		diags = append(diags, tm13UnusedLayout(fs.STMLPages, fs.Layouts, defaultLayoutFromManifest(fs))...)
	}

	return diags
}
