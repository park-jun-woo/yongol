//ff:type feature=tsx-parser type=model
//ff:what APICall — apiClient.<op>(...) 호출 1건 + 위치 정보

package tsx

// APICall captures a single apiClient.<operationId>(...) invocation.
// Kind is a best-effort hint reserved for future heuristics (e.g. useQuery
// vs useMutation surrounding context). Phase001 always stores "raw".
type APICall struct {
	OperationID string       // "listWorkflows"
	Kind        string       // reserved: "query" | "mutation" | "raw"
	Args        []ArgBinding // keys of the first ObjectExpression argument
	Line        int          // 1-based source line of the callee
	Col         int          // 1-based source column of the callee
}
