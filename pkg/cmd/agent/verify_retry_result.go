//ff:type feature=agent type=helper
//ff:what verifyRetryResult — OpenAPI verify-retry 1회차 결과

package agent

type verifyRetryResult struct {
	verified bool
	stopped  bool
}
