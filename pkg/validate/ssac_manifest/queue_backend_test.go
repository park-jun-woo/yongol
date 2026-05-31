//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what TestSSaCManifestHelpers — unit tests for the pure ssac_manifest helper functions
package ssac_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQueueBackend(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Queue: &manifest.QueueBackend{Backend: "postgres"}}}
	if got := queueBackend(fs); got != "postgres" {
		t.Errorf("got %q, want postgres", got)
	}
	// nil manifest → "".
	if got := queueBackend(&yongol.Fullstack{}); got != "" {
		t.Errorf("nil manifest: %q", got)
	}
	// nil Queue → "".
	if got := queueBackend(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}); got != "" {
		t.Errorf("nil queue: %q", got)
	}
}
