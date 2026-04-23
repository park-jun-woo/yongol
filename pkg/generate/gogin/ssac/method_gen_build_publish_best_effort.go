//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.buildPublishBestEffort — UseTx=false 경로. queue.Publish + slog.Error best-effort 라인 생성
package ssac

import (
	"fmt"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func (g *methodGen) buildPublishBestEffort(seq ssacparser.Sequence, fields []string, imports *[]string) []string {
	*imports = append(*imports, `"log/slog"`)
	lines := []string{
		fmt.Sprintf("if err := queue.Publish(ctx, %q, map[string]any{", seq.Topic),
	}
	lines = append(lines, fields...)
	lines = append(lines,
		"}); err != nil {",
		fmt.Sprintf("\tslog.Error(\"publish failed\", \"op\", %q, \"topic\", %q, \"err\", err)", g.FuncName, seq.Topic),
		"}",
	)
	return lines
}
