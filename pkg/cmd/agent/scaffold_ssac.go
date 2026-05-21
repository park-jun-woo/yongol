//ff:func feature=agent type=command control=iteration dimension=1
//ff:what scaffoldSSaC — features.yaml ops로부터 SSaC 서비스 파일 자동 생성

package agent

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/docs"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// scaffoldSSaC generates specs/service/{domain}/{Op}.ssac from features.yaml.
// Each feature op produces one LLM call that returns one SSaC file.
// Existing files are skipped to protect user modifications.
func scaffoldSSaC(specsDir string, ff *features.FeaturesFile, openapiContent string, llmFn LLMCallFunc, cfg Config, out io.Writer) (int, error) {
	if len(ff.Features) == 0 {
		return 0, nil
	}

	systemDoc, err := docs.FS.ReadFile("ssac.md")
	if err != nil {
		return 0, fmt.Errorf("read ssac.md docs: %w", err)
	}
	systemPrompt := string(systemDoc)

	count := 0
	for _, feat := range ff.Features {
		domain := domainFromPath(feat.Path)
		serviceDir := filepath.Join(specsDir, "service", domain)
		outPath := filepath.Join(serviceDir, feat.Op+".ssac")

		if _, err := os.Stat(outPath); err == nil {
			fmt.Fprintf(out, "  scaffold ssac: skipped %s/%s.ssac (exists)\n", domain, feat.Op)
			continue
		}

		if err := os.MkdirAll(serviceDir, 0755); err != nil {
			return 0, fmt.Errorf("create service/%s dir: %w", domain, err)
		}

		ddlContent := readDDLForTable(specsDir, feat.Table)
		queryNames := readSQLcQueryNames(specsDir, feat.Table)
		pathBlock := extractPathBlockForOp(openapiContent, feat.Op)
		userPrompt := buildSSaCUserPrompt(feat, ddlContent, queryNames, pathBlock)

		numCtx := int(float64(len(systemPrompt)+len(userPrompt))/4*1.5) + 2048

		reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, systemPrompt, userPrompt, numCtx)
		if err != nil {
			return 0, fmt.Errorf("scaffold ssac %s: %w", feat.Op, err)
		}

		content := stripCodeBlock(reply)
		if content == "" {
			return 0, fmt.Errorf("scaffold ssac %s: empty LLM response", feat.Op)
		}

		if err := os.WriteFile(outPath, []byte(content+"\n"), 0644); err != nil {
			return 0, fmt.Errorf("scaffold ssac %s: write: %w", feat.Op, err)
		}

		fmt.Fprintf(out, "  scaffold ssac: created %s/%s.ssac\n", domain, feat.Op)
		count++
	}

	return count, nil
}

// buildSSaCUserPrompt builds the user prompt for generating a single SSaC file.
func buildSSaCUserPrompt(feat features.Feature, ddlContent string, queryNames []string, pathBlock string) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Feature:\n")
	fmt.Fprintf(&b, "  op: %s\n", feat.Op)
	fmt.Fprintf(&b, "  path: %s\n", feat.Path)
	fmt.Fprintf(&b, "  table: %s\n", feat.Table)
	fmt.Fprintf(&b, "  public: %v\n", feat.Public)
	fmt.Fprintf(&b, "  desc: %s\n", feat.Desc)

	if ddlContent != "" {
		fmt.Fprintf(&b, "\nDDL for table %s:\n%s\n", feat.Table, ddlContent)
	}

	if len(queryNames) > 0 {
		b.WriteString("\nAvailable sqlc queries:\n")
		for _, name := range queryNames {
			fmt.Fprintf(&b, "  %s\n", name)
		}
	}

	if pathBlock != "" {
		fmt.Fprintf(&b, "\nOpenAPI path block:\n%s\n", pathBlock)
	}

	b.WriteString("\nGenerate a single SSaC service file for this feature.")
	b.WriteString("\nOutput ONLY the SSaC file content. No explanations. No markdown fences.")

	return b.String()
}

// readSQLcQueryNames reads the sqlc query file for a table and returns
// all "-- name:" lines (query name declarations).
func readSQLcQueryNames(specsDir, tableName string) []string {
	if tableName == "" {
		return nil
	}
	path := filepath.Join(specsDir, "db", "queries", tableName+".sql")
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	var names []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "-- name:") {
			names = append(names, line)
		}
	}
	return names
}

// domainFromPath extracts the domain name from an API path.
// /workflows/{id} -> workflow
// /auth/login -> auth
// /payment-intents/{id} -> payment_intent
func domainFromPath(path string) string {
	fields := strings.Fields(path)
	if len(fields) >= 2 {
		path = fields[1]
	}
	path = strings.TrimPrefix(path, "/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) == 0 || parts[0] == "" {
		return "default"
	}

	segment := parts[0]
	// Convert kebab-case to snake_case
	segment = strings.ReplaceAll(segment, "-", "_")
	// Remove trailing 's' for plural to singular (simple heuristic)
	if strings.HasSuffix(segment, "s") && !strings.HasSuffix(segment, "ss") {
		segment = segment[:len(segment)-1]
	}
	return segment
}
