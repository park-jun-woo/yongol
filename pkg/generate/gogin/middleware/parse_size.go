//ff:func feature=gen-gogin type=util control=iteration dimension=1 topic=size-parse
//ff:what ParseSize — "1MiB" / "32MiB" / "500KB" 같은 human-readable size 를 bytes 로 변환

package middleware

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseSize converts a human-readable size literal (e.g. "1MiB", "32MiB",
// "500KB", "1024") to a byte count. Supported suffixes: B, KB, MB, GB, TB,
// KiB, MiB, GiB, TiB (case-insensitive). "0" returns 0 and is interpreted by
// the caller as "no limit" for override semantics.
//
// Implemented manually (no github.com/dustin/go-humanize dep) because yongol
// does not otherwise import that package.
func ParseSize(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty size")
	}
	up := strings.ToUpper(s)
	var mult int64 = 1
	// Binary (KiB/MiB/GiB/TiB) checked first so their plain counterparts do not mis-match.
	for _, suf := range []struct {
		tag string
		v   int64
	}{
		{"KIB", 1 << 10},
		{"MIB", 1 << 20},
		{"GIB", 1 << 30},
		{"TIB", 1 << 40},
		{"KB", 1000},
		{"MB", 1000 * 1000},
		{"GB", 1000 * 1000 * 1000},
		{"TB", 1000 * 1000 * 1000 * 1000},
		{"B", 1},
	} {
		if strings.HasSuffix(up, suf.tag) {
			up = strings.TrimSuffix(up, suf.tag)
			mult = suf.v
			break
		}
	}
	up = strings.TrimSpace(up)
	n, err := strconv.ParseInt(up, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse size %q: %w", s, err)
	}
	if n < 0 {
		return 0, fmt.Errorf("negative size %q", s)
	}
	return n * mult, nil
}
