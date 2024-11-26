#!/bin/zsh

# Save the current directory
CURRENT_DIR=$(pwd)

# Target directory where the Goose migrations are located
TARGET_DIR="sql/schema"

# Change to the target directory
cd "$TARGET_DIR" || { echo "Failed to cd to $TARGET_DIR"; exit 1; }

# Run the Goose migration command
goose postgres postgres://nsanders:@localhost:5432/hyperboreacs up || { echo "Goose command failed"; cd "$CURRENT_DIR"; exit 1; }

# Return to the original directory
cd "$CURRENT_DIR" || { echo "Failed to return to $CURRENT_DIR"; exit 1; }

# Confirmation message
echo "Migrations completed and returned to $CURRENT_DIR"
