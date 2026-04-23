//ff:func feature=gen-gogin type=util control=iteration dimension=1 topic=observability
//ff:what bucketsLiteral — []float64 를 Go 슬라이스 리터럴 문자열로 포맷

package middleware

import (
	"fmt"
	"strings"
)

// bucketsLiteral formats the histogram buckets as a Go expression. An empty
// or nil slice yields "prometheus.DefBuckets" so the generated code keeps
// the widely-recognised web-API default (0.005s ~ 10s).
func bucketsLiteral(buckets []float64) string {
	if len(buckets) == 0 {
		return "prometheus.DefBuckets"
	}
	parts := make([]string, 0, len(buckets))
	for _, b := range buckets {
		parts = append(parts, fmt.Sprintf("%g", b))
	}
	return "[]float64{" + strings.Join(parts, ", ") + "}"
}
