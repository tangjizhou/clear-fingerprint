package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var encryptCount = 0

func Clear(fileChannel *chan PathedFile, scanCompleteChannel *chan bool) {
	for true {
		select {
		case file := <-*fileChannel:
			doClear(file)
		case <-*scanCompleteChannel:
			var length = len(*fileChannel)
			for i := 0; i < length; i++ {
				doClear(<-*fileChannel)
			}
			goto end
		}
	}
end:
	fmt.Println("encryptAesCBC complete,total: ", encryptCount)
}

func doClear(file PathedFile) {
	openedFile, err := os.OpenFile(file.path, os.O_RDWR, file.info.Mode().Perm())
	defer func() { _ = openedFile.Close() }()
	if err != nil {
		panic(err)
	}

	_, err = openedFile.Seek(0, io.SeekEnd)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(openedFile)
	line, err := reader.ReadBytes('\n')
	fmt.Println(line)
}
