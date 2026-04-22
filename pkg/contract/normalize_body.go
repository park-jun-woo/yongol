//ff:func feature=contract type=util control=sequence
//ff:what NormalizeBody — CRLF→LF, BOM 제거, 마지막 newline 보장 정규화

package contract

import (
	"bytes"
)

// NormalizeBody normalises a file body for deterministic hashing:
//
//  1. strip a leading UTF-8 BOM (0xEF 0xBB 0xBF) if present
//  2. convert CRLF ("\r\n") to LF ("\n")
//  3. convert bare CR ("\r") to LF ("\n")
//  4. ensure the result ends with a single "\n"
//
// Empty input returns an empty result (no trailing newline forced on
// empty content) so downstream callers can treat "" distinct from "\n".
func NormalizeBody(src []byte) []byte {
	if len(src) == 0 {
		return src
	}
	bom := []byte{0xEF, 0xBB, 0xBF}
	if bytes.HasPrefix(src, bom) {
		src = src[len(bom):]
	}
	src = bytes.ReplaceAll(src, []byte("\r\n"), []byte("\n"))
	src = bytes.ReplaceAll(src, []byte("\r"), []byte("\n"))
	if len(src) == 0 {
		return src
	}
	if src[len(src)-1] != '\n' {
		out := make([]byte, len(src)+1)
		copy(out, src)
		out[len(src)] = '\n'
		return out
	}
	return src
}
