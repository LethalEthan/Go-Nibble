package nibble

type DualNum struct {
	num1 byte
	num2 byte
}

func CreateNibble(B *DualNum) byte {
	num1 := B.num1 << 4
	num2 := B.num2 << 4
	num2 = num2 >> 4
	num3 := num1 | num2 //mix the numbers together
	return num3
}

func CreateDualNum(N1 byte, N2 byte) *DualNum {
	DNUM := new(DualNum)
	DNUM.num1 = N1
	DNUM.num2 = N2
	if N1 > 15 {
		print("Error: num1 over 15 correcting to 15")
		DNUM.num1 = 15
	}
	if N2 > 15 {
		print("Error: num2 over 15 correcting to 15")
		DNUM.num2 = 15
	}
	return DNUM
}

func ReadNibble(B byte) (byte, byte) {
	num1 := B >> 4
	num2 := B << 4
	num2 = num2 >> 4
	return num1, num2
}

func CreateNibbleArray(length int) []byte {
	Array := make([]byte, length)
	return Array
}

func ReadNibble1(B byte) byte {
	B = B >> 4
	return B
}

func ReadNibble2(B byte) byte {
	B = B << 4
	B = B >> 4
	return B
}
