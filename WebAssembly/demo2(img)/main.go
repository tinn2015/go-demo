package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/EdlinOrg/prominentcolor"
)

func main() {
	// load image
	// path, _ := os.Getwd()
	path, _ := filepath.Abs("asset/ff.jpg")
	fmt.Println("path", path)
	img, err := lodaImage(path)
	if err != nil {
		log.Fatal("[filed to load image] ", err)
	}
	fmt.Println("img.Bounds", img.Bounds().Max)

	// 处理出片
	colours, err := prominentcolor.Kmeans(img)

	if err != nil {
		log.Fatal("filed process img:", err)
	}

	fmt.Println("colours:", colours)

	fmt.Println("Domain colors:")
	for _, colour := range colours {
		fmt.Println("#" + colour.AsString())
	}
}

func lodaImage(fileInput string) (image.Image, error) {
	f, err := os.Open(fileInput)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	return img, err
}
