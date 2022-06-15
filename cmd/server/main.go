package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/linyuan-t/quick-search/backend/server"
)

var _version_ = ""
var _branch_ = ""

func main() {

	//1.从输入参数读取文件名，比如 ./backend --filename=../../config/backend-dev.yaml
	filename := kingpin.Flag("filename", "Config file name.").String()
	version := kingpin.Flag("version", "Display Version.").Bool()
	kingpin.Parse()

	fmt.Printf("Version: %s\nBranch: %s\n", _version_, _branch_)
	if *version {
		os.Exit(-1)
	}

	server.Run(*filename)
}
