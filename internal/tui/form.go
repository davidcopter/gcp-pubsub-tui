package tui

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
)

type SubscriptionOptions struct {
	SubscriptionID string
	AutoAck        bool
	PullInterval   int // in seconds
}

func RunSetupForm() (*SubscriptionOptions, error) {
	var (
		subID       string
		autoAck     bool   = true
		intervalStr string = "1"
	)

	// Create the form
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Subscription ID").
				Description("Enter the Google Cloud Pub/Sub Subscription ID").
				Value(&subID).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("subscription ID is required")
					}
					return nil
				}),

			huh.NewConfirm().
				Title("Auto Acknowledge").
				Description("Automatically acknowledge messages upon receipt?").
				Value(&autoAck),

			huh.NewInput().
				Title("Pulling Interval (seconds)").
				Description("How often to pull for new messages (default 1s)").
				Value(&intervalStr).
				Validate(func(str string) error {
					if str == "" {
						return nil // allow empty for default
					}
					val, err := strconv.Atoi(str)
					if err != nil || val < 0 {
						return errors.New("must be a positive integer")
					}
					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		return nil, err
	}

	interval, _ := strconv.Atoi(intervalStr)
	if interval == 0 {
		interval = 1
	}

	options := &SubscriptionOptions{
		SubscriptionID: subID,
		AutoAck:        autoAck,
		PullInterval:   interval,
	}

	// Show confirmation
	fmt.Println("\nConfiguration Confirmed:")
	fmt.Printf("  Subscription: %s\n", options.SubscriptionID)
	fmt.Printf("  Auto-Ack:     %v\n", options.AutoAck)
	fmt.Printf("  Interval:     %ds\n", options.PullInterval)
	fmt.Println("\nConnecting to Pub/Sub...")

	return options, nil
}
