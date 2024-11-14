package convert

// ApplyPointerToSlice применяет функцию fn к каждому элементу inputSlice, возвращает слайс
func ApplyPointerToSlice[in, out any](inputSlice []*in, fn func(*in) *out) []out {
	if inputSlice == nil {
		return nil
	}

	outputSlice := make([]out, 0, len(inputSlice))

	for _, item := range inputSlice {
		outputSlice = append(outputSlice, *fn(item))
	}

	return outputSlice
}

// ApplyToSlice применяет функцию fn к каждому элементу inputSlice, возвращает слайс
func ApplyToSlice[in, out any](inputSlice []in, fn func(in) out) []out {
	if inputSlice == nil {
		return nil
	}

	outputSlice := make([]out, 0, len(inputSlice))

	for _, item := range inputSlice {
		outputSlice = append(outputSlice, fn(item))
	}

	return outputSlice
}
