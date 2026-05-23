//ff:func feature=cli type=helper control=sequence
//ff:what printDiagFile — 진단 파일 경로+줄번호를 출력
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func printDiagFile(out io.Writer, d diagnostic.Diagnostic) {
	if d.File == "" {
		return
	}
	if d.Line > 0 {
		fmt.Fprintf(out, "    file: %s:%d\n", d.File, d.Line)
	} else {
		fmt.Fprintf(out, "    file: %s\n", d.File)
	}
}
