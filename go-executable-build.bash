#!/usr/bin/env bash

# For each defined platform/architecture build a binary from `main.go` in current directory named for basename of current directory and zip it basename-platform-architecture

# based on an example from https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-20-04 with iteration asssitance from ChatGPT

# Get the current directory name
dirname=$(basename "$PWD")

# Specify the path to the main.go file
package=./main.go
if [[ -z "$package" ]]; then
    echo "usage: $0 <package-name>"
    exit 1
fi

platforms=("linux/amd64" "darwin/arm64" "darwin/amd64" "windows/amd64")

for platform in "${platforms[@]}"; do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$dirname
    subdir_name=$GOOS'-'$GOARCH
    # Create a subdirectory with the name "platform-architecture"
    mkdir "$subdir_name"

    if [ "$GOOS" = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS="$GOOS" GOARCH="$GOARCH" go build -o ./"$subdir_name"/"$output_name" $package

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi

    # Zip the binary file and remove the original file
    zip -j "$dirname"-"$subdir_name".zip "$subdir_name"/"$output_name"

    # Remove the subdirectory
    rm -rf "$subdir_name"
done
