//ff:type feature=gen-ir type=model
//ff:what TriggerKind -- 서비스 함수 트리거 유형 (HTTP | Subscribe)

package ir

// TriggerKind represents how a service function is triggered.
type TriggerKind string

const (
	// TriggerHTTP indicates the service is triggered by an HTTP request.
	TriggerHTTP TriggerKind = "HTTP"
	// TriggerSubscribe indicates the service is triggered by a queue message.
	TriggerSubscribe TriggerKind = "Subscribe"
)
