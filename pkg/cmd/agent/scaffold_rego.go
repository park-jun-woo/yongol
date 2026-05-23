//ff:func feature=agent type=command control=iteration dimension=1
//ff:what scaffoldRego — non-public features로부터 OPA Rego authz 정책 자동 생성

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/docs"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

const regoBatchSize = 50

// scaffoldRego generates specs/policy/authz.rego from non-public features.
// Features are batched (max 50 per LLM call) and results are merged.
// Existing file is skipped.
func scaffoldRego(specsDir string, ff *features.FeaturesFile, llmFn LLMCallFunc, cfg Config, out io.Writer) error {
	outPath := filepath.Join(specsDir, "policy", "authz.rego")
	if _, err := os.Stat(outPath); err == nil {
		fmt.Fprintf(out, "  scaffold rego: skipped (exists)\n")
		return nil
	}

	var nonPublic []features.Feature
	for _, feat := range ff.Features {
		if !feat.Public {
			nonPublic = append(nonPublic, feat)
		}
	}
	if len(nonPublic) == 0 {
		fmt.Fprintf(out, "  scaffold rego: skipped (no non-public features)\n")
		return nil
	}

	policyDir := filepath.Join(specsDir, "policy")
	if err := os.MkdirAll(policyDir, 0755); err != nil {
		return fmt.Errorf("create policy dir: %w", err)
	}

	systemDoc, err := docs.FS.ReadFile("policy.md")
	if err != nil {
		return fmt.Errorf("read policy.md docs: %w", err)
	}
	systemPrompt := string(systemDoc)

	var ruleBlocks []string
	for i := 0; i < len(nonPublic); i += regoBatchSize {
		end := i + regoBatchSize
		if end > len(nonPublic) {
			end = len(nonPublic)
		}
		content, err := scaffoldRegoBatch(nonPublic[i:end], i/regoBatchSize, systemPrompt, cfg)
		if err != nil {
			return err
		}
		ruleBlocks = append(ruleBlocks, content)
	}

	doc := assembleRego(ruleBlocks)
	if err := os.WriteFile(outPath, []byte(doc), 0644); err != nil {
		return fmt.Errorf("scaffold rego write: %w", err)
	}

	fmt.Fprintf(out, "  scaffold rego: created authz.rego (%d non-public features)\n", len(nonPublic))
	return nil
}
