package boot

// envHelperImports — the imports the env-helper funcs compile against.
var envHelperImports = []string{
	`"fmt"`,
	`"os"`,
	`"strconv"`,
	`"strings"`,
	`"time"`,
}

// envHelperFuncs — rendered verbatim into the generated main.go after main().
// Each entry is a complete function declaration. Order is not significant
// but kept alphabetical (mostly) for readability.
var envHelperFuncs = []string{
	envIntFunc,
	envInt64Func,
	parseSizeFunc,
	envDurationFunc,
	envStringListFunc,
	envBoolFunc,
	envStringFunc,
	envFloatFunc,
}

var envIntFunc = `func envInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}`

var envInt64Func = `func envInt64(key string, def int64) int64 {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := parseSize(v)
	if err != nil {
		return def
	}
	return n
}`

var parseSizeFunc = `func parseSize(s string) (int64, error) {
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
}`

var envDurationFunc = `func envDuration(key string, def time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return def
}`

var envStringListFunc = `func envStringList(key string, def []string) []string {
	if v := os.Getenv(key); v != "" {
		parts := strings.Split(v, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		return parts
	}
	return def
}`

var envBoolFunc = `func envBool(key string, def bool) bool {
	if v := os.Getenv(key); v != "" {
		return strings.EqualFold(v, "true") || v == "1"
	}
	return def
}`

var envStringFunc = `func envString(key string, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}`

var envFloatFunc = `func envFloat(key string, def float64) float64 {
	if v := os.Getenv(key); v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return def
}`
