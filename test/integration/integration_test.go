package integration

import (
	"context"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/pubsub"
	"gcp-pubsub-tui/internal/client"
)

func TestPubSubIntegration(t *testing.T) {
	if os.Getenv("PUBSUB_EMULATOR_HOST") == "" {
		t.Skip("Skipping integration test: PUBSUB_EMULATOR_HOST not set")
	}

	ctx := context.Background()
	projID := "test-project"

	// Create raw client to setup topic/sub
	rawClient, err := pubsub.NewClient(ctx, projID)
	if err != nil {
		t.Fatal(err)
	}
	defer rawClient.Close()

	topicID := "test-topic"
	subID := "test-sub"

	// Cleanup (ignore errors if not exist)
	rawClient.Topic(topicID).Delete(ctx)
	rawClient.Subscription(subID).Delete(ctx)

	// Create Topic
	topic, err := rawClient.CreateTopic(ctx, topicID)
	if err != nil {
		t.Fatal(err)
	}

	// Create Sub
	_, err = rawClient.CreateSubscription(ctx, subID, pubsub.SubscriptionConfig{
		Topic: topic,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Publish a message
	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(`{"test":"data"}`),
	})
	_, err = res.Get(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Now use OUR client
	// We need a dummy credential file for the constructor validation
	f, _ := os.CreateTemp("", "dummy-creds.json")
	f.Write([]byte(`{"project_id":"test-project"}`))
	f.Close()
	defer os.Remove(f.Name())

	c, err := client.NewClient(ctx, projID, f.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	msgChan := make(chan *pubsub.Message)

	// Create a context with timeout for the subscriber
	subCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	go func() {
		// Subscribe blocks
		c.Subscribe(subCtx, subID, true, msgChan)
	}()

	select {
	case msg := <-msgChan:
		if string(msg.Data) != `{"test":"data"}` {
			t.Errorf("Unexpected data: %s", msg.Data)
		}
	case <-subCtx.Done():
		t.Fatal("Timeout waiting for message")
	}
}
