package iter

func RepeatChar(root string, repeat int) string {
	var repeated string
	for i := 0; i < repeat; i++ {
		repeated += root
	}
	return repeated
}
