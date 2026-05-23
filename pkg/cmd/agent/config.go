//ff:type feature=agent type=command
//ff:what Config — agent CLI 플래그 설정 구조체

package agent

// Config holds CLI flags for the agent command.
type Config struct {
	SpecsDir  string
	Backend   string // "ollama", "xai", "gemini"
	Model     string // model name within the backend
	MaxRounds int
}
