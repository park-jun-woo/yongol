//ff:func feature=agent type=command control=iteration dimension=1
//ff:what fixSplitFile — OpenAPI/Rego/Hurl 블록 단위 수정 (operationId 기준)

package agent

import (
	"fmt"
	"os"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func fixSplitFile(specsDir string, ff *features.FeaturesFile, relFile, absPath, fileContent string, diags []diagnostic.Diagnostic, llmFn LLMCallFunc, cfg Config, l layer) error {
	opIDs := extractOperationIDs(diags)
	if len(opIDs) == 0 {
		return fmt.Errorf("no operationId in diagnostics for %s", relFile)
	}

	msgs := diagMessages(diags)
	lookup := buildFeatureLookupFromFF(ff)

	for _, opID := range opIDs {
		if err := fixSplitFileOp(specsDir, ff, &fileContent, opID, diags, msgs, lookup, llmFn, cfg, l); err != nil {
			continue
		}
	}

	return os.WriteFile(absPath, []byte(fileContent), 0644)
}
