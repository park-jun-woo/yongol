//ff:func feature=agent type=command control=iteration dimension=1
//ff:what Run — agent 메인 루프 (validate → group → fix → repeat)

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
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

// Run executes the agent loop: validate → group diagnostics → fix files → repeat.
func Run(w io.Writer, cfg Config) error {
	start := time.Now()
	totalFixed := 0

	absSpecs, err := filepath.Abs(cfg.SpecsDir)
	if err != nil {
		return fmt.Errorf("resolve specs dir: %w", err)
	}

	// Build feature lookup (op → Feature)
	featureLookup := loadFeatureLookup(cfg.SpecsDir)

	for round := 1; round <= cfg.MaxRounds; round++ {
		// 1. Detect + Parse + Validate
		diags, err := runValidate(cfg.SpecsDir)
		if err != nil {
			return fmt.Errorf("round %d: %w", round, err)
		}

		errorCount := countErrors(diags)
		fmt.Fprintf(w, "\nyongol agent: round %d — %d errors\n", round, errorCount)

		if errorCount == 0 {
			fmt.Fprintf(w, "  %d files fixed in %d rounds (%.1fs)\n", totalFixed, round, time.Since(start).Seconds())
			return nil
		}

		// 2. Group diagnostics by file, ordered by layer priority
		groups := groupByFile(diags, absSpecs)
		sortByLayerPriority(groups)

		roundFixed := 0
		for _, g := range groups {
			l := classifyFile(g.relFile)

			// Extract operationId and feature desc
			desc, path := resolveFeature(g.relFile, l, featureLookup)

			// For non-SSaC layers without chain resolution, skip
			if l != layerSSaC && desc == "" {
				// Try chain-based resolution
				desc, path = chainResolve(cfg.SpecsDir, g.relFile, featureLookup)
				if desc == "" {
					fmt.Fprintf(w, "  skipped: %s (chain unavailable)\n", g.relFile)
					continue
				}
			}

			// Read current file content
			absPath := filepath.Join(absSpecs, g.relFile)
			content, err := os.ReadFile(absPath)
			if err != nil {
				fmt.Fprintf(w, "  skipped: %s (read error: %v)\n", g.relFile, err)
				continue
			}

			// Build prompts
			messages := diagMessages(g.diags)
			systemPrompt := buildSystemPrompt(l, messages)
			userPrompt := buildUserPrompt(desc, path, filepath.Base(g.relFile), string(content), messages)

			// Call LLM
			reply, err := llmCall(cfg.Backend, cfg.Model, systemPrompt, userPrompt)
			if err != nil {
				fmt.Fprintf(w, "  skipped: %s (LLM error: %v)\n", g.relFile, err)
				continue
			}

			// Strip markdown fences
			fixed := stripMarkdownFences(reply)
			if fixed == "" {
				fmt.Fprintf(w, "  skipped: %s (empty LLM response)\n", g.relFile)
				continue
			}

			// Write fixed content back
			if err := os.WriteFile(absPath, []byte(fixed), 0644); err != nil {
				fmt.Fprintf(w, "  skipped: %s (write error: %v)\n", g.relFile, err)
				continue
			}

			hint := g.relFile
			if len(g.diags) > 0 {
				msg := g.diags[0].Message
				if len(msg) > 60 {
					msg = msg[:60] + "..."
				}
				hint = g.relFile + " — " + msg
			}
			fmt.Fprintf(w, "  fixed: %s\n", hint)
			roundFixed++
			totalFixed++
		}

		if roundFixed == 0 {
			fmt.Fprintf(w, "  no files fixed this round — stopping early\n")
			break
		}
	}

	// Final validation to report remaining errors
	diags, _ := runValidate(cfg.SpecsDir)
	remaining := countErrors(diags)
	if remaining > 0 {
		fmt.Fprintf(w, "\nyongol agent: %d errors remaining after %d rounds (%.1fs)\n", remaining, cfg.MaxRounds, time.Since(start).Seconds())
		for _, d := range diags {
			if d.Level == diagnostic.LevelError {
				rel := rebaseFile(d.File, absSpecs)
				fmt.Fprintf(w, "  %s:%d %s\n", rel, d.Line, d.Message)
			}
		}
		return fmt.Errorf("%d errors remaining", remaining)
	}
	fmt.Fprintf(w, "\nyongol agent: 0 errors — %d files fixed (%.1fs)\n", totalFixed, time.Since(start).Seconds())
	return nil
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
