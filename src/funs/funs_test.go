package funs

import (
	"bytes"
	"fmt"
	"kdl"
	"testing"
	"text/template"
)

func TestHelper(t *testing.T) {
	test_result := true // To test replace to false.
	kdl.Must(t, "print stack", true, true, test_result)
}

func TestContinuous(t *testing.T) {
	a, b := Vector{1, 2, 3}, Vector{1, 2, 3}
	kdl.Must(t, "sameness", a.Same(b), !a.Same(b.Add(a)))

	// z = [x y]
	x, y, z := Vector{3, -3, 1}, Vector{4, 9, 2}, Vector{-15, -2, 39}
	fmt.Println(x.Cross(y))
	kdl.Must(t, "cross", z.Same(x.Cross(y)))
	kdl.Must(t, "dot",
		x.Dot(z).Same(0),
		z.Dot(x).Same(0),
		z.Len2().Same(z.Dot(z)))

	// Scaling etc.
	kdl.Must(t, "normalize",
		a.Scale(2.3).Divide(2.3).Same(a),
		a.Norm().Len().Same(1),
		a.Dot(a.Norm()).Same(a.Len()),
		a.Cross(a.Norm()).Len().Same(0),
		a.Norm().Scale(-1).Sub(a.Norm()).Len().Same(2))
}

func TestMesh(t *testing.T) {
	a, b, c := &Vector{1, 0, 0}, &Vector{0, 1, 0}, &Vector{0, 0, 1}

	mesh := NewMesh()

	mesh.Face(a, b, c).Face(a, c, b)
	kdl.Must(t, "one triangle", len(mesh.Points()) == 3, len(mesh.Faces()) == 2)

	mesh = NewMesh()
	Vector{0, 0, 0}.Dump(mesh)
	kdl.Must(t, "dumped dot", len(mesh.Points()) == 6, len(mesh.Faces()) == 8)
}

func TestTemplates(t *testing.T) {
	simplest := template.New("Simple")
	a, _ := simplest.Parse("a")
	buf := &bytes.Buffer{}
	a.Execute(buf, nil)
	kdl.Must(t, "simplest", string(buf.Bytes()) == "a")

	
	simpler := template.New("Range")
	a, err := simpler.Parse("{{range .}}[{{.}}]{{end}}")
	if nil != err {
		print(err.Error())
		t.Fail()
	} else {
		buf.Reset()
		a.Execute(buf, []int{1, 2, 8})
		kdl.Must(t, "iterate", string(buf.Bytes()) == "[1][2][8]")
	}
}
