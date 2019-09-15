package util

import "flag"

type Args struct {
	ListenAddr string
}

func ProcessCmdLineArgs() Args {
	var args Args
	flag.StringVar(&(args.ListenAddr), "listen_addr", ":8080", "The address on which to listen for connections.")
	flag.Parse()
	return args
}