//ff:func feature=orchestrator type=parser control=sequence
//ff:what appendBodyLine — QuerySpec.Body 에 한 줄 누적 (개행 분리)
package sqlc

// appendBodyLine appends a single body line to the current QuerySpec, inserting
// a newline separator between lines. Keeping comments intact preserves source
// fidelity for downstream consumers (e.g. RETURNING extraction in XQS-20).
func appendBodyLine(current *QuerySpec, line string) {
	if current.Body == "" {
		current.Body = line
		return
	}
	current.Body = current.Body + "\n" + line
}
