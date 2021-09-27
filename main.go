package main

import (
	"runtime"

	"github.com/onedss/filebrowser/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
