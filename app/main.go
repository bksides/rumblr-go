package main

import (
	"github.com/gin-gonic/gin"
	"flag"
)

type Args struct {
	ListenAddr string
}

func processCmdLineArgs() Args {
	var args Args
	flag.StringVar(&(args.ListenAddr), "listen_addr", ":8080", "The address on which to listen for connections.")
	flag.Parse()
	return args
}

func boopHandler(beep *gin.Context) {
	beep.Status(200)
}

func main() {
	args := processCmdLineArgs()
	r := gin.Default()
	r.GET("/boop", boopHandler)
	r.Run(args.ListenAddr)
}