package visitor

// Square
type Square struct {
	Side int
}

func (s *Square) Accept(v visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

// Circle
type Circle struct {
	Radius int
}

func (c *Circle) Accept(v visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

// Rectangle
type Rectangle struct {
	L int
	B int
}

func (t *Rectangle) Accept(v visitor) {
	v.visitForrectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}
