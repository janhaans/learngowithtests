package iteration

func Repeat(char string, num int) string {
	var repeatedChar string
	for i := 0; i < num; i++ {
		repeatedChar += char
	}
	return repeatedChar
}
