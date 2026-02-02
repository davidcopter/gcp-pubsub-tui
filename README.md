# Google Cloud Pub/Sub TUI

A professional, interactive Terminal User Interface (TUI) for subscribing to Google Cloud Pub/Sub topics. Built with Go and the Charm ecosystem (`bubbletea`, `lipgloss`, `huh`).

## Features

*   **Interactive Setup**: Guided configuration for subscription details.
*   **Visual Message Display**: Beautifully formatted messages with color-coded metadata.
*   **JSON Pretty Printing**: Automatically formats JSON payloads for readability.
*   **Real-time Streaming**: Streams messages using the official Google Cloud Pub/Sub client.
*   **Graceful Shutdown**: Handles signals correctly.

## Installation

```bash
go install github.com/davidcopter/gcp-pubsub-tui/cmd/pubsub-tui@latest
```

Or build from source:

```bash
git clone https://github.com/davidcopter/gcp-pubsub-tui.git
cd gcp-pubsub-tui
go build -o pubsub-tui ./cmd/pubsub-tui
```

### Cross-Platform Build

To build binaries for multiple platforms (Linux, macOS, Windows), run the provided build script:

```bash
./scripts/build.sh
```

The binaries will be generated in the `dist/` directory.

## Usage

You must provide a Google Cloud Service Account JSON key file.

```bash
# Using flag
./pubsub-tui -key /path/to/service-account.json

# Using environment variable
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/service-account.json"
./pubsub-tui
```

### Interactive Controls

*   **Setup Screen**:
    *   Enter Subscription ID (e.g., `my-sub`).
    *   Enable/Disable Auto-Acknowledgment.
    *   Set Pull Interval (Simulated/Info).
*   **Main Dashboard**:
    *   `q` or `Ctrl+C`: Quit the application.
    *   `c`: Clear the message history.

## Configuration Guide

1.  **Create a Service Account**:
    *   Go to Google Cloud Console > IAM & Admin > Service Accounts.
    *   Create a new Service Account.
    *   Grant the `Pub/Sub Subscriber` role.
2.  **Generate Key**:
    *   Go to the "Keys" tab of the Service Account.
    *   Add Key > Create new key > JSON.
    *   Save the file securely.

## Development

### Running Tests

```bash
go test ./...
```

### Integration Tests

To run integration tests, you need the Pub/Sub Emulator running.

```bash
# Start Emulator
gcloud beta emulators pubsub start --project=test-project

# Run Tests
export PUBSUB_EMULATOR_HOST=localhost:8085
go test ./test/integration/...
```
