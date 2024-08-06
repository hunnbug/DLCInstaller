package main

import (
	"fmt"
	"os"

	"github.com/sqweek/dialog"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	pathToInstall, err := dialog.Directory().Title("Choose a civ6 folder: ").Browse()
	if err != nil {
		fmt.Println("an error occured while opening civ6 directory: ", err)
		return
	}

	pathToInstall = pathToInstall + "/Base/Binaries/Win64Steam/"

	currentDir, err := os.Getwd()
	check(err)

	filesToInsert, err := os.ReadDir(currentDir + "/files")
	check(err)

	for _, item := range filesToInsert {

		filePath := currentDir + "/files/" + item.Name()

		file, err := os.ReadFile(filePath)
		check(err)

		os.WriteFile(pathToInstall+item.Name(), file, 0644)
	}
	fmt.Println("success!")
}
