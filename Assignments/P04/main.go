package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "time"
	"strconv"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {

	length := len(urls)

	for i := 0; i < length; i++ {
	
		j := i + 1
	strNumber := strconv.Itoa(j)
     downloadImage(urls[i], "image" + strNumber + ".jpg")
	 fmt.Printf("sequential image saved successfully \n")
    }

}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	
	channel := make(chan string)

	for i, url := range urls {

		j := i + 1
		strNumber := strconv.Itoa(j)

		go func(url string) {

			//download individual image
			filename := "conccurent_image_" + strNumber  + ".jpg"
			img := downloadImage(url, filename)


			if img != nil {
				fmt.Printf("Error downloading %s: %s\n", url, img)
				return
			}
			channel <- filename
		}(url)
	}
	
	//print success message
	for i := 0; i < len(urls); i++ {
		filename := <-channel
		if filename != "" {
			fmt.Printf("conccurent image saved successfully \n")
		}
	}
	//close channel after finishing
	close(channel)

}


func main() {


    urls := []string{
		"https://images.unsplash.com/photo-1695653421371-cd48246a6200?q=80&w=2072&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1699898645601-4ab442ed6994?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
        "https://plus.unsplash.com/premium_photo-1699796414302-b34941b175ec?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
        "https://cdn.pixabay.com/photo/2023/11/06/04/53/woman-8368830_1280.jpg",
		"https://cdn.pixabay.com/photo/2023/10/12/08/24/bird-8310183_1280.png",
    }

    // Sequential download
    start := time.Now()
    downloadImagesSequential(urls)
	fmt.Printf("\n")
    fmt.Printf("Sequential download took: %v\n", time.Since(start))
	fmt.Printf("\n")
	fmt.Printf("\n")



    // Concurrent download
    start = time.Now()
    downloadImagesConcurrent(urls)
	fmt.Printf("\n")
    fmt.Printf("Concurrent download took: %v\n", time.Since(start))
	fmt.Printf("\n")

}



// Helper function to download and save a single image.
func downloadImage(url, filename string) error {


// Create a new `http.Request` object.
req, err := http.NewRequest("GET", url, nil)
if err != nil {
	fmt.Println(err)
	return err
}

// Create a new `http.Client` object.
client := &http.Client{}

// Do the request and get the response.
resp, err := client.Do(req)
if err != nil {
	fmt.Println(err)
	return err
}

// Check the response status code.
if resp.StatusCode != http.StatusOK {
	fmt.Println("Response status code:", resp.StatusCode)
	return err
}

// Create a new file to save the image to.
f, err := os.Create(filename)
if err != nil {
	fmt.Println(err)
	return err
}

// Copy the image from the response body to the file.
_, err = io.Copy(f, resp.Body)
if err != nil {
	fmt.Println(err)
	return err
}

	// Close the file.
	f.Close()

	// Print a success message.
	fmt.Println("Image saved as", filename)


	return nil
}
