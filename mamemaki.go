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
	b, _ := ioutil.ReadFile("setting.json")
	json.Unmarshal(b, &setting)
	rand.Seed(time.Now().UnixNano())

	var voronoi Voronoi
	voronoi.Width = setting.Width
	voronoi.Height = setting.Height

	for i := 0; i < setting.PointCount; i++ {
		x := rand.Intn(setting.Width-2*setting.MapPadding) + setting.MapPadding
		y := rand.Intn(setting.Height-2*setting.MapPadding) + setting.MapPadding
		if check(voronoi.Polygons, x, y) {
			bar := Polygon{
				Focus: Point{x, y},
				Edges: make([]Edge, 0),
			}
			voronoi.Polygons = append(voronoi.Polygons, bar)
		} else {
			i--
		}
	}

	output(voronoi)
}

func output(m Voronoi) {
	outJson, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s", outJson)
}

func check(polys []Polygon, x int, y int) bool {
	for _, v := range polys {
		if v.Focus.distance(Point{x, y}) < float64(setting.PointMargin) {
			return false
		}
	}
	return true
}

func (a Point) distance(b Point) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Edge struct {
	A          Point `json:"a"`
	B          Point `json:"b"`
	ParentID   []int `json:"parentID,omitempty"`
	isAbstract bool
}

type Polygon struct {
	Edges []Edge `json:"edges,omitempty"`
	Focus Point  `json:"focus"`
	Id    int    `json:"id,omitempty"`
}

type Voronoi struct {
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Polygons []Polygon `json:"polygons,omitempty"`
}

type Setting struct {
	Width       int
	Height      int
	PointCount  int
	PointMargin int
	MapPadding  int
}
