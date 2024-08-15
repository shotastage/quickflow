#!/usr/bin/env bash

# Error handling: Stop the script if any command fails
set -e

# Define necessary directory paths
EXPORT_DIR="./export"
DEST_DIR="$HOME/.quickflow-dev/bin"
REPO_URL="https://github.com/air-verse/air.git"
REPO_DIR="air"

# Clone the repository
git clone $REPO_URL
cd $REPO_DIR

# Run make release
make release

# Check if the export directory exists, if not create it
if [ ! -d "$EXPORT_DIR" ]; then
  mkdir -p "$EXPORT_DIR"
  echo "Created export directory at $EXPORT_DIR."
fi

# Copy the binary to the export directory
cp ./bin/darwin/air $EXPORT_DIR/air

# Clean up the working directory
cd ..
rm -rf $REPO_DIR

# Create QuickFlow Dev directory if it doesn't exist
if [ ! -d "$HOME/.quickflow-dev" ]; then
  mkdir -p "$HOME/.quickflow-dev/bin"
  echo "Created QuickFlow Dev Workspace."
else
  echo "QuickFlow Dev Workspace is already configured."
fi

# Check if the destination directory exists and create if necessary
if [ ! -d "$DEST_DIR" ]; then
  mkdir -p "$DEST_DIR"
fi

# Move the binary to the destination directory
mv $EXPORT_DIR/air $DEST_DIR/air
echo "Binary moved to $DEST_DIR/air."

# Clean up the export directory
if [ -d "$EXPORT_DIR" ]; then
  rm -rf "$EXPORT_DIR"
  echo "Cleaned up the export directory."
fi

# Check if $HOME/.quickflow-dev/bin is in the PATH
if [[ ":$PATH:" != *":$HOME/.quickflow-dev/bin:"* ]]; then
  # If not in PATH, add it to ~/.bashrc or ~/.zshrc depending on the shell
  if [ -n "$BASH_VERSION" ]; then
    echo 'export PATH="$HOME/.quickflow-dev/bin:$PATH"' >> "$HOME/.bashrc"
    echo "$HOME/.quickflow-dev/bin was added to your PATH in ~/.bashrc."
  elif [ -n "$ZSH_VERSION" ]; then
    echo 'export PATH="$HOME/.quickflow-dev/bin:$PATH"' >> "$HOME/.zshrc"
    echo "$HOME/.quickflow-dev/bin was added to your PATH in ~/.zshrc."
  fi
  # Apply the changes to the current shell session
  export PATH="$HOME/.quickflow-dev/bin:$PATH"
else
  echo "$HOME/.quickflow-dev/bin is already in your PATH."
fi
