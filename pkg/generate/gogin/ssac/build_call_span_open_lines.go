//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildCallSpanOpenLines — @call span.Start 라인 (WrapCalls=true)

package ssac

import "fmt"

func buildCallSpanOpenLines(pkgName, callFunc string) []string {
	spanName := fmt.Sprintf("call.%s.%s", pkgName, callFunc)
	return []string{
		fmt.Sprintf("callCtx, callSpan := otel.Tracer(\"ssac\").Start(ctx, %q)", spanName),
	}
}
