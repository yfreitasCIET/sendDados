package messaging

import (
	"context"
	"fmt"

	pubsub "cloud.google.com/go/pubsub"
)

type Messaging struct {
	sub     *pubsub.Subscription
	context context.Context
}

func CreateNewMessaging() (*Messaging, error) {
	client, err := pubsub.NewClient(context.Background(), "itau-techpartner")
	if err != nil {
		return nil, err
	}

	sub := client.Subscription("tech-partner-incoming-messages-sub")
	return &Messaging{
		sub:     sub,
		context: context.Background(),
	}, nil
}

func (m *Messaging) handleMessage(ctx context.Context, msg *pubsub.Message) {

	fmt.Printf("%v", string(msg.Data))

	msg.Ack()
}

func (m *Messaging) Start() {
	m.sub.Receive(m.context, m.handleMessage)
}

func (m *Messaging) Stop() {
	m.context.Done()
}
