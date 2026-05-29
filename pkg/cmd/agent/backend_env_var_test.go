//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestBackendEnvVar — backend 이름별 환경변수 매핑/기본값 검증

package agent

import "testing"

func TestBackendEnvVar(t *testing.T) {
	cases := []struct {
		backend string
		want    string
	}{
		{"xai", "XAI_API_KEY"},
		{"gemini", "GEMINI_API_KEY"},
		{"openai", ""},
		{"", ""},
	}
	for _, tc := range cases {
		t.Run(tc.backend, func(t *testing.T) {
			if got := backendEnvVar(tc.backend); got != tc.want {
				t.Errorf("backendEnvVar(%q) = %q, want %q", tc.backend, got, tc.want)
			}
		})
	}
}
