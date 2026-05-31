//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestElementHeadMetaRawType — elementHeadMeta.RawType head 반환 커버
package types

import "testing"

func TestElementHeadMetaRawType_ZeroCov(t *testing.T) {
	m := elementHeadMeta{head: "BIGINT"}
	if m.RawType() != "BIGINT" {
		t.Errorf("RawType = %q, want BIGINT", m.RawType())
	}
}
