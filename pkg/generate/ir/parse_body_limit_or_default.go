//ff:func feature=gen-ir type=util control=selection
//ff:what parseBodyLimitOrDefault -- 사람이 읽을 수 있는 크기 문자열 파싱 (MiB/KiB/GiB/MB/KB)

package ir

// parseBodyLimitOrDefault parses a human-readable size string and returns
// the default on failure. The actual parsing logic is in gogin/middleware;
// for the IR we keep a simplified version that handles common suffixes.
func parseBodyLimitOrDefault(s string, def int64) int64 {
	if s == "" {
		return def
	}
	var multiplier int64
	numStr := s
	switch {
	case len(s) > 3 && s[len(s)-3:] == "MiB":
		multiplier = 1024 * 1024
		numStr = s[:len(s)-3]
	case len(s) > 3 && s[len(s)-3:] == "KiB":
		multiplier = 1024
		numStr = s[:len(s)-3]
	case len(s) > 3 && s[len(s)-3:] == "GiB":
		multiplier = 1024 * 1024 * 1024
		numStr = s[:len(s)-3]
	case len(s) > 2 && s[len(s)-2:] == "MB":
		multiplier = 1000 * 1000
		numStr = s[:len(s)-2]
	case len(s) > 2 && s[len(s)-2:] == "KB":
		multiplier = 1000
		numStr = s[:len(s)-2]
	default:
		return def
	}
	n := parseDigits(numStr)
	if n <= 0 {
		return def
	}
	return n * multiplier
}
