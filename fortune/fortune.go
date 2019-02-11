package main

import (
	"encoding/json"
	"io/ioutil"
	V "mapgen-golang-/voronoi"
	"os"
)

var setting Setting
var vmap V.Voronoi

func main() {
}

func init() {
	b, _ := ioutil.ReadFile("setting.json")
	json.Unmarshal(b, &setting)
	b2, _ := ioutil.ReadFile(os.Args[1])
	json.Unmarshal(b2, &vmap)
}

type Setting struct {
	PointSize int
	LineWidth int
}

type SweepLine struct {
	L            float64
	beachPolys   []V.Polygon
	siteEvents   []Event
	circleEvents []Event
	vmap         V.Voronoi
}

type Event struct {
}
