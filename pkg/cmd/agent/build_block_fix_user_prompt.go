//ff:func feature=agent type=helper control=selection
//ff:what buildBlockFixUserPrompt — 단일 블록 수정용 cross-SSOT user prompt 구성

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func buildBlockFixUserPrompt(specsDir string, ff *features.FeaturesFile, opID, desc, block string, diags []diagnostic.Diagnostic, l layer) string {
	var b strings.Builder

	lookup := buildFeatureLookupFromFF(ff)
	if desc == "" {
		if feat, ok := lookup[opID]; ok {
			desc = feat.Desc
		}
	}
	if desc != "" {
		fmt.Fprintf(&b, "Feature: %s\n\n", desc)
	}

	feat, hasFeat := lookup[opID]
	switch l {
	case layerOpenAPI:
		if hasFeat {
			writeDDLContext(&b, specsDir, feat.Table)
			writeSSaCForOp(&b, specsDir, opID)
		}
	case layerRego:
		if hasFeat {
			writeFeaturePublicContext(&b, feat)
			writeSSaCForOp(&b, specsDir, opID)
		}
	}

	fmt.Fprintf(&b, "OperationId: %s\n\n", opID)
	fmt.Fprintf(&b, "Current block:\n%s\n\n", block)

	writeDiagErrors(&b, diags)

	b.WriteString("\nFix ONLY this block. Output ONLY the corrected block content. Do not add surrounding file content.")
	return b.String()
}
