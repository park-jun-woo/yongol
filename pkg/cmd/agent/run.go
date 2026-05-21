//ff:func feature=agent type=command control=iteration dimension=1
//ff:what Run — agent 메인 흐름 (scaffold → v2 validate loop)

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/park-jun-woo/yongol/pkg/chain"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/validate"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"gopkg.in/yaml.v3"
)

// Config holds CLI flags for the agent command.
type Config struct {
	SpecsDir  string
	Backend   string // "ollama", "xai", "gemini"
	Model     string // model name within the backend
	MaxRounds int
}

// Run executes the agent: scaffold SSOT files, then run the v2 validate loop.
func Run(w io.Writer, cfg Config) error {
	start := time.Now()

	// Scaffold: generate SSOT files from features.yaml before validation
	ff, _ := features.Load(cfg.SpecsDir)
	if ff != nil {
		if err := scaffold(cfg.SpecsDir, ff, llmCall, cfg, w); err != nil {
			return fmt.Errorf("scaffold: %w", err)
		}
	}
	if cfg.MaxRounds == 0 {
		fmt.Fprintf(w, "\nyongol agent: scaffold only (max-rounds=0), %.1fs\n", time.Since(start).Seconds())
		return nil
	}

	// v2 validate loop: validate → filter immutable → fix per file → repeat.
	return validateLoop(cfg.SpecsDir, ff, llmCall, cfg, w, os.Stderr, cfg.MaxRounds)
}

// runValidate runs DetectSSOTs → ParseAll → Validate and returns all diagnostics.
func runValidate(specsDir string) ([]diagnostic.Diagnostic, error) {
	detected, err := yongol.DetectSSOTs(specsDir)
	if err != nil {
		return nil, fmt.Errorf("detect SSOTs: %w", err)
	}
	fs := yongol.ParseAll(specsDir, detected)
	if len(fs.ParseDiagnostics) > 0 {
		return fs.ParseDiagnostics, nil
	}
	report := validate.Validate(fs)
	var all []diagnostic.Diagnostic
	for _, step := range report.Steps {
		all = append(all, step.Diagnostics...)
	}
	return all, nil
}

// rebaseFile converts an absolute file path to a specsDir-relative path.
// If already relative or conversion fails, returns the original.
func rebaseFile(file, absSpecs string) string {
	if !filepath.IsAbs(file) {
		return file
	}
	rel, err := filepath.Rel(absSpecs, file)
	if err != nil {
		return file
	}
	return filepath.ToSlash(rel)
}

// fileGroup groups diagnostics by file path.
type fileGroup struct {
	relFile string // specs-dir relative path
	layer   layer
	diags   []diagnostic.Diagnostic
}

// groupByFile groups diagnostics by their File field, converting to specs-dir relative paths.
func groupByFile(diags []diagnostic.Diagnostic, absSpecs string) []fileGroup {
	m := map[string]*fileGroup{}
	var order []string
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			continue
		}
		rel := rebaseFile(d.File, absSpecs)
		if _, ok := m[rel]; !ok {
			m[rel] = &fileGroup{relFile: rel, layer: classifyFile(rel)}
			order = append(order, rel)
		}
		m[rel].diags = append(m[rel].diags, d)
	}
	result := make([]fileGroup, 0, len(order))
	for _, f := range order {
		result = append(result, *m[f])
	}
	return result
}

// sortByLayerPriority sorts file groups by the layer priority order.
func sortByLayerPriority(groups []fileGroup) {
	priority := make(map[layer]int)
	for i, l := range layerPriority {
		priority[l] = i
	}
	sort.SliceStable(groups, func(i, j int) bool {
		pi, ok := priority[groups[i].layer]
		if !ok {
			pi = 999
		}
		pj, ok := priority[groups[j].layer]
		if !ok {
			pj = 999
		}
		return pi < pj
	})
}

// countErrors counts ERROR-level diagnostics.
func countErrors(diags []diagnostic.Diagnostic) int {
	n := 0
	for _, d := range diags {
		if d.Level == diagnostic.LevelError {
			n++
		}
	}
	return n
}

// diagMessages extracts message strings from diagnostics.
func diagMessages(diags []diagnostic.Diagnostic) []string {
	msgs := make([]string, len(diags))
	for i, d := range diags {
		msgs[i] = d.Message
	}
	return msgs
}

// loadFeatureLookup builds a map from operationId to Feature.
func loadFeatureLookup(specsDir string) map[string]features.Feature {
	path := filepath.Join(specsDir, "features.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var ff features.FeaturesFile
	if err := yaml.Unmarshal(data, &ff); err != nil {
		return nil
	}
	m := make(map[string]features.Feature, len(ff.Features))
	for _, f := range ff.Features {
		m[f.Op] = f
	}
	return m
}

// resolveFeature extracts operationId and feature desc for a file.
// For SSaC files, the operationId is the filename stem.
func resolveFeature(relPath string, l layer, lookup map[string]features.Feature) (desc, path string) {
	if l == layerSSaC {
		base := filepath.Base(relPath)
		op := strings.TrimSuffix(base, ".ssac")
		if f, ok := lookup[op]; ok {
			return f.Desc, f.Path
		}
		// SSaC without feature desc — still fixable
		return op + " (no desc)", ""
	}
	return "", ""
}

// chainResolve tries to use the chain package to find an operationId for a non-SSaC file.
func chainResolve(specsDir, relPath string, lookup map[string]features.Feature) (desc, path string) {
	// Parse the fullstack to use chain
	detected, err := yongol.DetectSSOTs(specsDir)
	if err != nil {
		return "", ""
	}
	fs := yongol.ParseAll(specsDir, detected)
	if len(fs.ParseDiagnostics) > 0 {
		return "", ""
	}

	// Try each operationId from features to see if chain links to this file
	for op, feat := range lookup {
		links, err := chain.Chain(fs, op)
		if err != nil {
			continue
		}
		for _, link := range links {
			if link.File == relPath {
				return feat.Desc, feat.Path
			}
		}
	}
	return "", ""
}

// reOperationID matches operationId references in diagnostic messages.
// Patterns: "operationId X", "operationId: X", "operationId \"X\"",
// "operationId 'X'", "func X", "action == \"X\"", "# X" comment refs,
// "SSaC authorize (X, ...)", "Missing: X, Y, ...".
var reOperationID = regexp.MustCompile(
	`(?:operationId[:\s]+\"?([A-Z][A-Za-z0-9]+)\"?` +
		`|SSaC func ([A-Z][A-Za-z0-9]+)` +
		`|SSaC authorize \(([A-Z][A-Za-z0-9]+)` +
		`|input\.action == "([A-Z][A-Za-z0-9]+)"` +
		`|# ([A-Z][A-Za-z0-9]+))`,
)

// reMissingList matches "Missing: X, Y, Z" lists from XOH-11 style diagnostics.
var reMissingList = regexp.MustCompile(`Missing:\s*([A-Z][A-Za-z0-9]+(?:\s*,\s*[A-Z][A-Za-z0-9]+)*)`)

// reMissingItem extracts individual operationIds from a comma-separated list.
var reMissingItem = regexp.MustCompile(`[A-Z][A-Za-z0-9]+`)

// extractOperationIDs extracts unique operationIds from diagnostic messages.
func extractOperationIDs(diags []diagnostic.Diagnostic) []string {
	seen := map[string]struct{}{}
	var result []string
	addUnique := func(id string) {
		if _, ok := seen[id]; !ok {
			seen[id] = struct{}{}
			result = append(result, id)
		}
	}
	for _, d := range diags {
		matches := reOperationID.FindAllStringSubmatch(d.Message, -1)
		for _, m := range matches {
			for _, g := range m[1:] {
				if g != "" {
					addUnique(g)
				}
			}
		}
		// Handle "Missing: X, Y, Z" lists (XOH-11 style)
		if ml := reMissingList.FindStringSubmatch(d.Message); len(ml) > 1 {
			items := reMissingItem.FindAllString(ml[1], -1)
			for _, item := range items {
				addUnique(item)
			}
		}
	}
	return result
}

// filterMessagesByOp returns diagnostic messages that mention the given operationId.
func filterMessagesByOp(messages []string, opID string) []string {
	var filtered []string
	for _, m := range messages {
		if strings.Contains(m, opID) {
			filtered = append(filtered, m)
		}
	}
	return filtered
}

// fixSingleBlock extracts a block for an operationId, sends it to LLM, and merges
// the result back into fileContent. Returns true if the block was fixed.
// If the block does not exist, attempts to generate new content using SSaC + features.
func fixSingleBlock(w io.Writer, cfg Config, l layer, relFile, absPath string, fileContent *string, opID, desc, path string, diagMsgs []string) bool {
	var block string
	var startLine, endLine int
	var extractErr error

	switch l {
	case layerOpenAPI:
		block, startLine, endLine, extractErr = extractOpenAPIBlock(*fileContent, opID)
	case layerRego:
		block, startLine, endLine, extractErr = extractRegoBlock(*fileContent, opID)
	case layerHurl:
		block, startLine, endLine, extractErr = extractHurlBlock(*fileContent, opID)
	default:
		return false
	}

	// Block not found — try generating new content
	if extractErr != nil {
		return generateNewBlock(w, cfg, l, relFile, absPath, fileContent, opID, desc, path)
	}

	// Block found — fix existing content (original flow)
	systemPrompt := buildSystemPrompt(l, diagMsgs)
	userPrompt := buildBlockUserPrompt(desc, path, filepath.Base(relFile), opID, block, diagMsgs)

	reply, err := llmCall(cfg.Backend, cfg.Model, systemPrompt, userPrompt)
	if err != nil {
		fmt.Fprintf(w, "  skipped block: %s/%s (LLM error: %v)\n", relFile, opID, err)
		return false
	}

	fixedBlock := stripMarkdownFences(reply)
	if fixedBlock == "" {
		fmt.Fprintf(w, "  skipped block: %s/%s (empty LLM response)\n", relFile, opID)
		return false
	}

	// Merge the fixed block back
	var merged string
	switch l {
	case layerOpenAPI:
		merged, err = mergeOpenAPIBlock(*fileContent, startLine, endLine, fixedBlock)
	case layerRego:
		merged, err = mergeRegoBlock(*fileContent, startLine, endLine, fixedBlock)
	case layerHurl:
		merged, err = mergeHurlBlock(*fileContent, startLine, endLine, fixedBlock)
	}

	if err != nil {
		fmt.Fprintf(w, "  skipped block: %s/%s (merge error: %v)\n", relFile, opID, err)
		return false
	}

	*fileContent = merged
	return true
}

// generateNewBlock generates a new block for an operationId using SSaC content + features,
// then inserts it into the file. Returns true if block was generated and inserted.
func generateNewBlock(w io.Writer, cfg Config, l layer, relFile, absPath string, fileContent *string, opID, desc, featurePath string) bool {
	// Find SSaC file: search service/*/{opID}.ssac
	specsDir := filepath.Dir(filepath.Dir(absPath))
	if l == layerRego {
		// For policy/*.rego, specs dir is one level up from policy/
		specsDir = filepath.Dir(filepath.Dir(absPath))
	}
	// Normalize: for api/openapi.yaml → specsDir is parent of api/
	// For policy/authz.rego → specsDir is parent of policy/
	// For tests/api.hurl → specsDir is parent of tests/
	// In all cases we want the specs root containing service/
	specsDir = resolveSpecsRoot(absPath, l)

	ssacContent, ssacFound := findSSaCFile(specsDir, opID)
	if !ssacFound {
		fmt.Fprintf(w, "  skipped generate: %s/%s (SSaC file not found)\n", relFile, opID)
		return false
	}

	// Build generate prompt
	systemPrompt := "You generate yongol SSOT content. Output ONLY the requested block. No explanations. No markdown fences.\n\nExample for " + layerName(l) + ":\n" + layerExample(l)
	userPrompt := buildGeneratePrompt(l, opID, desc, featurePath, ssacContent)

	reply, err := llmCall(cfg.Backend, cfg.Model, systemPrompt, userPrompt)
	if err != nil {
		fmt.Fprintf(w, "  skipped generate: %s/%s (LLM error: %v)\n", relFile, opID, err)
		return false
	}

	newBlock := stripMarkdownFences(reply)
	if newBlock == "" {
		fmt.Fprintf(w, "  skipped generate: %s/%s (empty LLM response)\n", relFile, opID)
		return false
	}

	// Insert new block into file content
	var inserted string
	switch l {
	case layerOpenAPI:
		inserted, err = insertOpenAPIBlock(*fileContent, newBlock)
	case layerRego:
		inserted, err = insertRegoBlock(*fileContent, newBlock)
	case layerHurl:
		inserted, err = insertHurlBlock(*fileContent, newBlock)
	default:
		return false
	}

	if err != nil {
		fmt.Fprintf(w, "  skipped generate: %s/%s (insert error: %v)\n", relFile, opID, err)
		return false
	}

	*fileContent = inserted
	fmt.Fprintf(w, "  generated: %s/%s\n", relFile, opID)
	return true
}

// resolveSpecsRoot derives the specs root directory from an absolute file path and its layer.
func resolveSpecsRoot(absPath string, l layer) string {
	switch l {
	case layerOpenAPI:
		// absPath = .../specs/api/openapi.yaml → specs root = parent of api/
		return filepath.Dir(filepath.Dir(absPath))
	case layerRego:
		// absPath = .../specs/policy/authz.rego → specs root = parent of policy/
		return filepath.Dir(filepath.Dir(absPath))
	case layerHurl:
		// absPath = .../specs/tests/api.hurl → specs root = parent of tests/
		return filepath.Dir(filepath.Dir(absPath))
	default:
		return filepath.Dir(absPath)
	}
}

// findSSaCFile searches for service/*/{operationId}.ssac under specsDir.
// Returns the file content and true if found.
func findSSaCFile(specsDir, operationId string) (string, bool) {
	serviceDir := filepath.Join(specsDir, "service")

	// Glob for service/*/{operationId}.ssac
	pattern := filepath.Join(serviceDir, "*", operationId+".ssac")
	matches, err := filepath.Glob(pattern)
	if err != nil || len(matches) == 0 {
		return "", false
	}

	data, err := os.ReadFile(matches[0])
	if err != nil {
		return "", false
	}
	return string(data), true
}

// buildBlockUserPrompt assembles the user prompt for a single block fix.
func buildBlockUserPrompt(desc, path, filename, opID, block string, messages []string) string {
	var b strings.Builder
	if desc != "" {
		fmt.Fprintf(&b, "Feature: %s\nPath: %s\n\n", desc, path)
	}
	fmt.Fprintf(&b, "OperationId: %s\nFile: %s\n\nCurrent block:\n%s\n\nValidate errors:\n", opID, filename, block)
	for _, m := range messages {
		b.WriteString(m)
		b.WriteByte('\n')
	}
	b.WriteString("\nFix ONLY this block. Output ONLY the corrected block content. Do not add surrounding file content.")
	return b.String()
}
