//ff:func feature=agent type=command control=iteration dimension=1
//ff:what scaffoldStateMachine — states가 있는 테이블로부터 Mermaid stateDiagram 자동 생성

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

// scaffoldStateMachine generates specs/states/{table}.md for tables with states.
// Each table with states produces one LLM call. Existing files are skipped.
func scaffoldStateMachine(specsDir string, ff *features.FeaturesFile, llmFn LLMCallFunc, cfg Config, out io.Writer) (int, error) {
	// Collect tables that have states
	type tableWithStates struct {
		name   string
		states []string
	}
	var targets []tableWithStates
	sorted := topoSortTables(ff.Tables)
	for _, name := range sorted {
		td := ff.Tables[name]
		if len(td.States) > 0 {
			targets = append(targets, tableWithStates{name: name, states: td.States})
		}
	}

	if len(targets) == 0 {
		return 0, nil
	}

	statesDir := filepath.Join(specsDir, "states")
	if err := os.MkdirAll(statesDir, 0755); err != nil {
		return 0, fmt.Errorf("create states dir: %w", err)
	}

	// Load states docs for system prompt; fallback to hardcoded if not found
	systemPrompt := defaultStateDiagramSystem()
	if systemDoc, err := docs.FS.ReadFile("states.md"); err == nil {
		systemPrompt = string(systemDoc)
	}

	tableFeatMap := buildTableFeatureMap(ff)
	count := 0

	for _, target := range targets {
		outPath := filepath.Join(statesDir, target.name+".md")
		if _, err := os.Stat(outPath); err == nil {
			fmt.Fprintf(out, "  scaffold states: skipped %s.md (exists)\n", target.name)
			continue
		}

		feats := tableFeatMap[target.name]
		userPrompt := buildStateMachineUserPrompt(target.name, target.states, feats)

		numCtx := int(float64(len(systemPrompt)+len(userPrompt))/4*1.5) + 2048

		reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, systemPrompt, userPrompt, numCtx)
		if err != nil {
			return 0, fmt.Errorf("scaffold states %s: %w", target.name, err)
		}

		content := strings.TrimSpace(reply)
		if content == "" {
			return 0, fmt.Errorf("scaffold states %s: empty LLM response", target.name)
		}

		if err := os.WriteFile(outPath, []byte(content+"\n"), 0644); err != nil {
			return 0, fmt.Errorf("scaffold states %s: write: %w", target.name, err)
		}

		fmt.Fprintf(out, "  scaffold states: created %s.md\n", target.name)
		count++
	}

	return count, nil
}

// buildStateMachineUserPrompt builds the user prompt for generating a state diagram.
func buildStateMachineUserPrompt(tableName string, states []string, feats []features.Feature) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Table: %s\n", tableName)
	fmt.Fprintf(&b, "States: %s\n", strings.Join(states, ", "))

	if len(feats) > 0 {
		b.WriteString("\nRelated features:\n")
		for _, f := range feats {
			fmt.Fprintf(&b, "  - %s %s: %s\n", f.Op, f.Path, f.Desc)
		}
	}

	b.WriteString("\nGenerate a Mermaid stateDiagram-v2 for this table.")
	b.WriteString("\nOutput a markdown file with '# <table_name>' heading followed by a mermaid fenced code block.")
	b.WriteString("\nInclude [*] --> first_state as the initial transition.")
	b.WriteString("\nUse operationId names for transition labels where applicable.")

	return b.String()
}

// defaultStateDiagramSystem returns a fallback system prompt for state diagram generation.
func defaultStateDiagramSystem() string {
	return `You generate Mermaid stateDiagram-v2 diagrams for yongol SSOT specs.

Rules:
- Use stateDiagram-v2 syntax
- Always start with [*] --> first_state
- Label transitions with operationId names (e.g. ActivateWorkflow)
- Each state must be one of the defined states from the table
- Output a markdown file with a heading and a mermaid fenced code block

Example:
# workflow

` + "```mermaid\nstateDiagram-v2\n    [*] --> draft\n    draft --> active : ActivateWorkflow\n    active --> paused : PauseWorkflow\n    paused --> active : ResumeWorkflow\n```"
}
