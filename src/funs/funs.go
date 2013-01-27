package funs

import (
	"fmt"
	"math"
)

// Data structures
type R float32
type I uint

const (
	ε = 0.0001
)

// Smooth point
type 𝕍 struct {
	x, y, z R
}

// Discrete point
type 𝕀 struct {
	x, y, z I
}

// Cell
type 𝕆 struct {
	x, y, z, r I
}

// Niceties

func (a 𝕍) String() string { return fmt.Sprintf("(%.2f,%.2f,%.2f)", a.x, a.y, a.z) }

func (a 𝕍) Same(b 𝕍) bool { return a.Sub(b).Len() < ε }
func (a 𝕀) Same(b 𝕀) bool { return a.x == b.x && a.y == b.y && a.z == b.z }
func (a 𝕆) Same(b 𝕆) bool { return a.x == b.x && a.y == b.y && a.z == b.z && a.r == b.r }

func (a R) Same(b R) bool { return math.Abs(float64(a-b)) < ε }

// Operations

func (a 𝕍) Add(b 𝕍) 𝕍 { return 𝕍{a.x + b.x, a.y + b.y, a.z + b.z} }
func (a 𝕍) Sub(b 𝕍) 𝕍 { return 𝕍{a.x - b.x, a.y - b.y, a.z - b.z} }
func (a 𝕀) Add(b 𝕀) 𝕀 { return 𝕀{a.x + b.x, a.y + b.y, a.z + b.z} }
func (a 𝕀) Sub(b 𝕀) 𝕀 { return 𝕀{a.x - b.x, a.y - b.y, a.z - b.z} }

func (a 𝕍) Scale(c R) 𝕍 { return 𝕍{a.x * c, a.y * c, a.z * c} }
func (a 𝕍) Divide(c R) 𝕍 { return 𝕍{a.x / c, a.y / c, a.z / c} }
func (a 𝕀) Scale(c I) 𝕀 { return 𝕀{a.x * c, a.y * c, a.z * c} }
func (a 𝕀) Shift(c I) 𝕀 { return 𝕀{a.x << c, a.y << c, a.z << c} }

func (a 𝕍) Len2() R { return a.x*a.x + a.y*a.y + a.z*a.z }
func (a 𝕍) Len() R  { return R(math.Sqrt(float64(a.Len2()))) }
func (a 𝕍) Norm() 𝕍 { l := a.Len(); return 𝕍{a.x / l, a.y / l, a.z / l} }

func (a 𝕍) Dot(b 𝕍) R   { return a.x*b.x + a.y*b.y + a.z*b.z }
func (a 𝕍) Cross(b 𝕍) 𝕍 { return 𝕍{a.y*b.z - a.z*b.y, -(a.x*b.z - a.z*b.x), a.x*b.y - a.y*b.x} }

// Dumping

// FaceWriter can be used to write mesh faces.
type FacesWriter interface {
	Face(a ... *𝕍) FacesWriter
}

// FacesReader allows to read points and corresponfing face indices.
type FacesReader interface {
	Points() []𝕍
	Faces() [][]I
}


type FacesWriterReader interface {
	FacesReader
	FacesWriter
}

// Takes at least twice as much memory as needed. Well, basic...
type basic_indexed_mesh struct {
	set map[*𝕍] I
	points []𝕍
	faces [][]I
}

func (a * basic_indexed_mesh) Face( pts ... *𝕍) FacesWriter { 
	face := []I{}

	// Check if it was seen before, and add if not.
	index_of := func(pt *𝕍) ( the_index I) {
		if existing_idx, ok := a.set[pt]; ok {
			the_index = existing_idx
		} else {
			new_idx := I(len( a.points))
			a.set[pt] = new_idx
			a.points = append( a.points, *pt)
			the_index = new_idx
		}
		return
	}

	// Loop over the face
	for _, pt := range pts {
		face = append( face, index_of( pt))
	}

	// Append to all the faces
	a.faces = append( a.faces, face)
	return a 
}

func (a * basic_indexed_mesh) Points() []𝕍 { return a.points }
func (a * basic_indexed_mesh) Faces() [][]I { return a.faces }

func NewMesh() FacesWriterReader { return &basic_indexed_mesh{ map[*𝕍]I{}, []𝕍{}, [][]I{} }}
