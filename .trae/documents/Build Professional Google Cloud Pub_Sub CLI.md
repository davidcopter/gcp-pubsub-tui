I will develop a professional TUI application using the **Charm** ecosystem (`bubbletea`, `lipgloss`, `huh`) to provide a modern, interactive experience, including a feature to clear the message display.

### **Project Structure**
```
.
├── cmd
│   └── pubsub-tui
│       └── main.go           # Entry point
├── internal
│   ├── config            # Configuration & Authentication logic
│   ├── client            # Google Cloud Pub/Sub client wrapper
│   ├── tui
│   │   ├── form.go       # "huh" form for interactive setup
│   │   ├── model.go      # Bubble Tea main model (State & Key Handling)
│   │   └── styles.go     # Lipgloss definitions for the "Message Box" layout
├── pkg
│   └── utils             # Helper functions (JSON formatting)
├── go.mod
└── README.md
```

### **Implementation Steps**

1.  **Project Initialization**
    *   Initialize `go.mod`.
    *   Install dependencies: `bubbletea`, `lipgloss`, `huh`, `pubsub`.

2.  **Configuration & Authentication (`internal/config`)**
    *   Parse CLI args/Env vars for `GOOGLE_APPLICATION_CREDENTIALS`.
    *   Validate the Service Account file exists and is valid JSON.

3.  **Interactive Setup Phase (`internal/tui/form.go`)**
    *   Use `github.com/charmbracelet/huh` for:
        *   **Subscription Name** (Input, required).
        *   **Auto-Ack** (Confirm, default `true`).
        *   **Pulling Interval** (Input, validate as integer).
    *   Display a confirmation state before starting the listener.

4.  **Pub/Sub Client (`internal/client`)**
    *   Implement `Subscribe` method that streams messages to the TUI via a Go channel.
    *   Handle thread-safe message processing and graceful cancellation.

5.  **Main TUI Dashboard (`internal/tui`)**
    *   **State Management**: Store received messages in a slice.
    *   **Key Bindings**: 
        *   Implement **'c' or 'C'** key handling to clear the message list.
        *   Implement **'q' or 'ctrl+c'** for graceful exit.
    *   **View Layer**:
        *   **Header**: Show subscription details.
        *   **Message List**: Render messages in the requested "Boxed" format (Cyan IDs, Yellow Timestamps, Green Attributes, White JSON Data).
        *   **Footer**: Show help text (e.g., "c: clear • q: quit").

6.  **Main Logic (`cmd/pubsub-tui/main.go`)**
    *   Orchestrate Config -> Setup Form -> Bubble Tea Program loop.

7.  **Quality Assurance**
    *   **Unit Tests**: Test config, JSON formatting, and TUI "Clear" logic.
    *   **Integration Tests**: Connectivity test with Pub/Sub Emulator.
    *   **Documentation**: Detailed README with instructions and key bindings.

### **Key Dependencies**
*   `github.com/charmbracelet/bubbletea`: TUI framework.
*   `github.com/charmbracelet/lipgloss`: Layout and styling.
*   `github.com/charmbracelet/huh`: Forms.
*   `cloud.google.com/go/pubsub`: GCP Pub/Sub client.
