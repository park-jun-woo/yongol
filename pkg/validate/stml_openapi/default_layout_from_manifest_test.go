//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestStmlOpenAPIHelpers — unit tests for the pure stml_openapi helper functions
package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestDefaultLayoutFromManifest(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	fs.Manifest.Frontend.DefaultLayout = "main"
	if got := defaultLayoutFromManifest(fs); got != "main" {
		t.Errorf("got %q, want main", got)
	}
	// nil manifest → "".
	if got := defaultLayoutFromManifest(&yongol.Fullstack{}); got != "" {
		t.Errorf("nil manifest: %q", got)
	}
}
