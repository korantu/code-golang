package dots

import (
	//	"bytes"
	//	"encoding/gob"
	"fmt"
	"runtime/debug"
	"testing"
)

// Must helps tag and check several things at once.
func Must(t *testing.T, reason string, deeds ...bool) {
	test_results := ""
	passed := true
	for _, outcome := range deeds {
		if outcome {
			test_results += "+"
		} else {
			test_results += "-"
			passed = false
		}
	}
	if !passed {
		t.Errorf("%s\nVerification %s failed: [%s]", string(debug.Stack()), reason, test_results)
	}
}

/*
func BuildGraph(){
	Reset()
	p := A("Person")
	f := A("Family")
	me := A("Kostya")
	jj := A("Jenia")
	mm := A("Nonna")
	pp := A("Daniil")

	levinski := A(".")
	Arrow(f, levinski)

	Arrows(levinski, Dots{ jj, me})
	Arrows( Dots{ mm, pp}, levinski)
	Arrows( p, Dots{ me, jj, mm, pp})
}

func TestGraph(t *testing.T){
	BuildGraph()
	Must(t, "find", MatchString("Person").Size() == 1, MatchString("Persssson").Size() == 0) 
}

func TestCore(t *testing.T) {
	a := A("1")
	b := A("2")
	Must(t, "creating",
		a.Means().AsString() == "1",
		b.Means().AsString() == "2",
		a.Less().Size() == 0, a.More().Size() == 0)

	Arrow(a, b)
	Arrow(a, b)

	Must(t, "arrow", a.Less().IsEmpty() && b.More().IsEmpty(),
		a.More().Size() == 1, b.Less().Size() == 1,
		a.More()[0].Means().AsString() == "2", b.Less()[0].Means().AsString() == "1")
}

type SomeState struct {
	U int
	M OneToMany
}

func SameState(a, b SomeState) bool {
	if a.U != b.U {
		return false
	}
	return a.M.Equals(b.M)
}

const (
	MagicSomething = 33
)

// CheckParamPassing modifies a map in the passed structure, to see if
// The map as cloned. It appears it was not.
func CheckParamsPassing(a SomeState) {
	a.M[MagicSomething] = Dots{MagicSomething}
}

func TestGob(t *testing.T) {
	buff := bytes.NewBuffer([]byte{})
	encoder := gob.NewEncoder(buff)

	a := SomeState{42, OneToMany{0: {42}}}
	a1 := SomeState{42, OneToMany{0: {42}}}
	b := SomeState{0, OneToMany{}}

	CheckParamsPassing(b)
	some, ok := b.M[MagicSomething];
	Must(t, "param passing", ok && some.Size() == 1 && some[0] == MagicSomething)

	b = SomeState{0, OneToMany{}} // gob will not re-init things.

	Must( t, "state checker check", SameState(a, a1) && (!SameState(a, b)))

	if err := encoder.Encode(&a); err != nil {
		t.Error(err.Error())
	}

	buff = bytes.NewBuffer(buff.Bytes())

	decoder := gob.NewDecoder(buff)
	if err := decoder.Decode(&b); err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("%v:%v", a, b)
	Must( t, "decoded", SameState(a, b))
}

func TestDots(t *testing.T) {
	dots := Dots{1, 2, 4}
	Must(t, "simple inclusion", !dots.Has(3), dots.Has(4))
	Must(t, "dots in dots", dots.Includes(Dots{1, 4}), !dots.Includes(Dots{4, 3}))
}

func TestOneToMany(t *testing.T) {
	a := OneToMany{1: {2, 3}, 4: {2, 3}}
	b := OneToMany{1: {2, 3}, 4: {2, 3}}
	a1 := OneToMany{2: {1, 4}, 3: {4, 1}}

	Must(t, "samety", a.Equals(b), !a.Equals(a1))
	Must(t, "inverse", a.Inverse().Equals(a1))
}

func TestDotsSimple ( t * testing.T){
	o := Oo{1,2}
	bad := Oo{3,2,1,2}
	good := Oo{1,2,3}
	Must(t, "sameness", o.Same(1,2), o.Same(Oo{1,2} ...), !o.Same(1,2,3))
	Must(t, "goodness", bad.Good().Same(good ...))
}
*/
// Graph

func TestGrapher( t * testing.T) {
	g := NewGrapher()
	g.Arrow(g.Dot(), g.Dot())
	g.Arrow(g.Dot(), g.Dot())
	fmt.Println(g)
	Must(t, "dots", len(g.Arrows()) == 2, g.Size() == 4)

	g = NewGrapher()
	a,b,c := g.Dot(), g.Dot(), g.Dot()
	g.Arrow(a, b)
	g.Arrow(a, c)
	Must(t, "dots", len(g.Arrows()) == 2, g.Size() == 3)
}

// Scanners, eventually sources of dots.

// SimpleNoter is a counting Noter
type SimpleNoter struct { up, down, max_depth, depth int }
func (a *SimpleNoter) Up() { a.up++; a.depth-- }
func (a *SimpleNoter) Down() { a.down++; a.depth++; if a.max_depth < a.depth { a.max_depth = a.depth} }
func (a *SimpleNoter) Token([]byte){}
func (a *SimpleNoter) Same( b* SimpleNoter) bool{
	return a.up == b.up && a.down == b.down && a.max_depth == b.max_depth
}


func TestElispScanner( t * testing.T){

	// Simplest usage
	noter := &SimpleNoter{};
	ElispScanner([]byte("()")).Scan( noter)
	fmt.Println( ">>>", noter)
	Must( t, "simplest", noter.Same( &SimpleNoter{1,1,1,0} ))

	testcases := [](struct { input string; answer *SimpleNoter }) {
		{"((()))", &SimpleNoter{3,3,3,0}},
		{"()()()", &SimpleNoter{3,3,1,0}},
		{"()();()", &SimpleNoter{2,2,1,0}},
		{`()(");\n(")`, &SimpleNoter{2,2,1,0}},
	}

	for i, test := range testcases {
		fmt.Println( "Test ", i)
		
		noter := &SimpleNoter{}
		ElispScanner([]byte(test.input)).Scan( noter)
		Must(t, "testcases", noter.Same( test.answer))
	}

}