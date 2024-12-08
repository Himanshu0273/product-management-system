package processor

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "image"
    "image/jpeg"
    "os"
    "github.com/nfnt/resize"
)

func ProcessImage(imageURL string) {
    fmt.Println("Downloading image:", imageURL)

    // Download the image
    response, err := http.Get(imageURL)
    if err != nil {
        log.Println("Failed to download image:", err)
        return
    }
    defer response.Body.Close()

    imgData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Println("Failed to read image:", err)
        return
    }

    // Decode the image
    img, _, err := image.Decode(bytes.NewReader(imgData))
    if err != nil {
        log.Println("Failed to decode image:", err)
        return
    }

    // Resize the image
    resizedImg := resize.Resize(100, 0, img, resize.Lanczos3)

    // Save the processed image
    outFile, err := os.Create("processed_image.jpg")
    if err != nil {
        log.Println("Failed to create file:", err)
        return
    }
    defer outFile.Close()

    err = jpeg.Encode(outFile, resizedImg, nil)
    if err != nil {
        log.Println("Failed to save image:", err)
        return
    }

    fmt.Println("Processed image saved as processed_image.jpg")
}
