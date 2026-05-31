//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestElementHeadMeta — elementHeadMeta 어댑터 CheckEnum/RawType 검증
package types

import (
	"testing"
)

func TestElementHeadMetaCheckEnum(t *testing.T) {
	m := elementHeadMeta{head: "BIGINT"}
	if m.CheckEnum() != nil {
		t.Errorf("CheckEnum() = %v, want nil", m.CheckEnum())
	}
}
