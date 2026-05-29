//ff:func feature=gen-gogin type=test control=sequence
//ff:what corsEnabled — manifest.backend.cors.enabled 여부 판정

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCorsEnabled(t *testing.T) {
	if corsEnabled(nil) {
		t.Errorf("nil fs should be false")
	}
	if corsEnabled(&yongol.Fullstack{}) {
		t.Errorf("nil manifest should be false")
	}
	disabled := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{CORS: &pmanifest.CORSConfig{Enabled: false}},
	}}
	if corsEnabled(disabled) {
		t.Errorf("cors.enabled=false should be false")
	}
	enabled := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{CORS: &pmanifest.CORSConfig{Enabled: true}},
	}}
	if !corsEnabled(enabled) {
		t.Errorf("cors.enabled=true should be true")
	}
}
