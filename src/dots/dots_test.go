package dots

import (
	"fmt"
	"kdl"
	"testing"
)

// Graph

func TestGrapher(t *testing.T) {
	g := NewGrapher()
	g.Arrow(g.Dot(), g.Dot())
	g.Arrow(g.Dot(), g.Dot())
	fmt.Println(g)
	kdl.Must(t, "dots", len(g.Arrows()) == 2, g.Size() == 4)

	g = NewGrapher()
	a, b, c := g.Dot(), g.Dot(), g.Dot()
	g.Arrow(a, b)
	g.Arrow(a, c)
	kdl.Must(t, "dots", len(g.Arrows()) == 2, g.Size() == 3)
}

// Scanners, eventually sources of dots.

// SimpleNoter is a counting Noter
type SimpleNoter struct{ up, down, max_depth, depth int }

func (a *SimpleNoter) Up() { a.up++; a.depth-- }
func (a *SimpleNoter) Down() {
	a.down++
	a.depth++
	if a.max_depth < a.depth {
		a.max_depth = a.depth
	}
}
func (a *SimpleNoter) Token([]byte) {}
func (a *SimpleNoter) Same(b *SimpleNoter) bool {
	return a.up == b.up && a.down == b.down && a.max_depth == b.max_depth
}

func TestElispScanner(t *testing.T) {

	// Simplest usage
	noter := &SimpleNoter{}
	ElispScanner([]byte("()")).Scan(noter)
	fmt.Println(">>>", noter)
	kdl.Must(t, "simplest", noter.Same(&SimpleNoter{1, 1, 1, 0}))

	testcases := [](struct {
		input  string
		answer *SimpleNoter
	}){
		{"((()))", &SimpleNoter{3, 3, 3, 0}},
		{"()()()", &SimpleNoter{3, 3, 1, 0}},
		{"()();()", &SimpleNoter{2, 2, 1, 0}},
		{`()(");\n(")`, &SimpleNoter{2, 2, 1, 0}},
	}

	for i, test := range testcases {
		fmt.Println("Test ", i)

		noter := &SimpleNoter{}
		ElispScanner([]byte(test.input)).Scan(noter)
		kdl.Must(t, "testcases", noter.Same(test.answer))
	}

}
