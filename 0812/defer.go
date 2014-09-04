package main

import (
	//"errors"
	"fmt"
	"io"
	"os"
)

func CopyFile(dst string, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

func main() {
	_, err := CopyFile("./2.txt", "./1.txt")
	if err == nil {
		fmt.Println("copy sucess")
	} else {
		fmt.Println(err)
	}
}
