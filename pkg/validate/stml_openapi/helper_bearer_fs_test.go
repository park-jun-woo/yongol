//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what makeAuthFS — backend.auth mode를 가진 테스트용 Fullstack fixture 생성

package stml_openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// makeAuthFS builds a Fullstack like makeFS but with backend.auth declared
// in the given mode ("bearer" / "cookie" / "hybrid"), for the TM-21/22/24
// flow rules.
func makeAuthFS(pages []stml.PageSpec, doc *openapi3.T, mode string) *yongol.Fullstack {
	fs := makeFS(pages, doc)
	fs.Manifest.Backend.Auth = &manifest.Auth{Type: "jwt", Mode: mode}
	return fs
}
