package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"syscall/js"
	"time"

	"github.com/EdlinOrg/prominentcolor"
)

func main() {
	// 信道阻塞主协程， 否则会报错go程序已退出
	done := make(chan int, 0)
	js.Global().Set("processImage", js.FuncOf(processImage))
	<-done
}
func loadLocalImg() {
	// load local image
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

func processImage(this js.Value, args []js.Value) interface{} {
	fmt.Println("go接收时间：", time.Now().UnixNano()/1e6)
	array := args[0]
	imgSlice := make([]uint8, array.Get("byteLength").Int())
	js.CopyBytesToGo(imgSlice, array)

	reader := bytes.NewReader(imgSlice)
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Println("err:", err)
		return nil
	}
	fmt.Println("已读取到图片待处理", time.Now().UnixNano()/1e6)
	// 处理出片
	colours, err := prominentcolor.Kmeans(img)
	fmt.Println("图片处理完毕：", time.Now().UnixNano()/1e6)

	if err != nil {
		log.Fatal("filed process img:", err)
	}

	fmt.Println("colours:", colours)

	fmt.Println("Domain colors:")
	res := ""
	for i, colour := range colours {
		fmt.Println("#" + colour.AsString())
		// res[i] = "#" + colour.AsString()
		if i < len(colours)-1 {
			res += colour.AsString() + ","
		} else {
			res += colour.AsString()
		}
	}
	fmt.Println("result:", res)
	return js.ValueOf(res)
}
