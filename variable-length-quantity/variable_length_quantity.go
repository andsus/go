package variablelengthquantity

import (
	"errors"
)

// Starts from left of 32 bit integer and returns the first bit that has a 1
func getMaxBit(v uint32) int {
	for i := 31; i >= 0; i-- {

	}
}

func getNthBit(v, n uint32) int {
	nthBit := n               // reverse the endian
	movedV := v >> nthBit     // move the nth bit to the first position
	maskedValue := movedV & 1 // mask bit, select only the first bit
	//return int( (v >> (32-n)) & 1)
	return int(maskedValue)

}

// DecodeVarint decode vlq to uint32
func DecodeVarint(input []byte) (output []uint32, err error) {
	var (
		outInt uint32
		done   bool
	)

	for _, v := range input {
		outInt = (outInt << 7) + uint32(v&0x7f)
		if done = (v&0x80 == 0); done {
			output = append(output, outInt)
			outInt = 0
		}
	}
	if !done {
		return nil, errors.New("incomplete sequence")
	}
	return output, nil
}

// EncodeVarint encode uint32 to vlq
func EncodeVarint(input []uint32) (output []byte) {
	for _, v := range input {
		enc := []byte{byte(v & 0x7f)}
		for v >>= 7; v > 0; v >>= 7 {
			enc = append([]byte{byte(v&0x7f) | 0x80}, enc...)
		}
		output = append(output, enc...)
	}
	return output
}
