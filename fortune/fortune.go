package main

import (
	"encoding/json"
	"io/ioutil"
	"mapgen-golang-/BST"
	V "mapgen-golang-/voronoi"
	"os"
)

var setting Setting
var vmap V.Voronoi

type Bar struct {
	data int
}

func (a Bar) LessThan(c BST.Comparable) bool {
	return a.data < c.(Bar).data
}
func (a Bar) GreaterThan(c BST.Comparable) bool {
	return a.data > c.(Bar).data
}
func (a Bar) EqualTo(c BST.Comparable) bool {
	return a.data == c.(Bar).data
}

func main() {
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
	beachPolys   []*V.Polygon
	siteEvents   []*Event
	circleEvents []*Event
	vmap         *V.Voronoi
}

type ByX []Event

func (a ByX) Len() int           { return len(a) }
func (a ByX) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByX) Less(i, j int) bool { return a[i].X < a[j].X }

type Event struct {
	relevant []*V.Polygon
	X        float64
	center   V.Point
	isCircle bool
}

func det(m [3][3]float64) float64 {
	return (m[0][0]*m[1][1]*m[2][2] + m[0][1]*m[1][2]*m[2][0] + m[0][2]*m[1][0]*m[2][1]) - (m[0][2]*m[1][1]*m[2][0] + m[0][1]*m[1][0]*m[2][2] + m[0][0]*m[1][2]*m[2][1])
}
