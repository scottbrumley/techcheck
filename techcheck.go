package main

import (
	"github.com/scottbrumley/epo"
	"fmt"
	"strings"
)

type TechCheckResults struct {
	totalSystems int
	systemTypes map[string]int
}

// Check SystemTree
func CheckSystemTree(data epo.SystemProperties)(systemChecks TechCheckResults){
	lineNum := 1
	osTypes := make(map[string]int)
	OSStr := ""

	for i := range data {
		fmt.Printf("Record # %v\n", lineNum)
		for key, value := range data[i] {
			fmt.Println("Key:", key, "Value:", value)

			// Collect McAfee Web Gateway
			if (key == "EPOComputerProperties.OSVersion"){
				OSStr = fmt.Sprint(value)
				if ( strings.Contains(OSStr,".mwg.") ){
					osTypes["MWG"] = osTypes["MWG"] + 1
				}
			}

			// Collect McAfee Components
			if (key == "EPOComputerProperties.OSOEMID"){
				OSStr = fmt.Sprint(value)
				if ( strings.Contains(OSStr,"McAfee") ){
					osTypes[OSStr] = osTypes[OSStr] + 1
				}

			}

			// Make a list of all Operating Systems and Count them
			if (key == "EPOComputerProperties.OSType"){
				OSStr = fmt.Sprint(value)
				osTypes[OSStr] = osTypes[OSStr] + 1
			}
		}
		lineNum = lineNum + 1
		fmt.Println("")
	}
	systemChecks.totalSystems = lineNum - 1
	systemChecks.systemTypes = osTypes

	return systemChecks
}

func main() {

	myParms := epo.GetParams()
	jsonStr := epo.GetUrl(myParms)
	data := epo.DecodeJson(jsonStr)

	checkResults := CheckSystemTree(data)
	fmt.Println(checkResults)

}