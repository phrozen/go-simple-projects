package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

var imagePath string
var option string
var blurFactor int
var help bool

func init() {
	flag.StringVar(&imagePath, "f", "nature.jpeg", "Image path")
	flag.StringVar(&option, "o", "blur", "image processing operation option")
	flag.IntVar(&blurFactor, "b", 0, "Blur factor for image")
	flag.BoolVar(&help, "h", false, "Help for running this program")
	flag.Parse()
}

func main() {
	if help == true {
		flag.Usage()
	}

	image, err := getImageFromFilePath(imagePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	if option == "blur" {
		blurImage(image, blurFactor)
	} else if option == "rotate" {
		rotateImage(image)
	} else {
		fmt.Println("Option not supported yet! Get help by adding this flag: -help")
	}
}

func blurImage(img image.Image, blurFactor int) {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	if blurFactor == 0 {
		blurFactor = (width + height) / 100
	}
	fmt.Printf("blurFactor=%d\n", blurFactor)
	fmt.Printf("dimensions: %d x %d\n", width, height)

	var factor uint32 = uint32(blurFactor)
	var factorSquared uint32 = factor * factor

	newImage := NewDrawableImage(img)

	for i := 0; i < height; i += int(factor) {
		for j := 0; j < width; j += int(factor) {
			var r uint32 = 0
			var g uint32 = 0
			var b uint32 = 0
			var a uint32 = 0
			for k := i; k < int(factor)+i && k < height; k++ {
				for l := j; l < int(factor)+j && l < width; l++ {
					red, green, blue, alpha := rgbaToPixel(img.At(l, k).RGBA())
					r += uint32(red)
					g += uint32(green)
					b += uint32(blue)
					a += uint32(alpha)
				}
			}
			r /= factorSquared
			g /= factorSquared
			b /= factorSquared
			a /= factorSquared

			for k := i; k < int(factor)+i && k < height; k++ {
				for l := j; l < int(factor)+j && l < width; l++ {
					newRgb := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
					newImage.Set(l, k, newRgb)
				}
			}
		}
	}

	outFile, _ := os.Create("image_blurred.jpeg")
	jpeg.Encode(outFile, newImage, nil)

	fmt.Println("Image blurred successfully!")
}

func rotateImage(img image.Image) {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	fmt.Printf("dimensions: %d x %d\n", width, height)

	imageHolder := image.NewNRGBA(image.Rectangle{img.Bounds().Min, image.Point{height, width}})

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			red, green, blue, alpha := rgbaToPixel(img.At(j, i).RGBA())
			imageHolder.Set(i, j, color.RGBA{red, green, blue, alpha})
		}
	}

	outFile, _ := os.Create("image_rotated.jpeg")
	jpeg.Encode(outFile, imageHolder, nil)

	fmt.Println("Image rotated successfully!")
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) (uint8, uint8, uint8, uint8) {
	return uint8(r / 257), uint8(g / 257), uint8(b / 257), uint8(a / 257)
}

func NewDrawableImage(img image.Image) *image.NRGBA {
	return image.NewNRGBA(img.Bounds())
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(f)
	return image, err
}
