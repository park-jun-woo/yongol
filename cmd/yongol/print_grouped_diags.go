//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what printGroupedDiags — 같은 OperationID의 진단을 묶어서 출력
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func printGroupedDiags(out io.Writer, opID string, all []diagnostic.Diagnostic, specsDir string) {
	var group []diagnostic.Diagnostic
	for _, d := range all {
		if d.OperationID == opID {
			group = append(group, d)
		}
	}
	fmt.Fprintf(out, "\n%s (%d errors):\n", opID, len(group))
	for _, d := range group {
		fmt.Fprintf(out, "  [%s] %s\n", d.Level, d.Message)
		printDiagFile(out, d)
	}
	fmt.Fprintf(out, "  ▶ Fix all errors for %s. Then run `yongol next %s`. Don't stop. Keep going until complete.\n", opID, specsDir)
}
