package main

import (
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		fmt.Println("Error occured while reading directory: ", e)
		return
	}
}

func main() {

	files, err := os.ReadDir("/test/")
	checkError(err)

	fmt.Println("files in test:")
	for _, item := range files {
		fmt.Println(item)

		filePath := "/test/" + item.Name()

		file, err := os.ReadFile(filePath)
		checkError(err)

		dump := []byte(file)
		os.WriteFile("/test2/", dump, 0644)

		newFilePath := "/test2/" + item.Name()
		f, err := os.Create(newFilePath)
		checkError(err)

		defer f.Close()
	}

	files, err = os.ReadDir("/test2/")
	checkError(err)

	fmt.Println("files in test2:")
	if files != nil {
		for _, item := range files {
			fmt.Println(item)
		}
	} else {
		fmt.Println("no new files in test2")
	}

}
