package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sqweek/dialog"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	type File struct {
		Filename 	string `json:"Filename"`
		Bytes 		[]byte `json:"Bytes"`
	}

	response, err := http.Get("http://localhost:8080/files")
	check(err)

	responseBody, err := io.ReadAll(response.Body)
	check(err)

	var files  []File
	err = json.Unmarshal(responseBody, &files)
	check(err)

	pathToInstall, err := dialog.Directory().Title("Choose a civ6 folder: ").Browse()
	if err != nil {
		fmt.Println("an error occured while opening civ6 directory: ", err)
		return
	}
	
	pathToInstall = pathToInstall + "/Base/Binaries/Win64Steam/"

	for _, file := range files {
		os.WriteFile(pathToInstall + file.Filename, file.Bytes, 0644)
	}
	
	fmt.Println("success!")
}
