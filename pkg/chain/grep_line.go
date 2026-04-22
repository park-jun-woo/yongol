//ff:func feature=chain type=util control=iteration dimension=1
//ff:what grepLine returns the first line number containing substr, or 0 if not found.
package chain

import (
	"bufio"
	"os"
	"strings"
)

// grepLine returns the first line number (1-based) containing substr, or 0 if not found.
func grepLine(filePath string, substr string) int {
	f, err := os.Open(filePath)
	if err != nil {
		return 0
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		if strings.Contains(scanner.Text(), substr) {
			return lineNum
		}
	}
	return 0
}
