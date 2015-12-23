package main

import (
	"fmt"
	"image"
	"image/png"
	"image/draw"
	"image/color"
	"os"
)

func createDestImage(fileName string, img image.Image ) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	return png.Encode(file, img)
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
	src := loadSrcImage("d:\\xxwn.png")
	mask := loadMaskImage("d:\\mask.png")

	srcBounds := src.Bounds()
	maskBounds := mask.Bounds()

	println(src.Bounds().String());
	println(mask.Bounds().String());

	white := image.Uniform{ color.RGBA{255, 255, 255, 255} }
	whiteImg := image.NewRGBA(maskBounds)
	draw.Draw(whiteImg, maskBounds, &white, image.ZP, draw.Over)

	copyPoint := image.Pt(80, 80)

	// Create a new image with src bounds for final background
	copy1 := copySrc(src)
	// Get the src image without mask-bounds-region
	draw.DrawMask(copy1, srcBounds.Add(copyPoint), whiteImg, image.ZP, mask, maskBounds.Min, draw.Over)
	createDestImage("d:\\background.png", copy1)

	// Create a new image with mask bounds for final move block
	copy2 := copySrc(mask)
	// Get the part image in src image with the mask-bounds
	draw.DrawMask(copy2, maskBounds, src, copyPoint, mask, maskBounds.Min, draw.Over)
	createDestImage("d:\\moveblock.png", copy2)

	println("Generated");
}

func main() {
	GetSrcWithMaskOpImages()
}