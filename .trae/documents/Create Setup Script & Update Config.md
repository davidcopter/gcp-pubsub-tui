I will create a setup script to automate binary preparation and update the environment configuration.

### **Implementation Steps**

1.  **Create `setup.sh`**
    *   **Detect Environment**: Identify the current OS (Linux/Darwin/Windows) and Architecture (amd64/arm64).
    *   **Check Existence**: Look for `gcp-pubsub-tui` in the root.
    *   **Search Dist**: If missing, check `dist/pubsub-tui-<os>-<arch>`.
    *   **Build Fallback**: If not found in `dist/`, clean `dist/`, run `./scripts/build.sh`, and then copy the correct binary.
    *   **Finalize**: Rename the binary to `gcp-pubsub-tui` in the root directory.

2.  **Update `.gitignore`**
    *   Add `/gcp-pubsub-tui` to ignore the local binary.

3.  **Update `README.md`**
    *   Add a **"Quick Start"** section explaining how to use `./setup.sh`.

4.  **Set Permissions**
    *   Make `setup.sh` executable (`chmod +x`).

### **Script Logic**
The script will ensure that a usable binary is always available in the project root, either by reusing a pre-built one or building from source.
