//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what otelNoopCaseLines — switch "noop" case 본문 (exporter 없음)

package boot

// otelNoopCaseLines emits the `case "noop":` branch body. The SDK still
// runs so the instrumentation path is exercised end-to-end without any
// exported spans.
func otelNoopCaseLines() []string {
	return []string{
		`	case "noop":`,
		`		spanExporter = nil`,
	}
}
