//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what Run — STML↔OpenAPI 교차 검증 실행 (TM-01 ~ TM-13, XMO-10)
package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all STML↔OpenAPI cross-validation rules.
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

	// TM-10: class attribute directly on STML elements is prohibited
	diags = append(diags, tm10ClassProhibited(fs.STMLPages)...)

	// XMO-10: OpenAPI operationId not consumed by any STML page
	diags = append(diags, xmo10Unconsumed(fs.STMLPages, fs.OpenAPIDoc)...)

	// TM-11 ~ TM-13: layout cross-validation (skip if no layouts defined)
	if len(fs.Layouts) > 0 || (fs.Manifest != nil && fs.Manifest.Frontend.DefaultLayout != "") {
		diags = append(diags, tm11LayoutNotFound(fs.STMLPages, fs.Layouts)...)
		if fs.Manifest != nil {
			diags = append(diags, tm12DefaultLayoutNotFound(fs.Manifest.Frontend.DefaultLayout, fs.Layouts)...)
		}
		diags = append(diags, tm13UnusedLayout(fs.STMLPages, fs.Layouts, defaultLayoutFromManifest(fs))...)
	}

	return diags
}

// defaultLayoutFromManifest extracts the defaultLayout value from the manifest,
// returning empty string if the manifest is nil.
func defaultLayoutFromManifest(fs *yongol.Fullstack) string {
	if fs.Manifest == nil {
		return ""
	}
	return fs.Manifest.Frontend.DefaultLayout
}
