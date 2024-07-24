package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Helper function to download a file
func downloadFile(url, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	// URL of the Wikipedia page
	url := "https://en.wikipedia.org/wiki/Rider%E2%80%93Waite_Tarot"

	// Create a directory to save images
	dir := "../static/images"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// List of filenames to exclude
	excludeList := map[string]bool{
		"Arthur_Waite_Author.JPG":            true,
		"Commons-logo.svg":                   true,
		"P_religion_world.svg":               true,
		"Pamela_Colman_Smith_circa_1912.jpg": true,
		"People_icon.svg":                    true,
		"Symbol_category_class.svg":          true,
	}

	// Request the page
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	// Find all image tags
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		imgSrc, exists := s.Attr("src")
		if !exists {
			return
		}

		imgURL := "https:" + imgSrc

		// Change the URL to get the full resolution image
		if strings.Contains(imgURL, "thumb") {
			imgURL = strings.Replace(imgURL, "thumb/", "", 1)
			imgURL = imgURL[:strings.LastIndex(imgURL, "/")]
		}

		// Extract image name
		imgName := path.Base(imgURL)

		// Check if the image is in the exclude list
		if excludeList[imgName] {
			fmt.Println("Excluding:", imgName)
			return
		}

		// Download the image
		filepath := path.Join(dir, imgName)
		fmt.Println("Downloading:", imgURL)
		if err := downloadFile(imgURL, filepath); err != nil {
			fmt.Println("Error downloading image:", err)
		} else {
			fmt.Println("Downloaded:", imgName)
		}
	})
}
