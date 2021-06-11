package robotname

// import (
// 	"fmt"
// )

// // pre-generate the sequence
// const maxRobots int32 = 676000 // 26*26*10*10*10

// // Robot is a friendly automaton
// type Robot struct {
// 	name string
// }

// var zero = struct{}{}

// // randomly ordered sequence array
// var names = makeSequence()

// func makeSequence() [maxRobots]int32 {
// 	m := mapMaker()
// 	var a [maxRobots]int32
// 	i := 0
// 	//range over map gives random index k back
// 	for k := range m {
// 		a[i] = k
// 		i++
// 	}
// 	return a
// }

// ////makes a map of empty structs
// func mapMaker() map[int32]struct{} {
// 	m := make(map[int32]struct{}, maxRobots+1)
// 	for i := int32(0); i < maxRobots; i++ {
// 		m[i] = zero
// 	}
// 	return m
// }

// // keeps track of where we are in the sequence
// var count int32

// // Name returns a robots name
// func (r *Robot) Name() (string, error) {
// 	var err error
// 	if r.name == "" {
// 		err = r.Reset()
// 	}
// 	return r.name, err
// }

// // Reset generates a new name for this Robot
// func (r *Robot) Reset() error {
// 	p := names[count%maxRobots]
// 	r.name = generateName(p)
// 	count++
// 	if count > maxRobots {
// 		return fmt.Errorf("maxRobots exceeded %d (got %d)", maxRobots, count)
// 	}
// 	return nil
// }

// // given a number p generate a name in the pattern AA###
// func generateName(i int32) string {
// 	nameRunes := make([]rune, 5)
// 	for pos, r := range "AA000" {
// 		switch r {
// 		case 'A':
// 			nameRunes[pos], i = getRuneAndRemainder(i, r, 26)
// 		default:
// 			nameRunes[pos], i = getRuneAndRemainder(i, r, 10)
// 		}
// 	}
// 	return string(nameRunes)
// }

// func getRuneAndRemainder(remainder int32, start int32, max int32) (rune, int32) {
// 	r := rune(start + remainder%max)
// 	return r, remainder / max
// }

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// //Robot is just a name (for now)
// type Robot string

// //namesInUse keeps track of which names are in use
// var namesInUse = make(map[Robot]bool)

// // Name returns a robots name
// func (r *Robot) Name() (string, error) {
// 	rand.Seed(time.Now().UnixNano())
// 	for inUse, ok := string(*r) == "", true; ok && inUse; inUse, ok = namesInUse[*r] {
// 		*r = Robot(letter() + letter() + number())
// 	}
// 	namesInUse[*r] = true
// 	return string(*r), nil
// }

// // Reset give a robot a new name
// func (r *Robot) Reset() {
// 	*r = Robot("")
// }

// // letter generates a random upper case letter
// func letter() string {
// 	return fmt.Sprintf("%c", rune('A'+rand.Intn('Z'-'A')))
// }

// //number generates a random three digit number
// func number() string {
// 	return fmt.Sprintf("%03d", rand.Intn(1000))
// }

// type Robot struct {
//     name string
// }

// const NAMESPACE_SIZE = ('Z'-'A'+1) * ('Z'-'A'+1) * 1000

// var issuedNames = make(map[string]bool)

// func (r *Robot) Name() (string, error) {
//     if r.name == "" {
//         r.Reset()
//         if r.name == "" {
//             return "", fmt.Errorf("Namespace exhausted")
//         }
//     }
//     return r.name, nil
// }

// func (r *Robot) Reset() {
//     if len(issuedNames) == NAMESPACE_SIZE {
//         r.name = ""
//         return
//     }

//     for {
//         l1 := rune('A' + rand.Intn('Z'-'A'+1))
//         l2 := rune('A' + rand.Intn('Z'-'A'+1))
//         num := rand.Intn(1000)
//         name := fmt.Sprintf("%c%c%03d", l1, l2, num)

//         _, nameIssued := issuedNames[name]
//         if !nameIssued {
//             r.name = name
//             issuedNames[name] = true
//             break
//         }
//     }
// }
