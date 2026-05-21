//ff:func feature=agent type=command control=iteration dimension=1
//ff:what scaffoldRego — non-public features로부터 OPA Rego authz 정책 자동 생성

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

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

	// Collect non-public features
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

	// Batch features and call LLM
	var ruleBlocks []string
	for i := 0; i < len(nonPublic); i += regoBatchSize {
		end := i + regoBatchSize
		if end > len(nonPublic) {
			end = len(nonPublic)
		}
		batch := nonPublic[i:end]

		userPrompt := buildRegoUserPrompt(batch)
		numCtx := len(systemPrompt) + len(userPrompt) + 2048

		reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, systemPrompt, userPrompt, numCtx)
		if err != nil {
			return fmt.Errorf("scaffold rego batch %d: %w", i/regoBatchSize, err)
		}

		content := stripCodeBlock(reply)
		if content == "" {
			return fmt.Errorf("scaffold rego batch %d: empty LLM response", i/regoBatchSize)
		}

		ruleBlocks = append(ruleBlocks, content)
	}

	// Assemble final rego file
	doc := assembleRego(ruleBlocks)

	if err := os.WriteFile(outPath, []byte(doc), 0644); err != nil {
		return fmt.Errorf("scaffold rego write: %w", err)
	}

	fmt.Fprintf(out, "  scaffold rego: created authz.rego (%d non-public features)\n", len(nonPublic))
	return nil
}

// buildRegoUserPrompt builds the user prompt for a batch of features.
func buildRegoUserPrompt(feats []features.Feature) string {
	var b strings.Builder

	b.WriteString("Non-public features requiring authorization:\n\n")
	for _, f := range feats {
		resource := f.Table
		if resource == "" {
			resource = domainFromPath(f.Path)
		}
		fmt.Fprintf(&b, "  - op: %s, resource: %s\n", f.Op, resource)
	}

	b.WriteString("\nGenerate OPA Rego allow rules for these features.")
	b.WriteString("\nUse 'allow if { ... }' syntax with input.action == \"<operationId>\" checks.")
	b.WriteString("\nOutput ONLY the allow rule blocks. No package declaration. No import. No markdown fences.")

	return b.String()
}

// assembleRego combines rule blocks into a complete Rego policy file.
func assembleRego(ruleBlocks []string) string {
	var b strings.Builder

	b.WriteString("package authz\n\nimport rego.v1\n\ndefault allow := false\n")

	for _, block := range ruleBlocks {
		block = strings.TrimSpace(block)
		// Strip duplicate package/import/default lines from LLM output
		lines := strings.Split(block, "\n")
		var filtered []string
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "package ") {
				continue
			}
			if strings.HasPrefix(trimmed, "import ") {
				continue
			}
			if strings.HasPrefix(trimmed, "default allow") {
				continue
			}
			filtered = append(filtered, line)
		}
		cleaned := strings.TrimSpace(strings.Join(filtered, "\n"))
		if cleaned != "" {
			b.WriteString("\n")
			b.WriteString(cleaned)
			b.WriteString("\n")
		}
	}

	return b.String()
}
