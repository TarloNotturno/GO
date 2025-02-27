package main

// function cleaning the input and converting upper case to lower case
func clearUnexpectedChar(inputString string) string {
	var result string
	for n, _ := range inputString {
		if _, ok := fromCharToInt[string(lowUpConversion[string(inputString[n])])]; ok {
			result += string(lowUpConversion[string(inputString[n])])
		}
	}
	return result
}

// main conversion function
func FromRomanToArab(inputString string) int64 {
	n := 0
	inputString = clearUnexpectedChar(inputString)
	var convertedNumb int64
	oldVal := int64(0)
	for _, val := range inputString {
		app := int64(fromCharToInt[string(val)])
		if oldVal >= app {
			convertedNumb += oldVal
		} else {
			convertedNumb += -oldVal
		}
		oldVal = app
		n++
	}
	convertedNumb += oldVal
	return convertedNumb
}
