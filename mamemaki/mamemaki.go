package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	V "mapgen-golang-/voronoi"
	"math"
	"math/rand"
	"os"
	"time"
)

var setting Setting

func main() {
	b, _ := ioutil.ReadFile("setting.json")
	json.Unmarshal(b, &setting)
	rand.Seed(time.Now().UnixNano())

	var voronoi V.Voronoi
	voronoi.Width = setting.Width
	voronoi.Height = setting.Height

	for i := 0; i < setting.PointCount; i++ {
		x := rand.Intn(setting.Width-2*setting.MapPadding) + setting.MapPadding
		y := rand.Intn(setting.Height-2*setting.MapPadding) + setting.MapPadding
		if check(voronoi.Polygons, x, y) {
			bar := V.Polygon{
				Focus: V.Point{X: x, Y: y},
				Edges: make([]V.Edge, 0),
			}
			voronoi.Polygons = append(voronoi.Polygons, bar)
		} else {
			i--
		}
	}

	imgOutput(voronoi)
}

func output(m V.Voronoi) {
	outJson, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s", outJson)
}

func imgOutput(m V.Voronoi) {
	myImage := image.NewRGBA(image.Rect(0, 0, setting.Width, setting.Height))

	outputFile, err := os.Create("result.png")
	defer func() {
		outputFile.Close()
	}()

	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	for i := 0; i < setting.Width; i++ {
		for j := 0; j < setting.Height; j++ {
			myImage.Set(i, j, color.White)
		}
	}
	for _, p := range m.Polygons {
		fillCircle(myImage, p.Focus.X, p.Focus.Y, setting.PointSize, color.RGBA{255, 0, 0, 255})
		//myImage.Set(p.Focus.X, p.Focus.Y, color.RGBA{255, 0, 0, 255})
	}

	png.Encode(outputFile, myImage)
}

func fillCircle(img *image.RGBA, x0, y0, R int, c color.RGBA) {
	r := float64(R) / 2
	err := 0.0
	if R%2 == 0 {
		err = 0.5
	}
	for x := math.Floor(r) - err; x >= -math.Floor(r); x-- {
		// round half down y = ceil(x-0.5)
		for y := math.Ceil(math.Sqrt(r*r-x*x)-0.5) - err; y >= -math.Ceil(math.Sqrt(r*r-x*x)-0.5); y-- {
			dx := x0 + int(x+err)
			dy := y0 + int(y+err)
			if dx < 0 || dx >= img.Rect.Dx() || dy < 0 || dy >= img.Rect.Dy() {
				continue
			} else {
				img.Set(dx, dy, c)
			}
		}
	}
}

func check(polys []V.Polygon, x int, y int) bool {
	for _, v := range polys {
		if v.Focus.Distance(V.Point{X: x, Y: y}) < float64(setting.PointMargin) {
			return false
		}
	}
	return true
}

type Setting struct {
	Width       int
	Height      int
	PointCount  int
	PointMargin int
	MapPadding  int
	PointSize   int
}
