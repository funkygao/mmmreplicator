package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

const (
	VERSION = "v0.0.1.alpha"
	AUTHOR  = "funky.gao@gmail.com"
)

var (
	confiFile   string
	showVersion bool

	BuildID = "unknown" // git version id, passed in from shell
)

func parseFlags() {
	flag.StringVar(&confiFile, "conf", "etc/mmmd.cf", "config file")
	flag.BoolVar(&showVersion, "version", false, "show version and exit")

	flag.Parse()
}

func ShowVersionAndExit() {
	fmt.Fprintf(os.Stderr, "%s %s (build: %s)\n", os.Args[0], VERSION,
		BuildID)
	fmt.Fprintf(os.Stderr, "Built with %s %s for %s/%s\n",
		runtime.Compiler, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	os.Exit(0)
}
