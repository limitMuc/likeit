package main

import (
	"flag"
	"likeit/pkg"
	"log"
)

const (
	GO_FILE_EXT = ".go"
)

var (
	errHandleTpl = []byte(`err != nil`)
)

func main() {
	pth := flag.String("p", "", "files path")
	flag.Parse()

	path, err := pkg.ValidFilePath(*pth)
	if err != nil {
		log.Printf("'%s' isn't valid path. error: %s\n", *pth, err)
		return
	}

	node := pkg.ParsePath2TreeNode(path, GO_FILE_EXT, errHandleTpl)
	node.Print()
}
