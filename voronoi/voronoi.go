package voronoi

import (
	"math"
	"sort"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (a Point) Distance(b Point) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}

func vector(a Point, b Point) Point {
	return Point{b.X - a.X, b.Y - a.Y}
}

func cross(o Point, a Point, b Point) float64 {
	v1, v2 := vector(o, a), vector(o, b)
	return v1.X*v2.Y - v1.Y*v2.X
}

func midPoint(points []Point) Point {
	var Sx, Sy float64
	for _, v := range points {
		Sx += v.X
		Sy += v.Y
	}
	return Point{Sx / float64(len(points)), Sy / float64(len(points))}
}

type Edge struct {
	A        Point `json:"a"`
	B        Point `json:"b"`
	ParentID []int `json:"parentID,omitempty"`
}

// get the n'th point of that edge
// index of greater than 2 will return last point
func (e Edge) Get(index int) Point {
	if index == 0 {
		return e.A
	} else {
		return e.B
	}
}

// return distance fron point p to line e
func (e Edge) Distance(p Point) float64 {
	m := (e.B.Y - e.A.Y) / (e.B.X - e.A.X)
	a, b, c := m, float64(-1), -m*e.A.X+e.A.Y
	return math.Abs(a*p.X+b*p.Y+c) / math.Sqrt(a*a+b*b)
}

type Polygon struct {
	Edges     []Edge `json:"edges,omitempty"`
	Focus     Point  `json:"focus"`
	Id        int    `json:"id,omitempty"`
	organized bool
}

type foo struct {
	key float64
	val Edge
}

type ByKey []foo

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].key < a[j].key }

// sort edges to counterclockwise order
func (poly *Polygon) organize() {
	arr := make([]foo, len(poly.Edges))
	for k, v := range poly.Edges {
		mid := midPoint([]Point{v.A, v.B})
		vec := vector(poly.Focus, mid)
		angle := math.Atan2(vec.Y, vec.X)
		arr[k] = foo{key: angle, val: v}
	}
	sort.Sort(ByKey(arr))
	for k, v := range arr {
		poly.Edges[k] = v.val
	}
	for k, _ := range poly.Edges {
		if cross(poly.Focus, poly.Edges[k].A, poly.Edges[k].B) < 0 {
			poly.Edges[k].A, poly.Edges[k].B = poly.Edges[k].B, poly.Edges[k].A
		}
	}
	poly.organized = true
}

func (poly *Polygon) Contain(p Point) bool {
	n := len(poly.Edges)
	if !poly.organized {
		poly.organize()
	}
	points := make([]Point, n)
	for k, v := range poly.Edges {
		points[k] = v.A
	}
	if len(points) < 3 {
		return false
	}
	if cross(points[0], p, points[1]) > 0 {
		return false
	}
	if cross(points[0], p, points[n-1]) < 0 {
		return false
	}

	l, r := 2, n-1
	line := -1
	for l <= r {
		mid := (l + r) / 2
		if cross(points[0], p, points[mid]) > 0 {
			line = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return cross(points[line-1], p, points[line]) < 0
}

type Voronoi struct {
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Polygons []Polygon `json:"polygons,omitempty"`
}
