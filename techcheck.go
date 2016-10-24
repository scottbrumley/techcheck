package main

import (
	"github.com/scottbrumley/epo"
	"fmt"
)

func main() {

	myParms := epo.GetParams()
	jsonStr := epo.GetUrl(myParms)
	data := epo.DecodeJson(jsonStr)

	lineNum := 1
	for i := range data {
		fmt.Printf("Record # %v\n", lineNum)
		for key, value := range data[i] {
			fmt.Println("Key:", key, "Value:", value)
		}
		lineNum = lineNum + 1
		fmt.Println("")
	}

}