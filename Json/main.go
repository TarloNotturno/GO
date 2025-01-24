package main

import (
	"encoding/json"
	"fmt"
)

type StixPackageLookupResponse struct {
	Result            string   `json:"result"`
	MaliciousEntities []string `json:"malicious_entities"`
	ExecTimeUs        string   `json:"exec_time_us"`
}

type StixPackageLookupResponseUntagged struct {
	Result            string
	MaliciousEntities []string
	ExecTimeUs        string
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	app := StixPackageLookupResponse{Result: "app", MaliciousEntities: []string{"li", "ss"}, ExecTimeUs: "1ms"}
	fmt.Println(app)

	res1D := &StixPackageLookupResponse{
		Result:            "app",
		MaliciousEntities: []string{"li", "ss"},
		ExecTimeUs:        "1ms"}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &StixPackageLookupResponseUntagged{
		Result:            "ginoPasqualino",
		MaliciousEntities: []string{"li", "ss", "piùroba"},
		ExecTimeUs:        "10ms"}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"Result":"ginoPasqualino","malicious_entities":["li", "ss", "piùroba"],"exec_time_us":"6ms"}`)
	byt1 := byt
	var dat1 map[string]interface{}
	var dat StixPackageLookupResponse
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(byt1, &dat1); err != nil {
		panic(err)
	}
	//dat = dat.(StixPackageLookupResponse)
	fmt.Println("with type", dat) //.Result, dat.ExecTimeUs)
	fmt.Println(dat1["Result"])
	fmt.Println(dat1["Result"])

	//var res StixPackageLookupResponse
	//decoded := json.Unmarshal([]byte(StixPackageLookupResponse), &res)

}
