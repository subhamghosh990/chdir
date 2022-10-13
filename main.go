package main

import (
	"fmt"
	cd "main/code"
	"os"
)

func main() {
	var currPath, newPath string
	if len(os.Args) == 3 {
		currPath = os.Args[1]
		newPath = os.Args[2]
	}
	op := cd.OutPutPath(currPath, newPath)
	fmt.Println("Op : ", op)
}
