package main

import (
	"os"
	"strings"
	"sync"

	"main.go/FileDealer"
)

func UpdateInterfacesList() {
	var ListOfInterfaces []FileDealer.SWCInterfaces
	var (
		wg        sync.WaitGroup
		scheduler sync.WaitGroup
	)
	//First of all find all c file related to the project
	listOfFiles := "C:/Progetti/Go/ListOfFiles.txt"
	FileDealer.FindAllASWFiles("C:/GIT/obcg6_sw_app/Components/ASW", listOfFiles)
	scheduler.Add(1)
	currentDir, _ := os.Getwd()
	app := strings.Split(currentDir, "\\")
	currentDir = app[0]
	for i := 1; i < len(app)-1; i++ {
		currentDir = currentDir + "/" + app[i]
	}
	destinationFile := currentDir + "/ListOfInterfaces.txt"
	go FileDealer.LaunchFileRead(listOfFiles, &ListOfInterfaces, &wg, &scheduler, destinationFile)
	scheduler.Wait()
	for _, interfaceCurr := range ListOfInterfaces {
		interfaceCurr.PrintInterface()
	}

}
