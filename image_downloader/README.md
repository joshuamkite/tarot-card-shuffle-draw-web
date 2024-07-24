# Wikipedia Image Downloader

This Go script downloads images from a specified Wikipedia page and saves them to a local directory, excluding certain specified images.

## Usage

```sh
go run main.go
```

## Script Details

The script performs the following steps:

1. Sets the URL of the Wikipedia page to scrape.
2. Creates a directory to save images.
3. Defines a list of filenames to exclude from downloading.
4. Fetches the HTML content of the specified Wikipedia page.
5. Parses the HTML to find all image tags.
6. Processes each image tag to construct the full-resolution image URL.
7. Checks if the image is in the exclude list and skips it if necessary.
8. Downloads and saves each image to the specified directory.

## Directory Structure

```
project-root/
│
├── main.go
└── static/
    └── images/
        ├── Cups01.jpg
        ├── Cups02.jpg
        ├── ... (other downloaded images)
```
