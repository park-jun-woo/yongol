//ff:func feature=migration type=util control=iteration dimension=1
//ff:what NextSequenceNumber — artifacts/db/migrations/ 를 스캔해 다음 NNNN 반환
package migration

import (
	"os"
	"strconv"
	"strings"
)

// NextSequenceNumber returns the next 1-based sequence number for a new
// migration file. Scans <dir> for entries matching `NNNN_*.up.sql`;
// returns 1 when the directory is empty or missing. Down stub files
// (`.down.sql`) are skipped so every sequence number is counted once.
func NextSequenceNumber(dir string) (int, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return 1, nil
		}
		return 0, err
	}
	max := 0
	for _, e := range entries {
		name := e.Name()
		if !strings.HasSuffix(name, ".up.sql") {
			continue
		}
		// Up to the first `_`.
		i := strings.Index(name, "_")
		if i <= 0 {
			continue
		}
		n, err := strconv.Atoi(name[:i])
		if err != nil {
			continue
		}
		if n > max {
			max = n
		}
	}
	return max + 1, nil
}
