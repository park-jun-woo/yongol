//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what TestSSaCManifestHelpers — unit tests for the pure ssac_manifest helper functions
package ssac_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestUsesFile(t *testing.T) {
	if !usesFile(fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("file.Save")}})) {
		t.Error("file. prefix expected")
	}
	if !usesFile(fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("storage.Upload")}})) {
		t.Error("storage. prefix expected")
	}
	if usesFile(fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})) {
		t.Error("cache.Get should not count as file use")
	}
	// non-call sequence ignored.
	if usesFile(fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "get", Model: "file.X"}}})) {
		t.Error("non-call seq should be ignored")
	}
}
