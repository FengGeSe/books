package main

import (
	server "books/svc/server"

	global "books/global"
	_ "github.com/mkevac/debugcharts"
)

func main() {

	global.SetPid(global.ProjectRealPath + "/runtime/pid")

	server.Run()
}
