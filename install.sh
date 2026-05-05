#!/bin/bash

# Configuration
REPO="rounakkraaj-1744/dawker"
BINARY_NAME="dawker"

# Detect OS and Arch
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [[ "$ARCH" == "x86_64" ]]; then
    ARCH="amd64"
elif [[ "$ARCH" == "arm64" || "$ARCH" == "aarch64" ]]; then
    ARCH="arm64"
fi

# Map OS name to Go/Goreleaser style
case "$OS" in
    darwin)  OS="Darwin" ;;
    linux)   OS="Linux" ;;
    msys*|cygwin*|mingw*) OS="Windows" ;;
esac

# Get the latest release tag from GitHub
LATEST_TAG=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep "tag_name" | cut -d '"' -f 4)

if [ -z "$LATEST_TAG" ]; then
    echo "Error: Could not find the latest release for $REPO. Ensure you have created a Release on GitHub."
    exit 1
fi

# Construct download URL (Standard Goreleaser naming: dawker_Darwin_arm64.tar.gz)
DOWNLOAD_URL="https://github.com/$REPO/releases/download/${LATEST_TAG}/${BINARY_NAME}_${OS}_${ARCH}.tar.gz"

echo "Downloading $BINARY_NAME $LATEST_TAG for $OS/$ARCH..."
curl -sSL "$DOWNLOAD_URL" -o "${BINARY_NAME}.tar.gz"

if [ $? -ne 0 ]; then
    echo "Error: Download failed. Check your internet connection or if the release exists."
    exit 1
fi

# Extract and install
echo "Installing to /usr/local/bin..."
tar -xzf "${BINARY_NAME}.tar.gz"
chmod +x "$BINARY_NAME"
sudo mv "$BINARY_NAME" /usr/local/bin/

# Cleanup
rm "${BINARY_NAME}.tar.gz"

echo "Successfully installed $BINARY_NAME to /usr/local/bin!"
dawker --version
