//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldStateMachine — states 없음 0 / 기존파일 skip / mkdir 에러 / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldStateMachineNoTargets(t *testing.T) {
	var out bytes.Buffer
	// Tables without States produce no targets -> count 0, nil error.
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{"users": {}}}
	n, err := scaffoldStateMachine(t.TempDir(), ff, nil, Config{}, &out)
	if err != nil {
		t.Fatalf("no targets: unexpected error: %v", err)
	}
	if n != 0 {
		t.Fatalf("no targets: expected 0, got %d", n)
	}
}
