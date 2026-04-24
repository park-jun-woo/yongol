//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what ScanSentinelInserts — sentinel INSERT 스캐너의 공개 API (SentinelScanResult slice)
package ddl

// ScanSentinelInserts exposes the parser's sentinel scanner for use by
// other pkg/ modules (e.g. pkg/generate/migration) so the quote-aware
// terminator logic lives in one place.
func ScanSentinelInserts(content string) []SentinelScanResult {
	raw := parseSentinelInserts(content)
	out := make([]SentinelScanResult, 0, len(raw))
	for _, r := range raw {
		out = append(out, SentinelScanResult{
			Table:     r.Table,
			SQL:       r.SQL,
			StartLine: r.StartLine,
			Annotated: r.Annotated,
		})
	}
	return out
}
