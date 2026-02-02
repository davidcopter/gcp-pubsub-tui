I will create a comprehensive Bash build script to cross-compile the application for major operating systems and architectures.

### **Implementation Steps**

1.  **Create Build Script (`scripts/build.sh`)**
    *   Define target platforms:
        *   **Linux**: `amd64`, `arm64`
        *   **macOS (Darwin)**: `amd64` (Intel), `arm64` (Apple Silicon)
        *   **Windows**: `amd64`
    *   Set output directory to `dist/`.
    *   Loop through targets and execute `go build` with appropriate `GOOS` and `GOARCH` environment variables.
    *   Handle `.exe` extension for Windows binaries.
    *   Add error handling to stop on build failures.

2.  **Make Script Executable**
    *   Run `chmod +x scripts/build.sh`.

3.  **Update Documentation**
    *   Add a "Building for Release" section to `README.md` explaining how to use the script.

### **Script Logic Preview**
The script will iterate over a list of platforms (e.g., "linux/amd64", "darwin/arm64") and output binaries named `pubsub-tui-<os>-<arch>`.
