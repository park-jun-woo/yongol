//ff:func feature=agent type=loader control=sequence
//ff:what loadAPIKey — backend별 API 키 로드 (환경변수 → XDG credentials.yaml)

package agent

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// loadAPIKey returns the API key for the given backend.
// Priority: environment variable → $XDG_CONFIG_HOME/yongol/credentials.yaml.
func loadAPIKey(backend string) (string, error) {
	envVar := backendEnvVar(backend)
	if v := os.Getenv(envVar); v != "" {
		return v, nil
	}

	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("load API key for %s: %w", backend, err)
		}
		configDir = filepath.Join(home, ".config")
	}

	credPath := filepath.Join(configDir, "yongol", "credentials.yaml")
	data, err := os.ReadFile(credPath)
	if err != nil {
		return "", fmt.Errorf("load API key for %s: env %s not set and %s: %w", backend, envVar, credPath, err)
	}

	var creds map[string]string
	if err := yaml.Unmarshal(data, &creds); err != nil {
		return "", fmt.Errorf("parse %s: %w", credPath, err)
	}

	key, ok := creds[backend]
	if !ok || key == "" {
		return "", fmt.Errorf("load API key for %s: key not found in %s", backend, credPath)
	}
	return key, nil
}

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
