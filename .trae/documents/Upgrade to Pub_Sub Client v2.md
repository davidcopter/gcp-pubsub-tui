I will upgrade the Google Cloud Pub/Sub dependency to the latest version (`v2`) as recommended by the deprecation warning.

### **Implementation Steps**

1.  **Update Imports**
    *   Replace `cloud.google.com/go/pubsub` with `cloud.google.com/go/pubsub/v2` in the following files:
        *   `internal/client/client.go`
        *   `internal/tui/model.go`
        *   `cmd/pubsub-tui/main.go`
        *   `test/integration/integration_test.go`

2.  **Update Dependencies**
    *   Run `go mod tidy` to resolve the new dependency and remove the old one.

3.  **Verify Build & Tests**
    *   Run `go build ./...` to ensure no compilation errors (checking for breaking API changes).
    *   Run `go test ./...` to ensure functionality remains intact.

### **Why this is needed**
The `cloud.google.com/go/pubsub` package has been deprecated in favor of `v2`, which supports modern Go idioms and long-term stability.
