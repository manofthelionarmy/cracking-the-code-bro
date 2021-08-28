package main

func getShift(r rune) uint64 { // Make this more robust (a0 fails because both point to 0)
	var char rune
	if r >= 'a' && r <= 'z' {
		char = 'a'
	} else if r >= 'A' && r <= 'Z' {
		char = 'A'
	} else if r >= '0' && r <= '9' {
		char = '0'
	}

	return uint64(r - char)
}

func BitVectorHasUniqueCharacters(input string) bool {
	bitVector := uint64(0)
	// TODO: make sure to deep process this
	for _, r := range input {
		flippedBit := uint64(1)
		flippedBit = flippedBit << getShift(r) // TODO: need a way to get the bit at the bit position in the bit vector
		bitMask := bitVector ^ flippedBit      // flip the bit for inverted
		// flip the bit for inverted
		bitMask = (bitVector | flippedBit) ^ (bitMask) // Read as: append the bit to the bit vector, then xor against the reinverted bits
		if bitMask != flippedBit {
			bitVector = bitVector | flippedBit
		} else {
			return false
		}
	}

	return true
}
