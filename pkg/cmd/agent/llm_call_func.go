//ff:type feature=agent type=adapter
//ff:what LLMCallFunc — LLM 호출 함수 시그니처 타입

package agent

// LLMCallFunc is the signature for LLM call functions used by scaffold.
type LLMCallFunc func(backend, model, system, user string) (string, error)
