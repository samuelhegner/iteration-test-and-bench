package iteration

func Repeat(c string, iterations int) string {

	var repeated string

	for i := 0; i < iterations; i++ {
		repeated += c
	}

	return repeated
}
