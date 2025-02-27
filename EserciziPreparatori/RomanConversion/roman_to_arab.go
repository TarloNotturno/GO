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
	convertedNumb := int64(0)
	for n < len(inputString)-1 {
		currentVal := int64(fromCharToInt[string(inputString[n])])
		nextVal := int64(fromCharToInt[string(inputString[n+1])])
		if currentVal >= nextVal {
			convertedNumb += currentVal
		} else {
			convertedNumb += (nextVal - currentVal)
			n++
		}
		n++
	}
	if n == len(inputString)-1 {
		convertedNumb += int64(fromCharToInt[lowUpConversion[string(inputString[n])]])
	}
	return convertedNumb
}
