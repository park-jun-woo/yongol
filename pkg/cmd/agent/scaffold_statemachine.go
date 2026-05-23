//ff:func feature=agent type=command control=iteration dimension=1
//ff:what scaffoldStateMachine — states가 있는 테이블로부터 Mermaid stateDiagram 자동 생성

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

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

	systemPrompt := defaultStateDiagramSystem()
	if systemDoc, err := docs.FS.ReadFile("states.md"); err == nil {
		systemPrompt = string(systemDoc)
	}

	tableFeatMap := buildTableFeatureMap(ff)
	count := 0

	for _, target := range targets {
		created, err := scaffoldStateMachineTarget(statesDir, target.name, target.states, tableFeatMap[target.name], systemPrompt, cfg, out)
		if err != nil {
			return 0, err
		}
		if created {
			count++
		}
	}

	return count, nil
}
