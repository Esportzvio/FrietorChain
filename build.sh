#!/bin/bash

# Define the output directory for the builds
output_dir="builds"

# Ensure the output directory exists
mkdir -p "$output_dir"

# Get the list of all platforms and architectures supported by go tool dist
platforms_and_architectures=$(go tool dist list)

# Loop through each platform and architecture
for platform_arch in $platforms_and_architectures; do
    # Split the platform and architecture
    IFS='/' read -r -a parts <<< "$platform_arch"
    os="${parts[0]}"
    arch="${parts[1]}"

    # Set the output file name
    output_name="frietor_${os}_${arch}"
    echo "Building $output_name..."

    # Build the program for the current platform
    GOOS="$os" GOARCH="$arch" go build -o "$output_dir/$output_name" .

    # Check if the build was successful
    if [ $? -eq 0 ]; then
        echo "Build for $platform_arch completed successfully."
    else
        echo "Build for $platform_arch failed."
    fi
done

echo "All builds completed."
