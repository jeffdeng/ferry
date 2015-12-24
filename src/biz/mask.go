package main

import (
	"fmt"
	"image"
	"image/png"
	"image/draw"
	"image/color"
	"os"
	"time"
	"io"
)

func createImageFile(fileName string, img image.Image ) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	return png.Encode(file, img)
}

func sendImage(w io.Writer, img image.Image) error {
	return png.Encode(w, img)
}

func loadPngImage(fileName string) image.Image {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	img, _ := png.Decode(file)

	return img
}

func loadMaskImage(fileName string) image.Image {
	return loadPngImage(fileName)
}

func loadSrcImage(fileName string) image.Image {
	return loadPngImage(fileName)
}

func copySrc(src image.Image) draw.Image {
	srcBounds := src.Bounds()
	copy := image.NewRGBA(srcBounds)
	draw.Draw(copy, srcBounds, src, image.ZP, draw.Over)
	return copy
}

func GetSrcWithMaskOpImages() {
	a := time.Now().Nanosecond()
	src := loadSrcImage("d:\\xxwn.png")
	mask := loadMaskImage("d:\\mask.png")

	b := time.Now().Nanosecond()
	copy1, copy2, _ := GetImages(src, mask, image.Pt(80, 80));
	c := time.Now().Nanosecond()



	createImageFile("d:\\background.png", copy1)
	createImageFile("d:\\moveblock.png", copy2)

	d := time.Now().Nanosecond()

	println("Generated and use ", (c - b) / 1000000, "ms", (d - a) / 1000000, "ms");
}

func GetImages(src, mask image.Image, copyPoint image.Point) (draw.Image, draw.Image, error) {
	srcBounds := src.Bounds()
	maskBounds := mask.Bounds()

	white := image.Uniform{ color.RGBA{255, 255, 255, 255} }
	whiteImg := image.NewRGBA(maskBounds)
	draw.Draw(whiteImg, maskBounds, &white, image.ZP, draw.Over)

	// Create a new image with src bounds for final background
	copy1 := copySrc(src)
	// Get the src image without mask-bounds-region
	draw.DrawMask(copy1, srcBounds.Add(copyPoint), whiteImg, image.ZP, mask, maskBounds.Min, draw.Over)

	// Create a new image with mask bounds for final move block
	copy2 := copySrc(mask)
	// Get the part image in src image with the mask-bounds
	draw.DrawMask(copy2, maskBounds, src, copyPoint, mask, maskBounds.Min, draw.Over)

	return copy1, copy2, nil
}

func main() {
	GetSrcWithMaskOpImages()
}