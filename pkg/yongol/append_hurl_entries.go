//ff:func feature=orchestrator type=loader control=iteration dimension=1
//ff:what fs.HurlFiles 전체를 순회하며 엔트리를 누적하고 파서 diag 를 ParseDiagnostics 에 전파
package yongol

import (
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// appendHurlEntries parses every file in fs.HurlFiles, appends the resulting
// entries to fs.HurlEntries and forwards each file's parse diagnostics into
// fs.ParseDiagnostics so the CLI-level gate (len(fs.ParseDiagnostics) > 0)
// can trip on malformed hurl input.
func appendHurlEntries(fs *Fullstack) {
	for _, hf := range fs.HurlFiles {
		entries, diags := hurl.ParseFile(hf)
		fs.HurlEntries = append(fs.HurlEntries, entries...)
		fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	}
}
