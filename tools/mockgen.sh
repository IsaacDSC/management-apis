#!/bin/bash

# Directory containing the source files
SOURCE_DIR="internal/management/adapter"
# Directory to place the generated mocks
DEST_DIR="mocks"

# Create the destination directory if it doesn't exist
mkdir -p "$DEST_DIR"

echo "\n[*] starting to generate mocks\n"

# Loop through each Go file in the source directory
for file in "$SOURCE_DIR"/*.go; do
  # Extract the base name of the file (without the directory and extension)
  file_name=$(basename "$file" .go)
  # Run the mockgen command
  echo "$DEST_DIR/${file_name}_mock.go"
  mockgen -source="$file" -destination="$DEST_DIR/${file_name}_mock.go" -package=mocks
done


echo "\n[*] Mocks generated successfully\n"