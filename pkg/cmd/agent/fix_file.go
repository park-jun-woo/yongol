//ff:func feature=agent type=command control=sequence
//ff:what fixFile — 단일 파일의 validate 에러를 LLM으로 수정 (레이어별 컨텍스트 주입)

package agent

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/docs"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// fixFile reads the target file, builds layer-specific context, calls LLM to
// fix validate errors, and writes the result back. For OpenAPI and Rego the
// fix is done per-block (operationId granularity) using existing split logic.
func fixFile(specsDir string, ff *features.FeaturesFile, filename string, diags []diagnostic.Diagnostic, llmFn LLMCallFunc, cfg Config) error {
	l := classifyFile(filename)
	absPath := filepath.Join(specsDir, filename)

	content, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("read %s: %w", filename, err)
	}

	// For split-based layers (OpenAPI, Rego, Hurl) delegate to block-level fix.
	if l == layerOpenAPI || l == layerRego || l == layerHurl {
		return fixSplitFile(specsDir, ff, filename, absPath, string(content), diags, llmFn, cfg, l)
	}

	// Whole-file fix for other layers.
	systemPrompt := buildFixSystemPrompt(l, diags)
	userPrompt := buildFixUserPrompt(specsDir, ff, filename, string(content), diags, l)

	reply, err := llmFn(cfg.Backend, cfg.Model, systemPrompt, userPrompt)
	if err != nil {
		return fmt.Errorf("LLM %s: %w", filename, err)
	}

	fixed := stripMarkdownFences(reply)
	if fixed == "" {
		return fmt.Errorf("empty LLM response for %s", filename)
	}

	return os.WriteFile(absPath, []byte(fixed), 0644)
}

// fixSplitFile handles OpenAPI/Rego/Hurl files by extracting per-operationId
// blocks, fixing each, and merging back.
func fixSplitFile(specsDir string, ff *features.FeaturesFile, relFile, absPath, fileContent string, diags []diagnostic.Diagnostic, llmFn LLMCallFunc, cfg Config, l layer) error {
	opIDs := extractOperationIDs(diags)
	if len(opIDs) == 0 {
		return fmt.Errorf("no operationId in diagnostics for %s", relFile)
	}

	msgs := diagMessages(diags)
	lookup := buildFeatureLookupFromFF(ff)

	for _, opID := range opIDs {
		opMsgs := filterMessagesByOp(msgs, opID)
		if len(opMsgs) == 0 {
			opMsgs = msgs
		}

		// Filter diagnostics to this op for advice.
		var opDiags []diagnostic.Diagnostic
		for _, d := range diags {
			if strings.Contains(d.Message, opID) {
				opDiags = append(opDiags, d)
			}
		}
		if len(opDiags) == 0 {
			opDiags = diags
		}

		desc := ""
		if feat, ok := lookup[opID]; ok {
			desc = feat.Desc
		}

		systemPrompt := buildFixSystemPrompt(l, opDiags)

		var block string
		var startLine, endLine int
		var extractErr error

		switch l {
		case layerOpenAPI:
			block, startLine, endLine, extractErr = extractOpenAPIBlock(fileContent, opID)
		case layerRego:
			block, startLine, endLine, extractErr = extractRegoBlock(fileContent, opID)
		case layerHurl:
			block, startLine, endLine, extractErr = extractHurlBlock(fileContent, opID)
		}

		if extractErr != nil {
			continue // block not found — skip
		}

		// Build user prompt with cross-SSOT context.
		userPrompt := buildBlockFixUserPrompt(specsDir, ff, opID, desc, block, opDiags, l)

		reply, err := llmFn(cfg.Backend, cfg.Model, systemPrompt, userPrompt)
		if err != nil {
			continue
		}

		fixedBlock := stripMarkdownFences(reply)
		if fixedBlock == "" {
			continue
		}

		var merged string
		switch l {
		case layerOpenAPI:
			merged, err = mergeOpenAPIBlock(fileContent, startLine, endLine, fixedBlock)
		case layerRego:
			merged, err = mergeRegoBlock(fileContent, startLine, endLine, fixedBlock)
		case layerHurl:
			merged, err = mergeHurlBlock(fileContent, startLine, endLine, fixedBlock)
		}
		if err != nil {
			continue
		}
		fileContent = merged
	}

	return os.WriteFile(absPath, []byte(fileContent), 0644)
}

// buildFixSystemPrompt returns a system prompt loaded from embedded docs for
// the given layer, falling back to the generic prompt builder.
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

// buildFixUserPrompt builds a user prompt with layer-specific cross-SSOT
// context for whole-file fix. See the context table in PhaseC14 spec.
func buildFixUserPrompt(specsDir string, ff *features.FeaturesFile, filename, content string, diags []diagnostic.Diagnostic, l layer) string {
	var b strings.Builder

	// Resolve feature description from the file.
	lookup := buildFeatureLookupFromFF(ff)
	desc := resolveDescFromFile(filename, l, lookup)
	if desc != "" {
		fmt.Fprintf(&b, "Feature: %s\n\n", desc)
	}

	// Layer-specific context.
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

	// Current file content.
	fmt.Fprintf(&b, "Current file (%s):\n%s\n\n", filepath.Base(filename), content)

	// Validate errors with advice.
	b.WriteString("Validate errors:\n")
	for _, d := range diags {
		fmt.Fprintf(&b, "- %s\n", d.Message)
		if d.Advice != "" {
			fmt.Fprintf(&b, "  Advice: %s\n", d.Advice)
		}
	}

	b.WriteString("\nFix the file. Output ONLY the corrected file content.")
	return b.String()
}

// buildBlockFixUserPrompt builds a user prompt for a single block fix with
// cross-SSOT context.
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

	// Layer-specific context for the operationId.
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

	b.WriteString("Validate errors:\n")
	for _, d := range diags {
		fmt.Fprintf(&b, "- %s\n", d.Message)
		if d.Advice != "" {
			fmt.Fprintf(&b, "  Advice: %s\n", d.Advice)
		}
	}

	b.WriteString("\nFix ONLY this block. Output ONLY the corrected block content. Do not add surrounding file content.")
	return b.String()
}

// --- Context writer helpers ---

func writeDDLContext(b *strings.Builder, specsDir, table string) {
	if table == "" {
		return
	}
	ddl := readDDLForTable(specsDir, table)
	if ddl != "" {
		fmt.Fprintf(b, "DDL (%s.sql):\n%s\n\n", table, ddl)
	}
}

func writeSQLcQueryContext(b *strings.Builder, specsDir, table string) {
	names := readSQLcQueryNames(specsDir, table)
	if len(names) > 0 {
		b.WriteString("sqlc queries:\n")
		for _, n := range names {
			fmt.Fprintf(b, "  %s\n", n)
		}
		b.WriteByte('\n')
	}
}

func writeFeatureTableContext(b *strings.Builder, ff *features.FeaturesFile, table string) {
	if ff == nil || table == "" {
		return
	}
	var related []features.Feature
	for _, f := range ff.Features {
		if f.Table == table {
			related = append(related, f)
		}
	}
	if len(related) > 0 {
		b.WriteString("Related features:\n")
		for _, f := range related {
			fmt.Fprintf(b, "  - %s %s: %s\n", f.Op, f.Path, f.Desc)
		}
		b.WriteByte('\n')
	}
}

func writeOpenAPIPathContext(b *strings.Builder, specsDir string, lookup map[string]features.Feature, table string) {
	openapiPath := filepath.Join(specsDir, "api", "openapi.yaml")
	data, err := os.ReadFile(openapiPath)
	if err != nil {
		return
	}
	content := string(data)
	// Find ops for this table.
	for op, feat := range lookup {
		if feat.Table == table {
			block := extractPathBlockForOp(content, op)
			if block != "" {
				fmt.Fprintf(b, "OpenAPI path block (%s):\n%s\n", op, block)
			}
		}
	}
}

func writeOpenAPIPathBlockContext(b *strings.Builder, specsDir, opID string) {
	openapiPath := filepath.Join(specsDir, "api", "openapi.yaml")
	data, err := os.ReadFile(openapiPath)
	if err != nil {
		return
	}
	block := extractPathBlockForOp(string(data), opID)
	if block != "" {
		fmt.Fprintf(b, "OpenAPI path block (%s):\n%s\n", opID, block)
	}
}

func writeSSaCForOp(b *strings.Builder, specsDir, opID string) {
	content, ok := findSSaCFile(specsDir, opID)
	if ok && content != "" {
		fmt.Fprintf(b, "SSaC (%s.ssac):\n%s\n\n", opID, content)
	}
}

func writeSSaCContextForTable(b *strings.Builder, specsDir string, lookup map[string]features.Feature, table string) {
	for op, feat := range lookup {
		if feat.Table == table {
			writeSSaCForOp(b, specsDir, op)
		}
	}
}

func writeFeatureStatesContext(b *strings.Builder, ff *features.FeaturesFile, table string) {
	if ff == nil || table == "" {
		return
	}
	td, ok := ff.Tables[table]
	if !ok || len(td.States) == 0 {
		return
	}
	fmt.Fprintf(b, "States for %s: %s\n\n", table, strings.Join(td.States, ", "))
}

func writeFeaturePublicContext(b *strings.Builder, feat features.Feature) {
	fmt.Fprintf(b, "Feature public: %v\n\n", feat.Public)
}

// --- File path → entity resolution helpers ---

func tableFromDDLFile(relPath string) string {
	base := filepath.Base(relPath)
	return strings.TrimSuffix(base, ".sql")
}

func tableFromSQLcFile(relPath string) string {
	base := filepath.Base(relPath)
	return strings.TrimSuffix(base, ".sql")
}

func tableFromStatesFile(relPath string) string {
	base := filepath.Base(relPath)
	return strings.TrimSuffix(base, ".md")
}

func opFromSSaCFile(relPath string) string {
	base := filepath.Base(relPath)
	return strings.TrimSuffix(base, ".ssac")
}

func resolveDescFromFile(relPath string, l layer, lookup map[string]features.Feature) string {
	switch l {
	case layerSSaC:
		op := opFromSSaCFile(relPath)
		if f, ok := lookup[op]; ok {
			return f.Desc
		}
	case layerDDL:
		table := tableFromDDLFile(relPath)
		for _, f := range lookup {
			if f.Table == table {
				return f.Desc
			}
		}
	case layerSQLcQuery:
		table := tableFromSQLcFile(relPath)
		for _, f := range lookup {
			if f.Table == table {
				return f.Desc
			}
		}
	}
	return ""
}

// buildFeatureLookupFromFF builds op → Feature map from FeaturesFile.
func buildFeatureLookupFromFF(ff *features.FeaturesFile) map[string]features.Feature {
	if ff == nil {
		return nil
	}
	m := make(map[string]features.Feature, len(ff.Features))
	for _, f := range ff.Features {
		m[f.Op] = f
	}
	return m
}
