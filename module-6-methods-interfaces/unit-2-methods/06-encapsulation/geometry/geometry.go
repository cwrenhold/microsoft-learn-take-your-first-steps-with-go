package geometry

// Note: Upper case struct names are exported.
type Triangle struct {
	size int
}

// Note: Lower case method names are not exported.
func (t *Triangle) doubleSize() {
	t.size *= 2
}

// Note: Upper case method names are exported.
func (t *Triangle) SetSize(size int) {
	t.size = size
}

func (t *Triangle) Perimeter() int {
	t.doubleSize()
	return t.size * 3
}
