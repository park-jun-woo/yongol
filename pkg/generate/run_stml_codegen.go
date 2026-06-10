//ff:func feature=generate type=generator control=sequence
//ff:what runSTMLCodegen — STML 페이지 스펙에서 React TSX 페이지 파일 생성

package generate

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	stmlgen "github.com/park-jun-woo/yongol/pkg/generate/react/stml"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runSTMLCodegen generates page TSX files from STML specs when STMLPages
// are present. Skips silently if no STML pages were parsed.
func runSTMLCodegen(fs *yongol.Fullstack, artifactsDir string) error {
	if fs == nil || len(fs.STMLPages) == 0 {
		return nil
	}
	pagesDir := filepath.Join(artifactsDir, "frontend", "src", "pages")
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
		NoBodyOps:               noBodyOps,
		PathParamTypes:          oapiparser.ExtractPathParamTypes(fs.OpenAPIDoc),
	}
	_, err := stmlgen.Generate(fs.STMLPages, fs.SpecsDir, pagesDir, opt)
	return err
}
