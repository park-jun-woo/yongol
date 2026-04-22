//ff:type feature=contract type=model
//ff:what checkedLineRe — //ff:checked hash=<hex> 주석에서 hash 값을 캡처하는 정규식

package contract

import "regexp"

// checkedLineRe captures the hash value on a `//ff:checked llm=...
// hash=<hex>` annotation. The llm namespace is intentionally
// open-ended so that files checked by filefunc llmc also parse
// cleanly — callers that want to restrict to `llm=yongol-gen`
// inspect the llm field separately.
var checkedLineRe = regexp.MustCompile(`(?m)^\s*//ff:checked\b[^\n]*\bhash=([0-9a-fA-F]+)`)
