//ff:func feature=orchestrator type=test control=iteration dimension=2
//ff:what ParseAll — domains 블록 멀티 도메인 프로젝트의 도메인별 OpenAPI/STML/sitemap/layout 적재 + presence 파생 회귀
package yongol

import (
	"path/filepath"
	"testing"
)

// multidomainManifest declares two domains: "shop" (fully populated) and
// "admin" (every per-domain SSOT absent). It drives parseDomainsIfPresent through
// both the present and the absent branch of every per-domain loader in one pass.
const multidomainManifest = `apiVersion: yongol/v1
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
  admin:
    openapi: api/admin.yaml
    frontend: frontend/admin
    route_prefix: /api/admin
`

// validDomainOpenAPI is a minimal-but-real OpenAPI 3 document with one operation
// that carries a request-body constraint and a 200 response schema, so the loaded
// doc is non-nil and the line index is built.
const validDomainOpenAPI = `openapi: 3.0.0
info:
  title: shop
  version: "1"
paths:
  /things:
    post:
      operationId: CreateThing
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
                name:
                  type: string
                  maxLength: 10
      responses:
        "200":
          description: ok
          content:
            application/json:
              schema:
                type: object
                properties:
                  id:
                    type: string
`

const validDomainPage = `<main data-route="/things">
  <article data-fetch="ListThings">
    <span data-bind="Name"></span>
  </article>
</main>`

const validDomainSitemap = `<nav data-sitemap>
  <ul>
    <li data-page="things">Things</li>
  </ul>
</nav>`

const validDomainLayout = `<div>
  <slot data-outlet />
</div>`

// TestParseAllMultiDomain verifies the Phase004 domain loading: a manifest with a
// domains: block makes ParseAll populate the per-domain maps from a single pass,
// with presence derived from the same pass (present "shop" vs absent "admin").
func TestParseAllMultiDomain(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	writeFile(t, filepath.Join(tmp, "manifest.yaml"), multidomainManifest)
	// shop — fully populated per-domain SSOTs.
	writeFile(t, filepath.Join(tmp, "api", "shop.yaml"), validDomainOpenAPI)
	writeFile(t, filepath.Join(tmp, "frontend", "shop", "home.html"), validDomainPage)
	writeFile(t, filepath.Join(tmp, "frontend", "shop", "sitemap.html"), validDomainSitemap)
	writeFile(t, filepath.Join(tmp, "frontend", "shop", "layouts", "app.html"), validDomainLayout)
	// admin — intentionally NOTHING on disk: api/admin.yaml + frontend/admin absent.

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
		t.Fatalf("IsDomained() = false; want true for a manifest with a domains block")
	}

	// OpenAPI: shop present (doc + line index), admin absent.
	if _, ok := fs.DomainOpenAPIDocs["shop"]; !ok {
		t.Errorf("DomainOpenAPIDocs[shop] missing — per-domain OpenAPI not loaded")
	}
	if _, ok := fs.DomainOpenAPILines["shop"]; !ok {
		t.Errorf("DomainOpenAPILines[shop] missing — per-domain line index not built")
	}
	if _, ok := fs.DomainOpenAPIDocs["admin"]; ok {
		t.Errorf("DomainOpenAPIDocs[admin] present — absent OpenAPI must not be loaded")
	}

	// STML pages: shop populated, admin absent.
	if len(fs.DomainSTMLPages["shop"]) == 0 {
		t.Errorf("DomainSTMLPages[shop] empty — per-domain STML pages not loaded")
	}
	if _, ok := fs.DomainSTMLPages["admin"]; ok {
		t.Errorf("DomainSTMLPages[admin] present — absent frontend must not be loaded")
	}

	// Sitemap: shop populated, admin absent.
	if fs.DomainSitemaps["shop"] == nil {
		t.Errorf("DomainSitemaps[shop] nil — per-domain sitemap not loaded")
	}
	if _, ok := fs.DomainSitemaps["admin"]; ok {
		t.Errorf("DomainSitemaps[admin] present — absent sitemap must not be loaded")
	}

	// Layouts: shop populated, admin absent.
	if len(fs.DomainLayouts["shop"]) == 0 {
		t.Errorf("DomainLayouts[shop] empty — per-domain layouts not loaded")
	}
	if _, ok := fs.DomainLayouts["admin"]; ok {
		t.Errorf("DomainLayouts[admin] present — absent layouts must not be loaded")
	}

	// Presences derived from the same pass: shop populated, admin absent.
	if fs.DomainPresences["shop"][KindOpenAPI] != SSOTPopulated {
		t.Errorf("DomainPresences[shop][OpenAPI] = %v; want SSOTPopulated", fs.DomainPresences["shop"][KindOpenAPI])
	}
	if fs.DomainPresences["shop"][KindSTML] != SSOTPopulated {
		t.Errorf("DomainPresences[shop][STML] = %v; want SSOTPopulated", fs.DomainPresences["shop"][KindSTML])
	}
	if fs.DomainPresences["admin"][KindOpenAPI] != SSOTAbsent {
		t.Errorf("DomainPresences[admin][OpenAPI] = %v; want SSOTAbsent", fs.DomainPresences["admin"][KindOpenAPI])
	}
	if fs.DomainPresences["admin"][KindSTML] != SSOTAbsent {
		t.Errorf("DomainPresences[admin][STML] = %v; want SSOTAbsent", fs.DomainPresences["admin"][KindSTML])
	}

	// Presence never drifts from loaded data: a populated OpenAPI presence implies
	// a loaded doc, an absent one implies no doc.
	for name, pres := range fs.DomainPresences {
		_, hasDoc := fs.DomainOpenAPIDocs[name]
		if (pres[KindOpenAPI] == SSOTPopulated) != hasDoc {
			t.Errorf("domain %s: OpenAPI presence %v drifted from DomainOpenAPIDocs membership %v", name, pres[KindOpenAPI], hasDoc)
		}
	}
}
