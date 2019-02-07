package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	V "mapgen-golang-/voronoi"
	"math/rand"
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
				Focus: V.Point{x, y},
				Edges: make([]V.Edge, 0),
			}
			voronoi.Polygons = append(voronoi.Polygons, bar)
		} else {
			i--
		}
	}

	output(voronoi)
}

func output(m V.Voronoi) {
	outJson, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s", outJson)
}

func check(polys []V.Polygon, x int, y int) bool {
	for _, v := range polys {
		if v.Focus.Distance(V.Point{x, y}) < float64(setting.PointMargin) {
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
}
