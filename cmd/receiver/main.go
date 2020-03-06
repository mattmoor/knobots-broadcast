package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/kelseyhightower/envconfig"
)

type Receiver struct {
	client cloudevents.Client

	// If the K_SINK environment variable is set, then events are sent there,
	// otherwise we simply reply to the inbound request.
	Target string `envconfig:"K_SINK" required:true`
}

func main() {
	client, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	r := Receiver{client: client}
	if err := envconfig.Process("", &r); err != nil {
		log.Fatal(err.Error())
	}

	if err := client.StartReceiver(context.Background(), r.ReceiveAndSend); err != nil {
		log.Fatal(err.Error())
	}
}

// ReceiveAndSend is invoked whenever we receive an event.
func (recv *Receiver) ReceiveAndSend(ctx context.Context, event cloudevents.Event) error {
	ctx = cloudevents.ContextWithTarget(ctx, recv.Target)
	_, _, err := recv.client.Send(ctx, event)
	return err
}
