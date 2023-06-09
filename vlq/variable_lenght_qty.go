package variablelengthquantity

import (
	"errors"
	"math"
	"math/bits"
)

// EncodeVarint encodes an array of uint32 numbers into a byte slice using a variable-length quantity encoding scheme.
// The output byte slice is returned.
func EncodeVarint(input []uint32) (output []byte) {
	// Iterate over each number in the input array
	for _, number := range input {
		// Calculate the number of bytes required to encode the number
		numBytes := int(math.Ceil(float64(bits.Len32(number)) / 7))
		if numBytes == 0 {
			numBytes = 1
		}

		// Create a byte slice to hold the encoded number
		b := make([]byte, 5)

		// Encode the number into the byte slice
		for i := 0; i < len(b); i++ {
			byteValue := uint32(math.Pow(128, float64(len(b)-1-i)))
			b[i] = byte(number / byteValue)
			number -= (uint32(b[i]) * byteValue)
		}

		// Set the continuation bit for all but the last byte
		for i := len(b) - numBytes; i < len(b)-1; i++ {
			b[i] |= (1 << 7)
		}

		// Append the encoded number to the output byte slice
		output = append(output, b[len(b)-numBytes:]...)
	}

	// Return the output byte slice
	return output
}

// DecodeVarint decodes a byte slice into an array of uint32 numbers using a variable-length quantity encoding scheme.
// The input byte slice is expected to contain one or more encoded numbers.
// The output array of uint32 numbers is returned along with an error if the input byte slice is invalid.
func DecodeVarint(input []byte) ([]uint32, error) {
	// Initialize variables to hold the encoded values
	vList := [][]byte{}
	value := []byte{}

	// Iterate over each byte in the input byte slice
	for _, v := range input {
		// Append the byte to the current value slice
		value = append(value, v)

		// Check if the continuation bit is not set
		if v&(1<<7) == 0 {
			// Add the completed value slice to the list of encoded values
			vList = append(vList, value)

			// Reset the value slice for the next encoded value
			value = nil
		}
	}

	// Check if any encoded values were found
	if len(vList) == 0 {
		return nil, errors.New("incomplete sequence")
	}

	// Decode each encoded value into a uint32 number
	var output []uint32
	for _, v := range vList {
		var outputValue uint32
		for k, w := range v {
			if w&(1<<7) != 0 {
				w &^= (1 << 7)
			}
			outputValue += uint32(w) * uint32(math.Pow(128, float64(len(v)-1-k)))
		}
		output = append(output, outputValue)
	}

	// Return the decoded uint32 numbers and nil error
	return output, nil
}
