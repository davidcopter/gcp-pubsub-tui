#!/bin/bash

# Configuration
APP_NAME="pubsub-tui"
OUTPUT_DIR="dist"
MAIN_FILE="cmd/pubsub-tui/main.go"

# Platforms to build for
PLATFORMS=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
)

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo "Building $APP_NAME..."
mkdir -p $OUTPUT_DIR

# Clean previous builds
rm -rf $OUTPUT_DIR/*

for PLATFORM in "${PLATFORMS[@]}"; do
    # Split platform into OS and ARCH
    IFS='/' read -r -a PARTS <<< "$PLATFORM"
    GOOS="${PARTS[0]}"
    GOARCH="${PARTS[1]}"
    
    OUTPUT_NAME="$APP_NAME-$GOOS-$GOARCH"
    
    # Add .exe extension for Windows
    if [ "$GOOS" == "windows" ]; then
        OUTPUT_NAME+=".exe"
    fi
    
    echo -n "Building for $PLATFORM... "
    
    env GOOS=$GOOS GOARCH=$GOARCH go build -o "$OUTPUT_DIR/$OUTPUT_NAME" $MAIN_FILE
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}Success${NC}"
    else
        echo -e "${RED}Failed${NC}"
        exit 1
    fi
done

echo -e "\n${GREEN}Build complete! Binaries are in $OUTPUT_DIR/${NC}"
ls -lh $OUTPUT_DIR
