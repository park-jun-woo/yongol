//ff:func feature=generate type=generator control=sequence
//ff:what runSTMLCodegen — STML 페이지 스펙에서 React TSX 페이지 파일 생성

package generate

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	stmlgen "github.com/park-jun-woo/yongol/pkg/generate/react/stml"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runSTMLCodegen generates page TSX files from fs.STMLPages into
// <frontendDir>/src/pages. Skips silently if no STML pages were parsed.
//
// srcFrontendDir is the directory stmlgen resolves per-page
// <dir>/<page>.custom.ts against (collect_imports.go). Single-site passes
// fs.SpecsDir (unchanged), domain mode passes filepath.Join(fs.SpecsDir,
// cfg.Frontend) so the custom.ts sidecars resolve per domain — DomainView
// keeps SpecsDir shared, so the source dir must be threaded explicitly rather
// than synthesized from SpecsDir+convention (Decision N).
func runSTMLCodegen(fs *yongol.Fullstack, srcFrontendDir, frontendDir string) error {
	if fs == nil || len(fs.STMLPages) == 0 {
		return nil
	}
	pagesDir := filepath.Join(frontendDir, "src", "pages")
	// Prepared mode, not raw ResolvedMode(): keeps the capture-commit gate
	// aligned with the backend emitters and the react auth gates (Phase004
	// — including the BUG-014 jwt-without-mode → bearer rule).
	bearerAuth := prepared.AuthFor(fs).Mode == "bearer"
	constraints := fillDefaultRequestConstraints(fs.STMLPages, fs.OpenAPIDoc, fs.RequestConstraints)
	noBodyOps := oapiparser.ExtractNoBodyOps(fs.OpenAPIDoc)
	// STML field-less actions are also body-less (no form data to send).
	mergeFieldlessOps(noBodyOps, collectFieldlessActionOps(fs.STMLPages))
	opt := stmlgen.GenerateOptions{
		BearerAuth:              bearerAuth,
		RequestConstraints:      constraints,
		ResponseArrayItemFields: oapiparser.ExtractResponseArrayItemFields(fs.OpenAPIDoc),
		ResponseArrayItemTypes:  oapiparser.ExtractResponseArrayItemTypes(fs.OpenAPIDoc),
		ResponseBindTypes:       oapiparser.ExtractResponseFieldTypes(fs.OpenAPIDoc),
		NoBodyOps:               noBodyOps,
		PathParamTypes:          oapiparser.ExtractPathParamTypes(fs.OpenAPIDoc),
		RoutePatterns:           collectRoutePatterns(fs.STMLPages),
		DocumentTitles:          collectDocumentTitles(fs),
		CrumbFields:             stml.SitemapCrumbFields(fs.Sitemap),
		CrumbTitleSuffix:        crumbTitleSuffix(fs),
		ErrorDisplayField:       oapiparser.ExtractErrorDisplayField(fs.OpenAPIDoc),
	}
	_, err := stmlgen.Generate(fs.STMLPages, srcFrontendDir, pagesDir, opt)
	return err
}
