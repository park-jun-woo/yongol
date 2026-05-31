//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what renderPrometheusSources — prometheus 템플릿을 파일별 소스 맵으로 치환

package middleware

import "strings"

// renderPrometheusSources substitutes __BUCKETS__ in the middleware template
// and returns a map of filename to source content for multi-file emit.
func renderPrometheusSources(buckets []float64) map[string]string {
	lit := bucketsLiteral(buckets)
	return map[string]string{
		"prometheus_middleware.go": strings.ReplaceAll(prometheusMiddlewareSourceTemplate, "__BUCKETS__", lit),
		"prometheus_handler.go":    prometheusHandlerSource,
	}
}
