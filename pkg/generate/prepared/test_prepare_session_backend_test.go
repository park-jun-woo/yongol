//ff:func feature=generate type=test control=iteration dimension=1
//ff:what TestPrepareSessionBackend — manifest × SSaC 2x2 매트릭스 활성 판정

package prepared

import "testing"

func TestPrepareSessionBackend(t *testing.T) {
	for _, tc := range prepareSessionBackendCases() {
		t.Run(tc.name, runPrepareSessionBackendCase(tc))
	}
}
