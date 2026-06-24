//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseAll — 멀티 도메인 모드에서 각 도메인 doc 의 제약이 전역 Request/ResponseConstraints 에 병합되는지 (Phase005 전제)
package yongol

import (
	"path/filepath"
	"testing"
)

// twoPopulatedDomainsManifest declares two fully-populated domains whose OpenAPI
// docs carry DISTINCT operationIds, so the merge must accumulate both.
const twoPopulatedDomainsManifest = `apiVersion: yongol/v1
kind: Project
metadata:
  name: multisite
backend:
  lang: go
  framework: gin
  module: example.com/multisite
domains:
  shop:
    openapi: api/shop.yaml
    frontend: frontend/shop
    route_prefix: /api
  blog:
    openapi: api/blog.yaml
    frontend: frontend/blog
    route_prefix: /blog
`

// blogDomainOpenAPI mirrors validDomainOpenAPI but with a different operationId
// (CreatePost) so the two domains contribute non-overlapping constraint keys.
const blogDomainOpenAPI = `openapi: 3.0.0
info:
  title: blog
  version: "1"
paths:
  /posts:
    post:
      operationId: CreatePost
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
                title:
                  type: string
                  maxLength: 20
      responses:
        "200":
          description: ok
          content:
            application/json:
              schema:
                type: object
                properties:
                  slug:
                    type: string
`

// TestParseAllMultiDomainConstraintMerge guards the Phase004→Phase005 bridge: in
// domain mode the per-domain OpenAPI docs' request/response constraints must be
// merged into the GLOBAL fs.RequestConstraints / fs.ResponseConstraints maps
// (opIDs are globally unique per XDO-90). Without this, Phase005's constraint-based
// validation sees empty global maps for every multi-domain project.
func TestParseAllMultiDomainConstraintMerge(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	writeFile(t, filepath.Join(tmp, "manifest.yaml"), twoPopulatedDomainsManifest)
	writeFile(t, filepath.Join(tmp, "api", "shop.yaml"), validDomainOpenAPI) // CreateThing
	writeFile(t, filepath.Join(tmp, "api", "blog.yaml"), blogDomainOpenAPI)  // CreatePost

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}

	fs := ParseAll(tmp, detected)
	if fs == nil {
		t.Fatalf("ParseAll returned nil")
	}
	if len(fs.ParseDiagnostics) != 0 {
		t.Fatalf("expected no parse diagnostics, got %d: %+v", len(fs.ParseDiagnostics), fs.ParseDiagnostics)
	}
	if !fs.IsDomained() {
		t.Fatalf("IsDomained() = false; want true")
	}

	// Both domains' OpenAPI docs must be loaded.
	if len(fs.DomainOpenAPIDocs) != 2 {
		t.Fatalf("DomainOpenAPIDocs count = %d; want 2", len(fs.DomainOpenAPIDocs))
	}

	// Global request constraints must contain BOTH operationIds (merge, not overwrite).
	if fs.RequestConstraints == nil {
		t.Fatalf("RequestConstraints nil — domain constraints were not merged into the global map")
	}
	if _, ok := fs.RequestConstraints["CreateThing"]; !ok {
		t.Errorf("RequestConstraints[CreateThing] missing — shop domain not merged; keys=%d", len(fs.RequestConstraints))
	}
	if _, ok := fs.RequestConstraints["CreatePost"]; !ok {
		t.Errorf("RequestConstraints[CreatePost] missing — blog domain not merged; keys=%d", len(fs.RequestConstraints))
	}

	// Global response constraints must likewise contain BOTH operationIds.
	if fs.ResponseConstraints == nil {
		t.Fatalf("ResponseConstraints nil — domain constraints were not merged into the global map")
	}
	if _, ok := fs.ResponseConstraints["CreateThing"]; !ok {
		t.Errorf("ResponseConstraints[CreateThing] missing — shop domain not merged; keys=%d", len(fs.ResponseConstraints))
	}
	if _, ok := fs.ResponseConstraints["CreatePost"]; !ok {
		t.Errorf("ResponseConstraints[CreatePost] missing — blog domain not merged; keys=%d", len(fs.ResponseConstraints))
	}
}
