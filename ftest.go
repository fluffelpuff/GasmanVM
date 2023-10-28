package main

import (
	"fmt"

	"github.com/fluffelpuff/GasmanVM/fsysfile"
)

func main() {
	ffile, err := fsysfile.OpenVirtualFileSystemContainer("C:\\Users\\FluffelBuff\\Desktop\\test\\filesys")
	if err != nil {
		panic(err)
	}
	tresult, err := ffile.GetFileByFullPath("/test/Test_2/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(tresult)
}
