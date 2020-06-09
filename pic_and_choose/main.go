package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello world from pic_and_choose")
	if len(os.Args) < 2 {
		log.Fatal("Invalid script call. Should pass an image path as argument")
		return
	}

	fmt.Println("Image to process: " + os.Args[1])
	image, err := getImageFromFilePath(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("Choose an option (\"blur\", \"exit\"): ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		if option == "blur" {
			blurImage(image, 3)
		} else if option == "exit" {
			fmt.Println("Exiting...")
			return
		} else {
			fmt.Println("Option not supported yet!")
		}
	}

}

func blurImage(image image.Image, blurFactor int) {
	var factor uint16 = uint16(blurFactor)
	var factorSquared uint16 = factor * factor

	width := image.Bounds().Max.X
	height := image.Bounds().Max.Y

	fmt.Println("Width: " + strconv.Itoa(width))
	fmt.Println("Height: " + strconv.Itoa(height))

	newImage := NewMyImg(image)

	for i := 0; i < height; i += int(factor) {
		for j := 0; j < width; j += int(factor) {
			var r uint16 = 0
			var g uint16 = 0
			var b uint16 = 0
			var a uint16 = 0
			for k := i; k < int(factor)+i && k < height; k++ {
				for l := j; l < int(factor)+j && l < width; l++ {
					red, green, blue, alpha := rgbaToPixel(image.At(l, k).RGBA())
					r += uint16(red)
					g += uint16(green)
					b += uint16(blue)
					a += uint16(alpha)
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

	// Save new image
	outFile, _ := os.Create("changed.jpeg")
	jpeg.Encode(outFile, newImage, nil)

	fmt.Println("Image blurred successfully!")
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) (uint8, uint8, uint8, uint8) {
	return uint8(r / 257), uint8(g / 257), uint8(b / 257), uint8(a / 257)
}

type MyImg struct {
	image.Image
	custom map[image.Point]color.Color
}

func NewMyImg(img image.Image) *MyImg {
	return &MyImg{img, map[image.Point]color.Color{}}
}

func (m *MyImg) Set(x, y int, c color.Color) {
	m.custom[image.Point{x, y}] = c
}

func (m *MyImg) At(x, y int) color.Color {
	// Explicitly changed part: custom colors of the changed pixels:
	if c := m.custom[image.Point{x, y}]; c != nil {
		return c
	}
	// Unchanged part: colors of the original image:
	return m.Image.At(x, y)
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(f)
	return image, err
}
