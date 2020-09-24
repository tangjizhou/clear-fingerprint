package main

import (
	"./file"
	"flag"
	"fmt"
	"os"
)

var (
	path    string
	confirm bool
)

func main() {

	defer func() {
		if e := recover(); e != nil {
			fmt.Println("process failed.", e)
			os.Exit(1)
		}
	}()

	flag.BoolVar(&confirm, "y", false, "confirm")
	flag.StringVar(&path, "p", "", "file path")
	flag.Parse()

	//validate()
	askConfirm()
	path = "/Users/tangjizhou/work/project/clear-fingerprint/test/ip001"
	fileChannel, scanCompleteChannel := file.Scan(path)
	defer close(*fileChannel)
	defer close(*scanCompleteChannel)

	file.Clear(fileChannel, scanCompleteChannel)

}

func validate() {
	if path == "" {
		panic("file path required")
	}
}

func askConfirm() {
	if confirm {
		return
	}
	fmt.Print("is the path confirmed [", path, "],yes or no: ")
	var answer string
	_, err := fmt.Scanf("%s", &answer)
	if err != nil {
		panic(err)
	}
	if answer != "yes" && answer != "y" {
		panic("operation canceled")
	}
}
