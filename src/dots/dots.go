package dots

/*

 Plans:

 1. Simple connections, store, edit.
 2. Now the problem is - how, to, you know, plan it. When it is small and cute, it is easy to project huge capabilities, think about novel features, etc.
 3. Let's continue.

*/

import (
	// "io"
	"errors"
	// "sort"
)

var Error = errors.New("General failure.")

/*
 1. Dots - topology
 2. Soul - content
*/
/*
// Simplest.
type Soul string

type Dot uint
type Dots []Dot
type OneToMany map[Dot]Dots
type Souls []*Soul

// Topology
var more = OneToMany{}
var less = OneToMany{}

// Content
var meaning = Souls{}
var nonsense = Soul("empty")

func Reset() {
	more = OneToMany{}
	less = OneToMany{}
	meaning = Souls{}
}

func MatchString(match string) Dots {
	result := Dots{}
	for dot, soul := range meaning {
		if soul.AsString() == match {
			result = result.Add(Dot(dot))
		}
	}
	return result
}

//OneToMany
func (a OneToMany) Inverse() OneToMany {
	result := OneToMany{}
	for dot, dots := range a {
		for _, each_dot := range dots {
			if old_dots, ok := result[each_dot]; ok {
				result[each_dot] = old_dots.Add(dot)
			} else {
				result[each_dot] = Dots{dot}
			}
		}
	}
	return result
}

// Serialization

const (
	CodeStatePackSimple = iota
)

type StatePackSimple struct {
	Meaning Souls
	More    OneToMany
}

func State() *StatePackSimple {
	out := StatePackSimple{Souls{}, OneToMany{}}
	return &out
}

func (data *StatePackSimple) Pack() {
	data.Meaning = meaning
	data.More = more
}

func (data *StatePackSimple) Unpack() {
	Reset()
	meaning = data.Meaning
	more = data.More
	less = more.Inverse()
}
*/
/*
type Save( io.Writer ) error {
}

type Load( io.Reader ) error {
}
*/
/*
// AsString returns string representation of a soul.
func (a *Soul) AsString() string {
	return string(*a)
}

// A creates a piece of content
func A(some string) Dot {
	item := Soul(some)
	me := len(meaning)
	meaning = append(meaning, &item)
	return Dot(me)
}

// Means returns what the dot represents
func (a Dot) Means() *Soul {
	idea := &nonsense
	if int(a) < len(meaning) {
		idea = meaning[a]
	}
	return idea
}

// Get obtains connected dots from the mapping.
func (mapping OneToMany) Get(a Dot) Dots {
	if all, ok := mapping[a]; ok {
		return all
	}
	return Dots{}
}

func (b OneToMany) Includes(a OneToMany) bool {
	for dot_a, dots_a := range a {
		if dots_b, ok := b[dot_a]; ok {
			if !dots_b.Equals(dots_a) {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func (one OneToMany) Equals(another OneToMany) bool {
	return one.Includes(another) && another.Includes(one)
}

// --- Dots ---

func (list Dots) AsDots() Dots {
	return list
}

// Has checks if such a dot is in the list.
func (list Dots) Has(a Dot) bool {
	for _, v := range list {
		if a == v {
			return true // We got it.
		}
	}
	return false
}

func (list Dots) Includes(some Dots) bool {
	for _, a_dot := range some {
		if !list.Has(a_dot) {
			return false
		}
	}
	return true
}

func (list Dots) Equals(some Dots) bool {
	return list.Includes(some) && some.Includes(list)
}

func (list Dots) Size() int     { return len(list) }
func (list Dots) IsEmpty() bool { return list.Size() == 0 }

// Add a dot to the list if it is not already there.
func (list Dots) Add(a Dot) Dots {
	if list.Has(a) {
		return list
	}
	return append(list, a)
}

// Insert additional dot in the mapping.
func (mapping OneToMany) Insert(one Dot, other Dot) Dot {
	mapping[one] = mapping.Get(one).Add(other)
	return one
}

type Connectable interface {
	AsDots() Dots
}

// Explains specifies that "from" is some complex entity/concept, with one of its details/examples exposed in "to".
func Arrows(from Connectable, to Connectable) Connectable {
	if src, ok := from.(Dot); ok {
		for _, dst := range to.AsDots() {
			Arrow(src, dst)
		}
	} else {
		for _, src := range from.AsDots() {
			Arrows(src, to)
		}
	}
	return from
}

func Arrow(from, to Dot) {
	more.Insert(from, to)
	less.Insert(to, from)
}

func (from Dot) Arrows(to Dots) Dot {
	for _, dot := range to {
		more.Insert(from, dot)
		less.Insert(dot, from)
	}
	return from
}

func (a Dot) AsDots() Dots {
	return Dots{a}
}

// More specific
func (a Dot) More() Dots {
	return more.Get(a)
}

// Less specific 
func (a Dot) Less() Dots {
	return less.Get(a)
}

/// Version next.

/*

 Mixing graphs with node contents is confusing.

 Focus on connecting dots first.

*/
/*
type Oo []int // As in Ooo... dots.

func (a Oo) Same(b ...int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, dot := range a {
		if dot != b[i] {
			return false
		}
	}
	return true
}

// Good clones, sorts and removes duplicates.
func (a Oo) Good() Oo {
	sort.Ints(a)
	result := Oo{}
	for i, dot := range a {
		if i > 0 && result[len(result)-1] == dot {
			continue
		}
		result = append(result, dot)
	}
	return result
}

// Operation can do union/subtraction/xor/a-b; a,b assumed Good. TODO
func operation( a, b Oo) Oo{
	result := Oo{}
	return result
}
*/
// Graph
type Dot struct{ down []int }
type Arrow [2]int
type Graph []Dot

type Grapher interface {
	Arrow(a, b int) // Connects dots.
	Dot() int       // Gets a dot.
	Arrows() []Arrow
	Size() int
}

func (g *Graph) Arrow(a, b int) {
	dot := &(*g)[a]
	dot.down = append(dot.down, b)
}

func (g *Graph) Dot() int {
	index := len(*g)
	*g = append([]Dot(*g), Dot{[]int{}})
	return index
}

func (g *Graph) Arrows() []Arrow {
	arrows := []Arrow{}
	for from, to_dots := range *g {
		for _, to := range to_dots.down {
			arrows = append( arrows, Arrow{from, to})
		}
	} 
	return arrows
}

func (g *Graph) Size() int { return len(*g) }

func NewGrapher() Grapher {
	return &Graph{}
}

// Noter accepts information from a scanner.
type Noter interface {
	Down()
	Up()
	Token([]byte)
}

// Data precessors implement Scanner 
type Scanner interface {
	Scan(Noter)
}

// ElispScanner implements scanner interface for elisp files.
type ElispScanner []byte

// Scan goes through elisp, counting parens and reporting tokens (TODO) 
func (bytes ElispScanner) Scan(n Noter) {
	is_comment := false
	is_string := false
	do_escape := false
	for _, byte := range bytes {
		// String
		switch {
		case is_comment: // Do nothing
		case do_escape:
			do_escape = false
			continue
		case byte == '\\': // Not the best way...
			do_escape = true
		case byte == '"':
			is_string = !is_string
		}

		// Comment
		switch {
		case is_string: // Do nothing
		case byte == '\n':
			is_comment = false
		case byte == ';':
			is_comment = true
		}

		// Counting
		switch {
		case is_string:
		case is_comment: // Do nothing
		case byte == '(':
			n.Down()
		case byte == ')':
			n.Up()
		}
	}
}
