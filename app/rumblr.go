package main

import (
	"github.com/gin-gonic/gin"
	"github.com/bksides/rumblr-go/app/util"
	"github.com/bksides/rumblr-go/app/routers"
)

func main() {
	args := util.ProcessCmdLineArgs()
	r := gin.Default()
	routers.InitRouter(r)
	r.Run(args.ListenAddr)
}