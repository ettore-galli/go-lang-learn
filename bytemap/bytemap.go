package bytemap

type FibCalcFunction func(int, FibCalcFunction) int

func ToImageBuffer(imageMap []byte, width int16) []byte {

	buffer := make([]byte, len(imageMap))
	wBytes := width / 8

	for index, imageByte := range imageMap {
		for i := int16(0); i < 8; i++ {
			mask := byte(1 << (7 - i))
			imageBits := ((mask & imageByte) >> (7 - i)) << (7 - index%8)
			bytesBufferPos := int16(index)/wBytes + i*wBytes

			// fmt.Printf(
			// 	"index=%v, i=%v, bytesBufferPos=%v buffer[bytesBufferPos]=%v, mask=%v, imageByte=%v, imageBits=%v \n",
			// 	index, i, bytesBufferPos, buffer[bytesBufferPos], mask, imageByte, imageBits)

			buffer[bytesBufferPos] |= imageBits

		}
	}

	return buffer

}
