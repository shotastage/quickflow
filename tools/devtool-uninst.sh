#!/usr/bin/env bash

# Error handling: Stop the script if any command fails
set -e

# Define necessary directory paths
DEST_DIR="$HOME/.quickflow-dev/bin"
FULL_PATH="$DEST_DIR/air"

# Function to remove the air binary
remove_air() {
  if [ -f "$FULL_PATH" ]; then
    rm "$FULL_PATH"
    echo "Removed air binary from $FULL_PATH"
  else
    echo "air binary not found in $FULL_PATH"
  fi

  # Remove the QuickFlow Dev directory if it's empty
  if [ -d "$HOME/.quickflow-dev" ] && [ -z "$(ls -A "$HOME/.quickflow-dev")" ]; then
    rm -rf "$HOME/.quickflow-dev"
    echo "Removed empty QuickFlow Dev directory"
  fi
}

# Function to remove directory from PATH
remove_from_path() {
  # Determine the current shell and update the appropriate shell config file
  case "$SHELL" in
    */bash)
      if grep -q 'quickflow dev initialize' "$HOME/.bashrc"; then
        sed -i '' '/# >>> quickflow dev initialize >>>/,/# <<< quickflow dev initialize <<</d' "$HOME/.bashrc"
        echo "Removed QuickFlow Dev PATH entry from ~/.bashrc"
      else
        echo "QuickFlow Dev PATH entry not found in ~/.bashrc"
      fi
      ;;
    */zsh)
      if grep -q 'quickflow dev initialize' "$HOME/.zshenv"; then
        sed -i '' '/# >>> quickflow dev initialize >>>/,/# <<< quickflow dev initialize <<</d' "$HOME/.zshenv"
        echo "Removed QuickFlow Dev PATH entry from ~/.zshenv"
      else
        echo "QuickFlow Dev PATH entry not found in ~/.zshenv"
      fi
      ;;
    *)
      echo "Unknown shell. Please manually remove $HOME/.quickflow-dev/bin from your PATH if necessary."
      ;;
  esac
}

# Call the function to remove air
remove_air

# Call the function to remove the directory from the PATH
remove_from_path

echo "Cleanup complete. Please restart your terminal or run 'source ~/.bashrc' (for bash) or 'source ~/.zshenv' (for zsh) to apply PATH changes."
