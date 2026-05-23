//ff:func feature=agent type=helper control=sequence
//ff:what buildFixSystemPrompt — docs 기반 레이어별 fix system prompt 구성

package agent

import (
	"github.com/park-jun-woo/yongol/docs"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func buildFixSystemPrompt(l layer, diags []diagnostic.Diagnostic) string {
	filename := layerDocFile(l)
	if filename == "" {
		msgs := diagMessages(diags)
		return buildSystemPrompt(l, msgs)
	}
	data, err := docs.FS.ReadFile(filename)
	if err != nil {
		msgs := diagMessages(diags)
		return buildSystemPrompt(l, msgs)
	}
	return string(data)
}
