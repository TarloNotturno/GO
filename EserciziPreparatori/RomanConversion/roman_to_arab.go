package main

// count number of repetitions of a character
// and return the value of the character multiplied by the number of repetitions
func repeatedCharacter(inputString string, n int) (int64, int) {
	currentCharacter := lowUpConversion[string(inputString[n])]
	checkChar := currentCharacter
	checkIndex := n
	// find how many times current character is repeated
	numberOfRecurrrences := int64(0)
	for currentCharacter == checkChar && len(inputString) > checkIndex {
		checkChar = lowUpConversion[string(inputString[checkIndex])]
		checkIndex++
	}
	if len(inputString) != checkIndex || currentCharacter != checkChar {
		checkIndex--
	}
	numberOfRecurrrences = int64(checkIndex - n)
	/* if the char is repeated more than once, return the value of the char multiplied by the number of
	repetitions and its corresponding number in arab numbers */
	if numberOfRecurrrences > 0 {
		return numberOfRecurrrences * int64(fromCharToInt[currentCharacter]), checkIndex
	} else {
		return 0, n
	}

}

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
			currentVal, n = repeatedCharacter(inputString, n)
			convertedNumb += currentVal
		} else {
			n++
			nextVal, n = repeatedCharacter(inputString, n)
			convertedNumb += (nextVal - int64(currentVal))
		}
	}
	if n == len(inputString)-1 {
		convertedNumb += int64(fromCharToInt[lowUpConversion[string(inputString[n])]])
	}
	return convertedNumb
}
