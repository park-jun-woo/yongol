//ff:func feature=cli type=reporter control=iteration dimension=1
//ff:what printPreservedList — preserved 파일 경로·reason 목록 출력

package main

import (
	"fmt"
	"io"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

// printPreservedList writes a "Preserved Files (N)" section listing
// every preserved artifact path with its `//ff:preserve reason="..."`
// if present. Paths are sorted lexicographically so output is stable
// across runs regardless of filepath.WalkDir ordering.
func printPreservedList(w io.Writer, paths []string) {
	fmt.Fprintf(w, "\nPreserved Files (%d)\n", len(paths))
	if len(paths) == 0 {
		return
	}
	sorted := append([]string(nil), paths...)
	sort.Strings(sorted)
	for _, p := range sorted {
		fmt.Fprintf(w, "  %s\n", p)
		reason, err := contract.ParsePreserveReason(p)
		if err != nil || reason == "" {
			fmt.Fprintln(w, "    reason: <none>")
			continue
		}
		fmt.Fprintf(w, "    reason: %q\n", reason)
	}
}
