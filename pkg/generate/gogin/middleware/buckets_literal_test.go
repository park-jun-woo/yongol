//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBucketsLiteral — histogram buckets Go 리터럴 포맷 + 기본값 검증

package middleware

import "testing"

func TestBucketsLiteral(t *testing.T) {
	t.Run("EmptyDefault", func(t *testing.T) {
		if got := bucketsLiteral(nil); got != "prometheus.DefBuckets" {
			t.Errorf("nil: got %q, want prometheus.DefBuckets", got)
		}
		if got := bucketsLiteral([]float64{}); got != "prometheus.DefBuckets" {
			t.Errorf("empty: got %q, want prometheus.DefBuckets", got)
		}
	})

	t.Run("NonEmpty", func(t *testing.T) {
		got := bucketsLiteral([]float64{0.005, 0.1, 1})
		if got != "[]float64{0.005, 0.1, 1}" {
			t.Errorf("got %q", got)
		}
	})
}
