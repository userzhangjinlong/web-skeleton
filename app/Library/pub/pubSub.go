package pub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
)

//初始化pubClient
func NewPubClient(projectId string) (client *pubsub.Client, err error) {
	var ctx context.Context
	client, err = pubsub.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}

	return
}

//消息生产
func Produce(publishMessage string) (string, error) {
	var ctx context.Context
	client, err := NewPubClient("testProjectId")
	if err != nil {
		return "", err
	}

	topic := client.Topic("topic1")

	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(publishMessage),
	})

	msgID, err := res.Get(ctx)
	if err != nil {
		return "", err
	}

	return msgID, nil
}

//消息消费
func Consumer(msgId string) error {
	var ctx context.Context
	client, err := NewPubClient("testProjectId")
	if err != nil {
		return err
	}

	sub := client.Subscription(msgId)
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		//消息消费
		fmt.Println(m.Data)
		m.Ack() // Acknowledge that we've consumed the message.
	})
	if err != nil {
		return err
	}
	return nil
}
