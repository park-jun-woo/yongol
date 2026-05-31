//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestElementHeadMeta — elementHeadMeta 어댑터 CheckEnum/RawType 검증
package types

import (
	"testing"
)

func TestElementHeadMetaRawType(t *testing.T) {
	m := elementHeadMeta{head: "BIGINT"}
	if m.RawType() != "BIGINT" {
		t.Errorf("RawType() = %q, want BIGINT", m.RawType())
	}
}
