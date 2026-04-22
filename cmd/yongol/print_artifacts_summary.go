//ff:func feature=cli type=reporter control=sequence
//ff:what printArtifactsSummary — prints arts/backend file count, generated, preserved, and drift summary

package main

import (
	"fmt"
	"io"
	"path/filepath"
)

// printArtifactsSummary writes an "Artifacts (arts/backend)" block. It
// pairs the backend file count with already-computed preserved/drift
// counts and derives "generated" as total − preserved. The frontend
// line is a placeholder until TSX SSOT generation lands under
// pkg/generate/react/.
func printArtifactsSummary(w io.Writer, artsDir string, preservedCount, driftCount int) {
	fmt.Fprintf(w, "\nArtifacts (%s)\n", filepath.ToSlash(filepath.Join(artsDir, "backend")))
	total := countBackendFiles(artsDir)
	generated := total - preservedCount
	if generated < 0 {
		generated = 0
	}
	fmt.Fprintf(w, "  files        %d (generated=%d, preserved=%d, drift=%d)\n",
		total, generated, preservedCount, driftCount)
	fmt.Fprintf(w, "  frontend     -\n")
}
