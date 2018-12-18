package main

import (
	server "myservices/books/svc/server"

	_ "github.com/mkevac/debugcharts"
	global "myservices/books/global"
)

func main() {

	global.SetPid(global.ProjectRealPath + "/runtime/pid")

	server.Run()
}
