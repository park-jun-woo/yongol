//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestActivePorts — when 이 truthy 한 port 만 남기는지 검증
package infra

import (
	"testing"
)

func TestActivePortsEmpty(t *testing.T) {
	if got := activePorts(nil, map[string]any{}); got != nil {
		t.Errorf("expected nil for no ports, got: %v", got)
	}
}
