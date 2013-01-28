package dots

/* Plan: 
1. Efficient s-exp-like storage and processing
2. Entity names
3. Serialization

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
