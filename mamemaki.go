package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"time"
)

var setting Setting

func main() {
	var m []Point

	b, _ := ioutil.ReadFile("setting.json")
	json.Unmarshal(b, &setting)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < setting.PointCount; i++ {
		x := rand.Intn(setting.Width-2*setting.MapPadding) + setting.MapPadding
		y := rand.Intn(setting.Height-2*setting.MapPadding) + setting.MapPadding
		if check(m, x, y) {
			m = append(m, Point{x, y})
		} else {
			i--
		}
	}

	output(m)
}

func check(m []Point, x int, y int) bool {
	for _, v := range m {
		if distance(v, Point{x, y}) < float64(setting.PointMargin) {
			return false
		}
	}
	return true
}

func distance(a Point, b Point) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
}

func output(m []Point) {
	outJson, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s", outJson)
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Setting struct {
	Width       int
	Height      int
	PointCount  int
	PointMargin int
	MapPadding  int
}
