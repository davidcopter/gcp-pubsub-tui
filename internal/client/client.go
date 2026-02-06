package client

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub/v2"
	"cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	"google.golang.org/api/option"
)

type PubSubClient struct {
	client *pubsub.Client
}

// NewClient creates a new Google Cloud Pub/Sub client
func NewClient(ctx context.Context, projectID string, keyPath string) (*PubSubClient, error) {
	if projectID == "" {
		return nil, fmt.Errorf("project ID is required")
	}
	
	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(keyPath))
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %w", err)
	}
	return &PubSubClient{client: client}, nil
}

// Subscribe starts streaming messages from the subscription to the provided channel
func (c *PubSubClient) Subscribe(ctx context.Context, subID string, autoAck bool, msgChan chan<- *pubsub.Message) error {
	// Check if subscription exists using SubscriptionAdminClient
	// We need to construct the fully qualified name
	fullSubName := fmt.Sprintf("projects/%s/subscriptions/%s", c.client.Project(), subID)
	
	_, err := c.client.SubscriptionAdminClient.GetSubscription(ctx, &pubsubpb.GetSubscriptionRequest{
		Subscription: fullSubName,
	})
	if err != nil {
		return fmt.Errorf("failed to check subscription existence (or sub does not exist): %w", err)
	}

	// Get subscriber client
	sub := c.client.Subscriber(subID)
	
	// Start receiving messages
	// We use the context to cancel the subscription
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		if autoAck {
			msg.Ack()
		}
		
		// Non-blocking send or blocking?
		// If we block, we might hold up the stream.
		// Use a select to respect context cancellation.
		select {
		case msgChan <- msg:
		case <-ctx.Done():
		}
	})

	if err != nil {
		return fmt.Errorf("subscription error: %w", err)
	}

	return nil
}

// Close closes the underlying Pub/Sub client
func (c *PubSubClient) Close() error {
	return c.client.Close()
}
