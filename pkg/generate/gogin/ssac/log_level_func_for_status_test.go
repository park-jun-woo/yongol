//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what logLevelFuncForStatus 단위 테스트 (4xx→slog.Warn, 그 외→slog.Error)

package ssac

import "testing"

func TestLogLevelFuncForStatus(t *testing.T) {
	cases := map[int]string{
		400: "slog.Warn",
		404: "slog.Warn",
		499: "slog.Warn",
		500: "slog.Error",
		200: "slog.Error",
		502: "slog.Error",
	}
	for in, want := range cases {
		if got := logLevelFuncForStatus(in); got != want {
			t.Errorf("logLevelFuncForStatus(%d) = %q, want %q", in, got, want)
		}
	}
}
