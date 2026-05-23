//ff:func feature=agent type=helper control=sequence
//ff:what defaultStateDiagramSystem — state diagram 생성 fallback system prompt

package agent

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
