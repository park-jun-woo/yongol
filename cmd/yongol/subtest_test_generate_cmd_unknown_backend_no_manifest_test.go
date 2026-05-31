//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestGenerateCmdUnknownBackendNoManifest — UnknownBackendNoManifest 서브테스트
package main

import (
	"testing"
)

func subtestTestGenerateCmdUnknownBackendNoManifest(t *testing.T) {

	// An empty (but valid) specs dir parses clean and has no manifest, so an
	// unrecognized --backend falls through to the "unknown --backend" error
	// rather than manifest resolution.
	dir := t.TempDir()
	_, _, err := runCmd(t, "generate", "--backend", "no-such-backend", dir, t.TempDir())
	if err == nil {
		t.Fatal("expected unknown backend error, got nil")
	}

}
