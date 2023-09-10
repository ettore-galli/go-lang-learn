package bytemap

type FibCalcFunction func(int, FibCalcFunction) int

func ToImageBuffer(imageMap []byte, width int16) []byte {

	buffer := make([]byte, len(imageMap))
	wBytes := width / 8

	for index, imageByte := range imageMap {

		for i := int16(0); i < 8; i++ {

			mask := byte(1 << i)

			imageBits := (mask & imageByte) << (7 - i)

			bytesBufferPos := int16(index)/wBytes + i*wBytes

			if wBytes > 0 {
				buffer[bytesBufferPos] |= imageBits
			}
		}

	}

	return buffer

}
