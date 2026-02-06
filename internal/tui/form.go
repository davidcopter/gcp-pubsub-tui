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

func RunSetupForm(initialSubID string) (*SubscriptionOptions, error) {
	var (
		subID        string = initialSubID
		autoAck      bool   = true
		intervalStr  string = "1"
	)

	// If subscription ID is already provided, skip the input step for it or pre-fill it
	// For better UX, if provided via flag, we can either skip the form entirely if we had defaults for others,
	// or just pre-fill. Let's pre-fill for now, but if the user wants fully non-interactive, 
	// they would need more flags. 
	// However, the request implies using the flag might skip the form or at least pre-fill.
	// Let's assume if subID is provided, we still show form to confirm other options unless we want to be fully automatic.
	// But usually CLI flags override prompts.
	
	// Let's just pre-fill. If you want to skip, we can check if subID != "" and return immediately with defaults.
	if subID != "" {
		// If subscription is provided, we can return immediately with defaults
		// OR we can just pre-fill. The prompt says "Interactive prompt system", 
		// but flags usually imply automation.
		// Let's assume if flag is present, we use it. But we still need other params.
		// Let's pre-fill the value in the form.
	}

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
