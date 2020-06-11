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
	} else if option == "resize" {
		imageResize(image)
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

func imageResize(img image.Image) {
	newImage := NewDrawableImage(img)
	for i := 0; i < 1; i++ {
		energy := calculateImageEnergy(newImage)
		bestPath := getWeakestPath(energy)
		newImage = resize(bestPath, newImage)
	}

	outFile, _ := os.Create("image_resized.jpeg")
	jpeg.Encode(outFile, newImage, nil)

	fmt.Println("Image resized successfully!")
}
func calculateImageEnergy(img *image.NRGBA) [][]int {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	fmt.Printf("dimensions: %d x %d\n", width, height)

	var imageEnergy int = 0
	energy := make([][]int, height)
	for i := range energy {
		energy[i] = make([]int, width)
	}
	var left, right, up, down int
	var widthEnergy, heightEnergy int

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			left = Max(j-1, 0)
			right = Min(j+1, width-1)

			up = Max(i-1, 0)
			down = Min(i+1, height-1)

			r1, g1, b1, _ := rgbaToPixelToInt(img.At(left, i).RGBA())
			r2, g2, b2, _ := rgbaToPixelToInt(img.At(right, i).RGBA())

			r3, g3, b3, _ := rgbaToPixelToInt(img.At(j, up).RGBA())
			r4, g4, b4, _ := rgbaToPixelToInt(img.At(j, down).RGBA())

			widthEnergy = Square(r1-r2) + Square(g1-g2) + Square(b1-b2)
			heightEnergy = Square(r3-r4) + Square(g3-g4) + Square(b3-b4)

			energy[i][j] = widthEnergy + heightEnergy
			imageEnergy += energy[i][j]
		}
	}
	return energy
}
func getWeakestPath(energy [][]int) []int {
	height := len(energy)
	width := len(energy[0])

	pointers := make([][]Pointer, height)
	for i := range pointers {
		pointers[i] = make([]Pointer, width)
	}
	var previous []int = energy[0]
	var current []int
	var left int
	var right int

	for i := 0; i < height; i++ {
		current = energy[i]
		for j := 0; j < width; j++ {
			left = Max(j-1, 0)
			right = Min(j+1, width-1)

			if previous[left] < previous[right] {
				current[j] += previous[left]
				var p Pointer
				p.energy = previous[left]
				p.previousRowValue = left
				pointers[i][j] = p
			} else {
				current[j] += previous[right]
				var p Pointer
				p.energy = previous[right]
				p.previousRowValue = right
				pointers[i][j] = p
			}
		}
		previous = current
	}

	min := 0
	for i := 0; i < width; i++ {
		if pointers[height-1][i].energy < pointers[height-1][min].energy {
			min = i
		}
	}

	var list []int
	list = append(list, min)
	for i := height - 1; i >= 0; i-- {
		min = pointers[i][min].previousRowValue
		list = append(list, min)
	}

	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

	return list
}
func resize(list []int, img *image.NRGBA) *image.NRGBA {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	imageResized := image.NewNRGBA(image.Rectangle{img.Bounds().Min, image.Point{width - 1, height}})
	for i := 0; i < height; i++ {
		col := 0
		for j := 0; j < width; j++ {
			r, g, b, a := rgbaToPixel(img.At(j, i).RGBA())
			imageResized.Set(col, i, color.RGBA{r, g, b, a})
			col++
		}
	}

	return imageResized
}

type Pointer struct {
	energy           int
	previousRowValue int
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func Square(x int) int {
	return x * x
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) (uint8, uint8, uint8, uint8) {
	return uint8(r / 257), uint8(g / 257), uint8(b / 257), uint8(a / 257)
}
func rgbaToPixelToInt(r uint32, g uint32, b uint32, a uint32) (int, int, int, int) {
	return int(r / 257), int(g / 257), int(b / 257), int(a / 257)
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
