//ff:func feature=chain type=util control=iteration dimension=2
//ff:what findDDLTable finds the DDL file and line number for a table name.
package chain

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func findDDLTable(tableName string, specsDir string) (string, int) {
	dbDir := filepath.Join(specsDir, "db")
	entries, err := os.ReadDir(dbDir)
	if err != nil {
		return "db/?.sql", 0
	}
	createRe := regexp.MustCompile(`(?i)CREATE\s+TABLE\s+` + regexp.QuoteMeta(tableName))
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}
		fullPath := filepath.Join(dbDir, entry.Name())
		f, err := os.Open(fullPath)
		if err != nil {
			continue
		}
		scanner := bufio.NewScanner(f)
		lineNum := 0
		for scanner.Scan() {
			lineNum++
			if createRe.MatchString(scanner.Text()) {
				f.Close()
				return "db/" + entry.Name(), lineNum
			}
		}
		f.Close()
	}
	return "db/?.sql", 0
}
