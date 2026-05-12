//ff:func feature=generate type=generator control=sequence
//ff:what runSTMLCodegen — STML 페이지 스펙에서 React TSX 페이지 파일 생성

package generate

import (
	"path/filepath"

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
	var hasAuthz bool
	if fs.Manifest != nil {
		hasAuthz = fs.Manifest.Authz != nil
	}
	opt := stmlgen.GenerateOptions{
		HasAuthz:                hasAuthz,
		RequestConstraints:      fs.RequestConstraints,
		ResponseArrayItemFields: oapiparser.ExtractResponseArrayItemFields(fs.OpenAPIDoc),
		NoBodyOps:               oapiparser.ExtractNoBodyOps(fs.OpenAPIDoc),
	}
	_, err := stmlgen.Generate(fs.STMLPages, fs.SpecsDir, pagesDir, opt)
	return err
}
