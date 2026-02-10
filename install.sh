#!/bin/bash

set -e

echo "Installing gitadd-tui..."

# Download the binary
curl -L https://github.com/shrikant23codes/gitadd-tui/releases/latest/download/gitadd-tui -o /tmp/gitadd-tui

# Make it executable
chmod +x /tmp/gitadd-tui

# Move to bin (requires sudo)
sudo mv /tmp/gitadd-tui /usr/local/bin/

echo "✓ gitadd-tui installed successfully!"
echo "Run 'gitadd-tui' to get started"
