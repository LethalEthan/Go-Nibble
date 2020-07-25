package nibble

//CreateNibble - Creates a nibble by using half a byte that can be merged with another nibble later
func CreateNibble(N1 byte) byte {
	if N1 > 15 {
		N1 = 15
	}
	return N1
}

//CreateNibbleMerged - Creates a merged nibble by using 2 bytes (0-15) and merging into byte for space savings in arrays, can be separated with ReadNibbles()
func CreateNibbleMerged(N1 byte, N2 byte) byte {
	//prevent values over 15 and correct them to the highest value
	if N1 > 15 {
		N1 = 15
	}
	if N2 > 15 {
		N2 = 15
	}
	N1 = N1 << 4
	N3 := N1 | N2 //mix the numbers together
	return N3
}

//ReadNibbles - Reads Nibbles in a merged nibble byte and returns the values
func ReadNibbles(B byte) (byte, byte) {
	num1 := B >> 4
	num2 := B << 4
	num2 = num2 >> 4
	return num1, num2
}

//CreateNibbleArray - Creates a byte array that stores merged nibbles that can be read by ReadNibbles()
func CreateNibbleArray(length int) []byte {
	Array := make([]byte, length)
	return Array
}

//AppendUnmergedNibblestoArray - Append 2 unmerged nibbles into a byte array for space savings
func AppendUnmergedNibblestoArray(Array []byte, N1 byte, N2 byte) []byte {
	N1 = MergeNibbles(N1, N2)
	Array = append(Array, N1)
	return Array
}

//AppendMergedNibbletoNibbleArray - Append a merged nibble to nibble array
func AppendMergedNibbletoNibbleArray(Array []byte, N1 byte) []byte {
	Array = append(Array, N1)
	return Array
}

//ReadNibble1 - Reads the first nibble in a merged nibble byte
func ReadNibble1(B byte) byte {
	B = B >> 4
	if B > 15 {
		B = 15
	}
	return B
}

//ReadNibble2 - Reads the second nibble in a merged nibble byte
func ReadNibble2(B byte) byte {
	B = B << 4
	B = B >> 4
	if B > 15 {
		B = 15
	}
	return B
}

//MergeNibbles - Merges 2 non-merged nibbles togther into a byte
func MergeNibbles(N1 byte, N2 byte) byte {
	N1 = N1 << 4
	N3 := N1 | N2
	return N3
}
