package iteration

const repeatCount = 5

// Repeat takes character and returns it repeated 5 times
func Repeat(character string, count int) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
