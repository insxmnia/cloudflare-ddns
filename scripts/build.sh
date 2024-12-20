#!/bin/bash

# Build script for the Go project

# Ensure the /build directory exists
mkdir -p ../build

# Build the application
echo "Building the application..."
go build -o ../build/cloudflare-ddns ../cmd/

if [ $? -eq 0 ]; then
  echo "Build successful! Binary created in ../build/cloudflare-ddns"
else
  echo "Build failed!"
  exit 1
fi
