//ff:func feature=gen-gogin type=generator control=sequence
//ff:what loggerHelperParseLogLevel — LOG_LEVEL 문자열 → slog.Level 변환 헬퍼 소스 반환

package boot

// loggerHelperParseLogLevel returns the top-level parseLogLevel(level string) slog.Level
// helper source. Extracted from main() so that main no longer carries a depth-1 switch
// (filefunc A13: selection + loop mixed at depth 1 is forbidden). The helper is emitted
// as a sibling 1-file-1-func file under cmd/ via writeEnvHelperFiles.
func loggerHelperParseLogLevel() string {
	return `func parseLogLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return slog.LevelDebug
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	}
	return slog.LevelInfo
}`
}
