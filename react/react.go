package react

// A MyReactor implements the Reactor interface.
type MyReactor struct {
}

// A MyCell implements the Cell interface.
type MyCell struct {
	v            int
	dependencies []Cell
}

// Value returns the value of a Cell.
func (c MyCell) Value() int {
	return c.v
}

// A MyInputCell implements the InputCell interface.
type MyInputCell struct {
	*MyCell
}

// CreateInput creates an InputCell.
func (r MyReactor) CreateInput(v int) InputCell {
	return MyInputCell{MyCell: &MyCell{v: v}}
}

// CreateCompute1 creates a ComputeCell having one input.
func (r MyReactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	m := MyCompute1Cell{
		MyCell:    &MyCell{v: f(c.Value())},
		input:     &c,
		function:  f,
		callbacks: map[int]func(int){},
	}

	switch c.(type) {
	case MyInputCell:
		myC := c.(MyInputCell)
		myC.dependencies = append(myC.dependencies, &m)
	case MyCompute1Cell:
		myC := c.(MyCompute1Cell)
		myC.dependencies = append(myC.dependencies, &m)
	case MyCompute2Cell:
		myC := c.(MyCompute2Cell)
		myC.dependencies = append(myC.dependencies, &m)
	}

	return m
}

// CreateCompute2 creates a ComputeCell having two inputs.
func (r MyReactor) CreateCompute2(c, d Cell, f func(int, int) int) ComputeCell {
	m := MyCompute2Cell{
		MyCell:    &MyCell{v: f(c.Value(), d.Value())},
		input1:    &c,
		input2:    &d,
		function:  f,
		callbacks: map[int]func(int){},
	}

	switch c.(type) {
	case MyInputCell:
		myC := c.(MyInputCell)
		myC.dependencies = append(myC.dependencies, &m)
	case MyCompute1Cell:
		myC := c.(MyCompute1Cell)
		myC.dependencies = append(myC.dependencies, &m)
	case MyCompute2Cell:
		myC := c.(MyCompute2Cell)
		myC.dependencies = append(myC.dependencies, &m)
	}

	return m
}

// SetValue sets the value of an InputCell.
func (c MyInputCell) SetValue(v int) {
	c.v = v

	for _, d := range c.dependencies {
		d.Value()
	}
}

// A MyCompute1Cell is a ComputeCell with a single input.
type MyCompute1Cell struct {
	*MyCell
	input       *Cell
	function    func(int) int
	callbackKey int
	callbacks   map[int]func(int)
}

// Value returns the value of a Cell.
func (c MyCompute1Cell) Value() int {
	v := c.function((*c.input).Value())

	if c.v != v {
		c.v = v

		for _, f := range c.callbacks {
			f(v)
		}

		for _, d := range c.dependencies {
			d.Value()
		}
	}

	return v
}

// AddCallback adds a callback to a ComputeCell.
func (c MyCompute1Cell) AddCallback(f func(int)) Canceler {
	return addCallback(c.callbacks, f)
}

// A MyCompute2Cell is a ComputeCell with two inputs.
type MyCompute2Cell struct {
	*MyCell
	input1      *Cell
	input2      *Cell
	function    func(int, int) int
	callbackKey int
	callbacks   map[int]func(int)
}

// Value returns the value of a Cell.
func (c MyCompute2Cell) Value() int {
	v := c.function((*c.input1).Value(), (*c.input2).Value())

	if c.v != v {
		for _, f := range c.callbacks {
			f(v)
		}
		c.v = v
		for _, d := range c.dependencies {
			d.Value()
		}
	}

	return v
}

// AddCallback adds a callback to a ComputeCell.
func (c MyCompute2Cell) AddCallback(f func(int)) Canceler {
	return addCallback(c.callbacks, f)
}

// A MyCanceler removes a callback from a Cell.
type MyCanceler struct {
	function func()
}

// Cancel removes a callback from its Cell.
func (c MyCanceler) Cancel() {
	c.function()
}

// New creates a MyReactor.
func New() Reactor {
	return MyReactor{}
}

func addCallback(m map[int]func(int), f func(int)) Canceler {
	key, ok := 0, true
	for ok {
		key++
		_, ok = m[key]
	}

	m[key] = f

	return MyCanceler{func() { delete(m, key) }}
}
