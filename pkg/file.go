package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	goos  = runtime.GOOS
	Total = 0
)

func readEachLine(path string, dst []byte) (int, error) {
	total := 0
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num := bytes.Count(scanner.Bytes(), dst)
		total += num
	}
	if scanner.Err() != nil {
		return 0, err
	}

	return total, nil
}

func getOSPathSeparator() string {
	switch goos {
	case "windows":
		return "\\"
	default:
		return "/"
	}
}

func GetAllFilesOfAbsPath(path string) []string {
	dir := []string{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("read '%s' failed. error: %s\n", path, err)
		return dir
	}

	for _, file := range files {
		parentPth := path + getOSPathSeparator()
		if file.IsDir() {
			subPath := GetAllFilesOfAbsPath(parentPth + file.Name())
			dir = append(dir, subPath...)
			continue
		}
		dir = append(dir, parentPth+file.Name())
	}
	return dir
}

func ParsePath2TreeNode(path, ext string, word []byte) *Node {
	node := NewNode(path, path, 1)
	return ParseFileNode(node, ext, word)
}

func ParseFileNode(node *Node, ext string, word []byte) *Node {
	files, err := ioutil.ReadDir(node.absPath)
	if err != nil {
		log.Printf("read '%s' failed. error: %s\n", node.absPath, err)
		return node
	}

	filesLen := len(files)
	for idx, file := range files {
		parentPth := node.absPath + getOSPathSeparator()
		isLast := idx == filesLen-1
		subNode := &Node{
			name:      file.Name(),
			isLast:    isLast,
			needSpace: node.needSpace,
			absPath:   parentPth + file.Name(),
			level:     node.level + 1,
			subNode:   []Node{},
		}

		if node.level > 1 {
			subNode.needSpace += INTERVAL_SPACE
		}

		if isLast {
			subNode.stdPrint = subNode.needSpace + LAST_NODE + file.Name()
		} else {
			subNode.stdPrint = subNode.needSpace + INTERNAL_NODE + file.Name()
		}

		if IsSpecifiedFile(file.Name(), ext) {
			num := WordStatistic(parentPth+file.Name(), word)
			subNode.number = num
			Total += num
		}

		if file.IsDir() {
			subNode.isDir = true
			subNode = ParseFileNode(subNode, ext, word)
		}
		node.subNode = append(node.subNode, *subNode)
	}

	return node
}

func WordStatistic(path string, word []byte) int {
	total := 0
	num, err := readEachLine(path, word)
	if err != nil {
		log.Printf("count word of '%s' failed. error: %s\n", path, err)
		return 0
	}
	total += num
	return total
}

func IsSpecifiedFile(path, ext string) bool {
	return filepath.Ext(path) == ext
}

func ValidFilePath(path string) (string, error) {
	if len(path) == 0 {
		return "", fmt.Errorf("path can't be empty")
	}

	path = filepath.Clean(path)
	if filepath.IsAbs(path) {
		return path, nil
	}

	return filepath.Abs(path)
}