package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub/v2"
	tea "github.com/charmbracelet/bubbletea"

	"gcp-pubsub-tui/internal/client"
	"gcp-pubsub-tui/internal/config"
	"gcp-pubsub-tui/internal/tui"
)

func main() {
	// Setup file logging to avoid interfering with TUI
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	// 1. Config & Auth
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		os.Exit(1)
	}

	// 2. Setup Form (Interactive)
	// This runs before Bubble Tea takes over the full screen
	opts, err := tui.RunSetupForm()
	if err != nil {
		fmt.Printf("Setup aborted: %v\n", err)
		os.Exit(1)
	}

	// 3. Initialize Client
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if cfg.ProjectID == "" {
		fmt.Println("Error: Could not determine Project ID from Service Account file.")
		os.Exit(1)
	}

	pubsubClient, err := client.NewClient(ctx, cfg.ProjectID, cfg.ServiceAccountPath)
	if err != nil {
		fmt.Printf("Failed to create Pub/Sub client: %v\n", err)
		os.Exit(1)
	}
	defer pubsubClient.Close()

	// 4. Start Subscription
	msgChan := make(chan *pubsub.Message)

	// Run subscriber in background
	go func() {
		log.Printf("Starting subscription to %s", opts.SubscriptionID)
		// Note: Subscribe blocks until error or context cancel
		err := pubsubClient.Subscribe(ctx, opts.SubscriptionID, opts.AutoAck, msgChan)
		if err != nil {
			log.Printf("Subscription error: %v", err)
			// In a more complex app, we'd send an error Msg to the model
			close(msgChan)
		}
	}()

	// 5. Start TUI
	p := tea.NewProgram(
		tui.NewModel(opts.SubscriptionID, msgChan),
		tea.WithAltScreen(),       // Use full screen
		tea.WithMouseCellMotion(), // Enable mouse
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
