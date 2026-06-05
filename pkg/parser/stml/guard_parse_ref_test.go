//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what parseGuardRef — ref := model "." field 파싱 (정상/모델 누락/점 누락/필드 누락) 검증

package stml

import "testing"

func TestParseGuardRef(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantErr   bool
		wantModel string
		wantField string
	}{
		{name: "valid ref", input: "workflow.status", wantModel: "workflow", wantField: "status"},
		{name: "missing model", input: "=x", wantErr: true},
		{name: "missing dot", input: "workflow=x", wantErr: true},
		{name: "missing field", input: "workflow.=x", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertParseGuardRef(t, tt.input, tt.wantErr, tt.wantModel, tt.wantField)
		})
	}
}
