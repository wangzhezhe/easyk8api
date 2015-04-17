package main

import (
	//	"bytes"
	"easyk8api/Apitool"
	"fmt"
	//	"io/ioutil"
	//	"net/http"
	//	"os"
)

func main() {

	masterip := "10.10.103.86"
	masterport := "8080"
	//filename := "/testpatch.json"
	//	apitype_map := make(map[string]string)
	//	apitype_map["G"] = "GET"
	//	apitype_map["PO"] = "POST"
	//	apitype_map["PU"] = "PUST"
	//	apitype_map["D"] = "DELETE"

	//	apiversion_map := make(map[string]string)
	//	apitype_map["b1"] = "v1beta1"
	//	apitype_map["b2"] = "v1beta2"
	//	apitype_map["b3"] = "v1beta3"

	//getcommands := []string{"minions"}
	//postcommands := []string{"replicationControllers"}
	//deletecommands := []string{"replicationControllers", "redis-master-controller-test1"}
	//putcommands := []string{"replicationControllers", "frontend-controller"}
	patchcommands := []string{"replicationControllers", "frontend-controller"}

	//parameters := []string{"fields="}

	//status, _ := sentapi("GET", masterip, masterport, "v1beta1", commands, nil)
	//use "" to represent nil if you don't need filename
	//don't need to add "/" before filename
	status, message := sendapi.Sendapi("PATCH", masterip, masterport, "v1beta1", patchcommands, "")
	fmt.Println(status)
	fmt.Println(message)
}
