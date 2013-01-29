package funs

import (
	"fmt"
	"kdl"
	"testing"
)

func TestHelper(t *testing.T) {
	test_result := true // To test replace to false.
	kdl.Must(t, "print stack", true, true, test_result)
}

func TestContinuous(t *testing.T) {
 	a, b := 𝕍{1, 2, 3}, 𝕍{1, 2, 3}
  	kdl.Must(t, "sameness", a.Same(b), !a.Same(b.Add(a)))

	// z = [x y]
	x, y, z := 𝕍{3, -3, 1}, 𝕍{4, 9, 2}, 𝕍{-15, -2, 39}
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

func TestDiscrete(t *testing.T) {
	a, b := 𝕀{2, 3, 4}, 𝕀{6, 2, 1}
	kdl.Must(t, "sameness", a.Same(a), !a.Same(b),
		a.Same(a.Add(b).Add(b).Sub(b.Scale(2))))
	kdl.Must(t, "shift", a.Scale(4).Same(a.Shift(2)))
}

func TestMesh(t *testing.T) {
	a, b, c := &𝕍{1, 0, 0}, &𝕍{0, 1, 0}, &𝕍{0, 0, 1}

	mesh := NewMesh()

	mesh.Face(a, b, c).Face(a, c, b)
	kdl.Must(t, "one triangle", len(mesh.Points()) == 3, len(mesh.Faces()) == 2)
}
