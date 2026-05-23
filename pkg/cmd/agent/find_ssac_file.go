//ff:func feature=agent type=helper control=sequence
//ff:what findSSaCFile — service/*/{operationId}.ssac 파일 탐색 및 내용 반환

package agent

import (
	"os"
	"path/filepath"
)

// findSSaCFile searches for service/*/{operationId}.ssac under specsDir.
// Returns the file content and true if found.
func findSSaCFile(specsDir, operationId string) (string, bool) {
	serviceDir := filepath.Join(specsDir, "service")

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
