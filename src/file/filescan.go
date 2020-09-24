package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var scanFileSuffixes = []string{".v", ".sv", ".svkey"}
var fileChannel = make(chan PathedFile, 100)
var scanCompleteChannel = make(chan bool, 1)

type PathedFile struct {
	info os.FileInfo
	path string
}

func Scan(path string) (*chan PathedFile, *chan bool) {
	go doScan(path)
	return &fileChannel, &scanCompleteChannel
}

func doScan(rootPath string) {
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		if err != nil {
			return filepath.SkipDir
		}
		if isFileExtMatched(info.Name()) {
			fileChannel <- PathedFile{info, path}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		scanCompleteChannel <- true
	}()
}

func isFileExtMatched(filename string) bool {
	for _, value := range scanFileSuffixes {
		if strings.HasSuffix(filename, value) {
			return true
		}
	}
	return false
}
