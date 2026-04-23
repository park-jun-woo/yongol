//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-observability
//ff:what OBS-003 테스트 — 허용된 exporter 값(otlp/stdout/noop/"") 은 조용히 통과

package manifest

import "testing"

func TestObs03TracingExporter_Allowed(t *testing.T) {
	cases := []string{"otlp", "stdout", "noop", ""}
	for _, v := range cases {
		t.Run("exporter="+v, func(t *testing.T) {
			assertObs03ExporterAllowed(t, v)
		})
	}
}
