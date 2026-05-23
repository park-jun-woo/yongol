//ff:func feature=agent type=helper control=selection
//ff:what backendEnvVar — backend 이름에 대응하는 환경변수 반환

package agent

// backendEnvVar returns the environment variable name for a backend.
func backendEnvVar(backend string) string {
	switch backend {
	case "xai":
		return "XAI_API_KEY"
	case "gemini":
		return "GEMINI_API_KEY"
	default:
		return ""
	}
}
