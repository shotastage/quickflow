#!/usr/bin/env bash
# Error handling: Stop the script if any command fails
set -e

# Define necessary directory paths
DEST_DIR="$HOME/.quickflow-dev/bin"
REPO_URL="https://github.com/air-verse/air.git"
REPO_DIR="air"
GITLEAKS_REPO_URL="https://github.com/gitleaks/gitleaks.git"
GITLEAKS_REPO_DIR="gitleaks"

# Function to clone, build, and move the air binary
install_air() {
  # Clone the repository
  git clone $REPO_URL
  cd $REPO_DIR
  # Run make release
  make release
  # Check if the destination directory exists and create if necessary
  if [ ! -d "$DEST_DIR" ];then
    mkdir -p "$DEST_DIR"
  fi
  # Copy the binary directly to the destination directory
  cp ./bin/darwin/air $DEST_DIR/air
  echo "Binary copied to $DEST_DIR/air."
  # Clean up the working directory
  cd ..
  rm -rf $REPO_DIR
}

install_gitleaks() {
  # Clone the repository
  git clone $GITLEAKS_REPO_URL
  cd $GITLEAKS_REPO_DIR
  # Build gitleaks
  make build
  # Check if the destination directory exists and create if necessary
  if [ ! -d "$DEST_DIR" ];then
    mkdir -p "$DEST_DIR"
  fi
  # Copy the binary to the destination directory
  cp ./gitleaks $DEST_DIR/gitleaks
  echo "Binary copied to $DEST_DIR/gitleaks."
  # Clean up the working directory
  cd ..
  rm -rf $GITLEAKS_REPO_DIR
}

# Function to add directory to PATH
add_to_path() {
  # Check if $HOME/.quickflow-dev/bin is in the PATH
  if [[ ":$PATH:" != *":$HOME/.quickflow-dev/bin:"* ]]; then
    # Determine the current shell and update the appropriate shell config file
    case "$SHELL" in
      */bash)
        if ! grep -q 'quickflow dev initialize' "$HOME/.bashrc"; then
          cat << 'EOF' >> "$HOME/.bashrc"
# >>> quickflow dev initialize >>>
export PATH="$HOME/.quickflow-dev/bin:$PATH"
# <<< quickflow dev initialize <<<
EOF
          echo "$HOME/.quickflow-dev/bin was added to your PATH in ~/.bashrc."
        else
          echo "quickflow dev is already in your PATH in ~/.bashrc."
        fi
        ;;
      */zsh)
        if ! grep -q 'quickflow dev initialize' "$HOME/.zshenv"; then
          cat << 'EOF' >> "$HOME/.zshenv"
# >>> quickflow dev initialize >>>
export PATH="$HOME/.quickflow-dev/bin:$PATH"
# <<< quickflow dev initialize <<<
EOF
          echo "$HOME/.quickflow-dev/bin was added to your PATH in ~/.zshenv."
        else
          echo "quickflow dev is already in your PATH in ~/.zshenv."
        fi
        ;;
      *)
        echo "Unknown shell. Please manually add $HOME/.quickflow-dev/bin to your PATH."
        ;;
    esac
    # Apply the changes to the current shell session
    export PATH="$HOME/.quickflow-dev/bin:$PATH"
  else
    echo "$HOME/.quickflow-dev/bin is already in your PATH."
  fi
}

# Create QuickFlow Dev directory if it doesn't exist
if [ ! -d "$HOME/.quickflow-dev" ]; then
  mkdir -p "$HOME/.quickflow-dev/bin"
  echo "Created QuickFlow Dev Workspace."
else
  echo "QuickFlow Dev Workspace is already configured."
fi

# Call the function to install air
install_air

# Call the function to install gitleaks
install_gitleaks

# Call the function to add the directory to the PATH
add_to_path
