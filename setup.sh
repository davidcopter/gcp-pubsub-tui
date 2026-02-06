#!/bin/bash

# Configuration
BINARY_NAME="gcp-pubsub-tui"
APP_NAME="pubsub-tui"
DIST_DIR="dist"

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Detect OS and Architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map Architecture to Go conventions
case "$ARCH" in
    x86_64) GOARCH="amd64" ;;
    aarch64|arm64) GOARCH="arm64" ;;
    *) 
        echo -e "${RED}Unsupported architecture: $ARCH${NC}"
        exit 1 
        ;;
esac

# Map OS to Go conventions (mostly handles cases like MINGW/CYGWIN if running in git bash)
case "$OS" in
    linux*) GOOS="linux" ;;
    darwin*) GOOS="darwin" ;;
    msys*|mingw*|cygwin*) GOOS="windows" ;;
    *) 
        echo -e "${RED}Unsupported OS: $OS${NC}"
        exit 1 
        ;;
esac

TARGET_BINARY="$APP_NAME-$GOOS-$GOARCH"
if [ "$GOOS" == "windows" ]; then
    TARGET_BINARY+=".exe"
    BINARY_NAME+=".exe"
fi

echo -e "Detected System: ${GREEN}$GOOS/$GOARCH${NC}"

# Check if binary exists in root
if [ -f "./$BINARY_NAME" ]; then
    echo -e "${GREEN}Binary '$BINARY_NAME' already exists in root. Skipping setup.${NC}"
    exit 0
fi

echo -e "${YELLOW}Binary not found in root. Searching dist directory...${NC}"

# Function to copy binary
copy_binary() {
    echo -e "Copying $1 to ./$BINARY_NAME..."
    cp "$1" "./$BINARY_NAME"
    chmod +x "./$BINARY_NAME"
    echo -e "${GREEN}Setup complete! You can now run ./$BINARY_NAME${NC}"
}

# Check dist directory
DIST_PATH="$DIST_DIR/$TARGET_BINARY"

if [ -f "$DIST_PATH" ]; then
    echo -e "${GREEN}Found matching binary in dist: $TARGET_BINARY${NC}"
    copy_binary "$DIST_PATH"
else
    echo -e "${YELLOW}Binary not found in dist. Rebuilding...${NC}"
    
    # Run build script
    if [ -f "./scripts/build.sh" ]; then
        ./scripts/build.sh
        
        # Check again after build
        if [ -f "$DIST_PATH" ]; then
            copy_binary "$DIST_PATH"
        else
            echo -e "${RED}Build failed or binary not created at expected path: $DIST_PATH${NC}"
            exit 1
        fi
    else
        echo -e "${RED}Build script not found at ./scripts/build.sh${NC}"
        exit 1
    fi
fi
