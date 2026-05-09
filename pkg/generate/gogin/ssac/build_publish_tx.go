//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildPublishTx — UseTx=true 경로. queue.PublishTx(ctx, tx, ...) 호출 + err 전파 라인 생성
package ssac

import (
	"fmt"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func buildPublishTx(seq ssacparser.Sequence, fields []string, _ []string, isSubscribe bool) []string {
	retErr := "return nil, err"
	if isSubscribe {
		retErr = "return err"
	}
	lines := []string{
		fmt.Sprintf("if err := queue.PublishTx(ctx, tx, %q, map[string]any{", seq.Topic),
	}
	lines = append(lines, fields...)
	lines = append(lines,
		"}); err != nil {",
		"\t"+retErr,
		"}",
	)
	return lines
}
