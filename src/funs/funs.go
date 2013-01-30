package funs

/*
 Rule: No unneded incomplete features.
 Everything must have a reason to be here in this particular way.
*/

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
type Vector struct {
	x, y, z R
}
// Niceties

func (a Vector) String() string { return fmt.Sprintf("(%.2f,%.2f,%.2f)", a.x, a.y, a.z) }

func (a Vector) Same(b Vector) bool { return a.Sub(b).Len() < ε }

func (a R) Same(b R) bool { return math.Abs(float64(a-b)) < ε }


// Operations

func (a Vector) Add(b Vector) Vector { return Vector{a.x + b.x, a.y + b.y, a.z + b.z} }
func (a Vector) Sub(b Vector) Vector { return Vector{a.x - b.x, a.y - b.y, a.z - b.z} }

func (a Vector) Scale(c R) Vector  { return Vector{a.x * c, a.y * c, a.z * c} }
func (a Vector) Divide(c R) Vector { return Vector{a.x / c, a.y / c, a.z / c} }

func (a Vector) Len2() R      { return a.x*a.x + a.y*a.y + a.z*a.z }
func (a Vector) Len() R       { return R(math.Sqrt(float64(a.Len2()))) }
func (a Vector) Norm() Vector { l := a.Len(); return Vector{a.x / l, a.y / l, a.z / l} }

func (a Vector) Dot(b Vector) R { return a.x*b.x + a.y*b.y + a.z*b.z }
func (a Vector) Cross(b Vector) Vector {
	return Vector{a.y*b.z - a.z*b.y, -(a.x*b.z - a.z*b.x), a.x*b.y - a.y*b.x}
}

var X = Vector{1, 0, 0}
var Y = Vector{0, 1, 0}
var Z = Vector{0, 0, 1}

// Dumping

func (a Vector) Dump(f FacesWriter) {
	x := []Vector{a.Add(X), a.Sub(X)}
	y := []Vector{a.Add(Y), a.Sub(Y)}
	z := []Vector{a.Add(Z), a.Sub(Z)}
	f.
		Face(&x[0], &y[0], &z[0]).
		Face(&x[0], &z[0], &y[1]).
		Face(&x[0], &y[1], &z[1]).
		Face(&x[0], &z[1], &y[0]).
		Face(&x[1], &y[0], &z[0]).
		Face(&x[1], &z[0], &y[1]).
		Face(&x[1], &y[1], &z[1]).
		Face(&x[1], &z[1], &y[0])

}

// FaceWriter can be used to write mesh faces.
type FacesWriter interface {
	Face(a ...*Vector) FacesWriter
}

// FacesReader allows to read points and corresponfing face indices.
type FacesReader interface {
	Points() []Vector
	Faces() [][]I
}

type FacesWriterReader interface {
	FacesReader
	FacesWriter
}

// Takes at least twice as much memory as needed. Well, basic...
type basic_indexed_mesh struct {
	set    map[*Vector]I
	points []Vector
	faces  [][]I
}

func (a *basic_indexed_mesh) Face(pts ...*Vector) FacesWriter {
	face := []I{}

	// Check if it was seen before, and add if not.
	index_of := func(pt *Vector) (the_index I) {
		if existing_idx, ok := a.set[pt]; ok {
			the_index = existing_idx
		} else {
			new_idx := I(len(a.points))
			a.set[pt] = new_idx
			a.points = append(a.points, *pt)
			the_index = new_idx
		}
		return
	}

	// Loop over the face
	for _, pt := range pts {
		face = append(face, index_of(pt))
	}

	// Append to all the faces
	a.faces = append(a.faces, face)
	return a
}

func (a *basic_indexed_mesh) Points() []Vector { return a.points }
func (a *basic_indexed_mesh) Faces() [][]I     { return a.faces }

func NewMesh() FacesWriterReader { return &basic_indexed_mesh{map[*Vector]I{}, []Vector{}, [][]I{}} }
