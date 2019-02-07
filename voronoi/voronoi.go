package voronoi

import "math"

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (a Point) Distance(b Point) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
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
