package reverse

func String(input string) string {
	aInput := []rune(input)
	length := len(aInput)
	for i := 0; i < length/2; i++ {
		tmp := aInput[i]
		aInput[i] = aInput[length-i-1]
		aInput[length-i-1] = tmp
	}
	return string(aInput)

}
