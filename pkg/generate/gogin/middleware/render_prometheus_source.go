//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what renderPrometheusSource — prometheus 템플릿의 __BUCKETS__ 를 buckets 리터럴로 치환

package middleware

import "strings"

// renderPrometheusSource substitutes __BUCKETS__ with the buckets literal.
// When buckets is empty prometheus.DefBuckets is emitted so operators can
// still tune the slice without regenerating.
func renderPrometheusSource(buckets []float64) string {
	lit := bucketsLiteral(buckets)
	return strings.ReplaceAll(prometheusSourceTemplate, "__BUCKETS__", lit)
}
