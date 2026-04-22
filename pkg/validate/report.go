//ff:type feature=validate type=model
//ff:what Report — Validate 실행 전체 step 결과 집합
package validate

// Report aggregates all step results from a Validate run.
type Report struct {
	Steps []StepResult
}
