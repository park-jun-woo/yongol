//ff:type feature=ssac-parse type=model
//ff:what SubscribeInfo — type representing queue subscription trigger metadata
package ssac

// SubscribeInfo holds queue subscription trigger metadata.
type SubscribeInfo struct {
	Topic       string // "order.completed"
	MessageType string // "OnOrderCompletedMessage"
}
