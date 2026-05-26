//ff:type feature=gen-ir type=model
//ff:what PublishOp -- @publish 시퀀스의 IR 표현 (메시지 큐 발행)

package ir

// PublishOp represents a @publish sequence: publishing an event message to a
// queue topic.
type PublishOp struct {
	// Topic is the queue topic (e.g. "order.completed").
	Topic string

	// Payload are the key-value pairs forming the message payload.
	Payload []FieldArg

	// Options are optional publish parameters (e.g. delay).
	Options []FieldArg
}
