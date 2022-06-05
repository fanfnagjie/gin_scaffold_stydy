package main

import (
	"github.com/e421083458/gin_scaffold_study/golang_common/lib"
	"github.com/e421083458/gin_scaffold_study/router"
	//"github.com/e421083458/gin_scaffold_study/golang_common/lib"

	"os"
	"os/signal"
	"syscall"
)

func main()  {
	lib.InitModule("./conf/dev/")
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}

//
//export GOROOT="/usr/local/Cellar/go/1.18.2/libexec"
//export GOPATH="/Users/apple/doc/GoProjects"
//export PATH="/Users/apple/doc/GoProjects/bin:$PATH"
//export GO111MODULE=on
//export GOPROXY=https://mirrors.aliyun.com/goproxy/