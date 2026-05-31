//ff:func feature=gen-ir type=test control=sequence
//ff:what convertPublish/resolveExposeInternal/isCountResultType/ddlTableSingularIR/DDLTableSingularIR/findDDLTable
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveExposeInternal(t *testing.T) {
	if resolveExposeInternal(nil) {
		t.Errorf("nil should be false")
	}
	if resolveExposeInternal(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}) {
		t.Errorf("no error config should be false")
	}
	tru := true
	on := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{
		Error: &manifest.ErrorConfig{ExposeInternalError: &tru},
	}}}
	if !resolveExposeInternal(on) {
		t.Errorf("explicit true should be true")
	}
}
