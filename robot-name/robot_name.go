package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

//Robot is just a name (for now)
type Robot struct {
	name string
}

const maxNumberRobotNames = 26 * 26 * 1_000

//namesInUse keeps track of which names are in use
var namesInUse = make(map[string]bool)

var seedRandom = rand.New(rand.NewSource(time.Now().UnixNano()))

// Name returns a robots name
func (r *Robot) Name() (string, error) {
	if r.name != "" { // test case name sticks
		return r.name, nil
	}
	if len(namesInUse) >= maxNumberRobotNames {
		return "", fmt.Errorf("number of namesInUse exceeds %06d", maxNumberRobotNames)
	}
	r.name = newName()
	for namesInUse[r.name] {
		r.name = newName()
	}

	namesInUse[r.name] = true
	return r.name, nil
}

// Reset give a robot a new name
func (r *Robot) Reset() {
	r.name = ""
}

func newName() string {
	r1 := seedRandom.Intn(26) + 'A'
	r2 := seedRandom.Intn(26) + 'A'
	n := seedRandom.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, n)
}
