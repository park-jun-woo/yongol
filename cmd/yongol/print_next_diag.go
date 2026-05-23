//ff:func feature=cli type=helper control=sequence
//ff:what printNextDiag — 단일 diagnostic을 next 형식으로 출력 (1줄 에러 + file + fix 지시)
package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func printNextDiag(cmd *cobra.Command, d diagnostic.Diagnostic, specsDir string) {
	out := cmd.OutOrStdout()
	fmt.Fprintf(out, "\n[%s] %s\n", d.Level, d.Message)
	if d.File != "" {
		if d.Line > 0 {
			fmt.Fprintf(out, "  file: %s:%d\n", d.File, d.Line)
		} else {
			fmt.Fprintf(out, "  file: %s\n", d.File)
		}
	}
	fmt.Fprintf(out, "  ▶ Fix this error. Then run `yongol next %s`. Don't stop. Keep going until complete.\n", specsDir)
}
