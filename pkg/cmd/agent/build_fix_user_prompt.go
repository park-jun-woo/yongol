//ff:func feature=agent type=helper control=selection
//ff:what buildFixUserPrompt — 레이어별 cross-SSOT 컨텍스트 포함 fix user prompt 구성

package agent

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func buildFixUserPrompt(specsDir string, ff *features.FeaturesFile, filename, content string, diags []diagnostic.Diagnostic, l layer) string {
	var b strings.Builder

	lookup := buildFeatureLookupFromFF(ff)
	desc := resolveDescFromFile(filename, l, lookup)
	if desc != "" {
		fmt.Fprintf(&b, "Feature: %s\n\n", desc)
	}

	switch l {
	case layerDDL:
		table := tableFromDDLFile(filename)
		writeFeatureTableContext(&b, ff, table)
		writeSQLcQueryContext(&b, specsDir, table)
	case layerSQLcQuery:
		table := tableFromSQLcFile(filename)
		writeDDLContext(&b, specsDir, table)
		writeOpenAPIPathContext(&b, specsDir, lookup, table)
	case layerSSaC:
		op := opFromSSaCFile(filename)
		if feat, ok := lookup[op]; ok {
			writeDDLContext(&b, specsDir, feat.Table)
			writeSQLcQueryContext(&b, specsDir, feat.Table)
			writeOpenAPIPathBlockContext(&b, specsDir, op)
		}
	case layerStateDiagram:
		table := tableFromStatesFile(filename)
		writeFeatureStatesContext(&b, ff, table)
		writeSSaCContextForTable(&b, specsDir, lookup, table)
	}

	fmt.Fprintf(&b, "Current file (%s):\n%s\n\n", filepath.Base(filename), content)

	writeDiagErrors(&b, diags)

	b.WriteString("\nFix the file. Output ONLY the corrected file content.")
	return b.String()
}
