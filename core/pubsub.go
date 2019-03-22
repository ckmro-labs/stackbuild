package core

import "context"

// Message 定议
type Message struct {
	Repository string
	Data       []byte
}

//Pubsub 提供构建消息订阅/发布(广播)
type Pubsub interface {
	// Publish 广播所有消息至订阅者
	Publish(context.Context, *Message) error

	// Subscribe 订阅消息管道
	Subscribe(context.Context) (<-chan *Message, <-chan error)

	// Subscribers 返回消息订阅总数
	Subscribers() int
}
