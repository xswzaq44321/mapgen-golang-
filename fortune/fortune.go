package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	V "mapgen-golang-/voronoi"
	"os"
)

var setting Setting
var vmap V.Voronoi

func main() {
	poly := V.Polygon{
		Edges: []V.Edge{
			V.Edge{
				A: V.Point{658, 621},
				B: V.Point{514, 544},
			},
			V.Edge{
				A: V.Point{261, 109},
				B: V.Point{232, 0},
			},
			V.Edge{
				A: V.Point{232, 0},
				B: V.Point{558, 0},
			},
			V.Edge{
				A: V.Point{261, 109},
				B: V.Point{514, 544},
			},
			V.Edge{
				A: V.Point{558, 0},
				B: V.Point{658, 621},
			},
		},
		Focus: V.Point{519, 272},
	}

	fmt.Println(poly.Contain(V.Point{245, 93}))
}

func initial() {
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
