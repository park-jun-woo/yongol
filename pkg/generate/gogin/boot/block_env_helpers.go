//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockEnvHelpers — envInt / envDuration / envStringList / envBool 헬퍼 (top-level)

package boot

// blockEnvHelpers produces top-level function declarations (envInt,
// envDuration, envStringList, envBool) appended after main(). Used by
// DB pool config, CORS config, and any future env-driven block. The
// helpers silently fall back to the provided default on parse failure.
//
// Lines is empty — Funcs slot carries the entire payload. Imports include
// strconv / time / strings / os because main() itself doesn't need them
// just for helper usage; dedup merges with other blocks that already
// import os etc.
func blockEnvHelpers() MainBlock {
	return MainBlock{
		Name: "env-helpers",
		Imports: []string{
			`"fmt"`,
			`"os"`,
			`"strconv"`,
			`"strings"`,
			`"time"`,
		},
		Funcs: []string{
			`func envInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}`,
			`func envInt64(key string, def int64) int64 {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := parseSize(v)
	if err != nil {
		return def
	}
	return n
}`,
			`func parseSize(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty size")
	}
	up := strings.ToUpper(s)
	var mult int64 = 1
	for _, suf := range []struct {
		tag string
		v   int64
	}{
		{"KIB", 1 << 10}, {"MIB", 1 << 20}, {"GIB", 1 << 30}, {"TIB", 1 << 40},
		{"KB", 1000}, {"MB", 1000 * 1000}, {"GB", 1000 * 1000 * 1000}, {"TB", 1000 * 1000 * 1000 * 1000},
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
}`,
			`func envDuration(key string, def time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return def
}`,
			`func envStringList(key string, def []string) []string {
	if v := os.Getenv(key); v != "" {
		parts := strings.Split(v, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		return parts
	}
	return def
}`,
			`func envBool(key string, def bool) bool {
	if v := os.Getenv(key); v != "" {
		return strings.EqualFold(v, "true") || v == "1"
	}
	return def
}`,
			`func envString(key string, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}`,
			`func envFloat(key string, def float64) float64 {
	if v := os.Getenv(key); v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return def
}`,
		},
	}
}
