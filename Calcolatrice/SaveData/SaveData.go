package SaveData

import (
	"encoding/json"
	"fmt"
	"os"

	"main.go/geometry"
)

type FileData struct {
	StraightLines []geometry.StraightLine
	Points        []geometry.Point
	Plots         []geometry.CartesianPlot
}

func WriteData(filename string, input ...interface{}) {
	file, _ := os.OpenFile(filename, os.O_CREATE, os.ModePerm)
	for _, currentInput := range input {
		switch v := currentInput.(type) {
		case []geometry.StraightLine:
			encoder := json.NewEncoder(file)
			encoder.Encode("ArrayStraightLines")
			encoder.Encode(v)
		case geometry.StraightLine:
			encoder := json.NewEncoder(file)
			encoder.Encode("StraightLines")
			encoder.Encode(v)
		case FileData:
			encoder := json.NewEncoder(file)
			encoder.Encode("FileData:")
			encoder.Encode(v)
		default:
			fmt.Printf("The type is not a geometric asset %T!\n", v)
		}
	}
	file.Close()
}

func ReadData(filename string) (output FileData) {
	file, _ := os.Open(filename)
	defer file.Close()
	decoder := json.NewDecoder(file)
	///filteredData := []map[string]interface{}{}

	// Read the array open bracket
	decoder.Token()

	var data interface{}
	//data := map[string]interface{}{}
	 decoder.UnmarshalJSON(data)
	
		fmt.Println(n)
	
	//for decoder.More() {
	//	decoder.Decode(&data)
	//	fmt.Println(data)
	//	//fmt.Println(filteredData)
	//	//if filter(data) {
	//	//	filteredData = append(filteredData, data)
	//	//}
	//}

	return
}
