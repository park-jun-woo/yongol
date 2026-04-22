//ff:func feature=gen-ffhash type=util control=iteration dimension=1
//ff:what InjectCheckedLine — //ff: 블록 끝에 //ff:checked llm=yongol-gen hash=<sha> 라인 삽입/갱신

package ffhash

import (
	"bytes"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

// LLMName is the namespace recorded in the `llm=` field of the injected
// `//ff:checked` annotation. It intentionally differs from the real
// ollama model identifiers used by `filefunc llmc` so that operators
// can distinguish machine-generated artifacts from human-verified ones.
const LLMName = "yongol-gen"

// InjectCheckedLine returns src with a `//ff:checked llm=yongol-gen
// hash=<sha>` line present in the top annotation block. Behaviour:
//
//   - if src has no `//ff:` annotation line (no annotation block), the
//     input is returned unchanged; callers generate files with an
//     annotation block, so this path is only reached for foreign files
//     that should not be rewritten
//   - if a `//ff:checked` line already exists, the hash value is
//     updated in place
//   - otherwise a new `//ff:checked` line is inserted immediately after
//     the last `//ff:` annotation line (and before the next non-`//ff:`
//     line such as the package declaration)
//
// When ComputeBodyHash returns "" (no func decl found in the source),
// InjectCheckedLine returns src untouched — a hash line with an empty
// value would be misleading and filefunc A7 skips func-less files
// anyway, so omission is the correct behaviour.
//
// The input is first normalised (CRLF→LF, BOM removed, trailing LF
// ensured) so the output is deterministic regardless of how callers
// build the source buffer.
func InjectCheckedLine(src []byte) []byte {
	normalised := contract.NormalizeBody(src)
	hash := contract.ComputeBodyHash(normalised)
	if hash == "" {
		return normalised
	}
	checkedLine := "//ff:checked llm=" + LLMName + " hash=" + hash
	lines := bytes.Split(normalised, []byte("\n"))
	lastFFIdx := -1
	checkedIdx := -1
	for i, line := range lines {
		trim := strings.TrimSpace(string(line))
		if !strings.HasPrefix(trim, "//ff:") {
			continue
		}
		lastFFIdx = i
		if strings.HasPrefix(trim, "//ff:checked") {
			checkedIdx = i
		}
	}
	if lastFFIdx < 0 {
		return normalised
	}
	if checkedIdx >= 0 {
		lines[checkedIdx] = []byte(checkedLine)
		return bytes.Join(lines, []byte("\n"))
	}
	out := make([][]byte, 0, len(lines)+1)
	out = append(out, lines[:lastFFIdx+1]...)
	out = append(out, []byte(checkedLine))
	out = append(out, lines[lastFFIdx+1:]...)
	return bytes.Join(out, []byte("\n"))
}
