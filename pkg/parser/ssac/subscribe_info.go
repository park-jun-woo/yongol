//ff:type feature=ssac-parse type=model
//ff:what 큐 구독 트리거 정보 타입
package ssac

// SubscribeInfo는 큐 구독 트리거 정보다.
type SubscribeInfo struct {
	Topic       string // "order.completed"
	MessageType string // "OnOrderCompletedMessage"
}
