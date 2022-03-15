package main

import (
	"fmt"
	"os"
)

func main() {

	// 1. Create file
	newFile, err := os.Create("test001.txt")
	if err != nil {
		fmt.Println(err)
	}
	newFile.Close()

	// 2. File info
	fileInfo, err := os.Stat("test001.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileInfo)

	// 3. Remove
	err = os.Remove("test001.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 4. Chmod
	err = os.Chmod("test001.txt", 0777)
	if err != nil {
		fmt.Println(err)
	}

	// 5. File exsit
	fileInfo, err = os.Stat("test001.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(err)
		}
	}

}
