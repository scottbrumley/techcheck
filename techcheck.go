package main

import (
	"github.com/scottbrumley/epo"
	"fmt"
	"strings"
)

type TechCheckResults struct {
	totalSystems int
	systemTypes map[string]int
	systemAgents map[string]int
}

// SystemTree Settings
func CheckSystemTree(data epo.SystemProperties)(systemChecks TechCheckResults){
	lineNum := 1
	osTypes := make(map[string]int)
	agentVersion := make(map[string]int)
	OSStr := ""

	for i := range data {
		//fmt.Printf("Record # %v\n", lineNum)
		for key, value := range data[i] {
			//fmt.Println("Key:", key, "Value:", value)

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

			// Make a list of all Operating Systems and Count them
			if (key == "EPOLeafNode.AgentVersion"){
				OSStr = fmt.Sprint(value)
				agentVersion[OSStr] = agentVersion[OSStr] + 1
			}
		}
		lineNum = lineNum + 1
		//fmt.Println("")
	}
	systemChecks.totalSystems = lineNum - 1
	systemChecks.systemTypes = osTypes
	systemChecks.systemAgents = agentVersion

	return systemChecks
}

// EPO Settings
type EpoSettings struct {
	version string
}

func importDashboards(myParms epo.ParamStruct){
	// Importing Dashboards
	fmt.Println("")
	fmt.Println("Importing Tech Check Dashboards")
	// Edit Settings for Import
	myParms.Cmd = "clienttask.importClientTask"
	myParms.Parms = "importFileName=tech_check_operations_health_dashboard.xml"
	myParms.Query = myParms.Url + "/" + myParms.Cmd + "?" + myParms.Parms + "&:output=" + myParms.Output
	fmt.Println(myParms.Query)

	jsonStr := epo.GetUrl(myParms)
	fmt.Println(jsonStr)
}

func GetEpoVersion(myParms epo.ParamStruct){
	// ePO Version
	// Importing Dashboards
	fmt.Println("")
	fmt.Println("Checking ePO ...")
	// Edit Settings for Import
	myParms.Cmd = "epo.getVersion"
	myParms.Query = myParms.Url + "/" + myParms.Cmd + "?:output=" + myParms.Output
	jsonStr := epo.GetUrl(myParms)
	fmt.Println(jsonStr)
}

func CheckSystems(myParms epo.ParamStruct){
	// Checking System Tree
	fmt.Println("")
	fmt.Println("Checking System Tree ...")

	// Edit Settings for System Tree
	myParms.Cmd = "system.find"
	myParms.Parms = "searchText=."
	myParms.Query = myParms.Url + "/" + myParms.Cmd + "?" + myParms.Parms + "&:output=" + myParms.Output

	jsonStr := epo.GetUrl(myParms)
	data := epo.DecodeJson(jsonStr)
	treeResults := CheckSystemTree(data)
	fmt.Println("Results")
	fmt.Println(treeResults)
}

func main() {
	var myParms epo.ParamStruct
	myParms = epo.GetParams()
	myParms.SslIgnore = true

	CheckSystems(myParms)
	GetEpoVersion(myParms)
}