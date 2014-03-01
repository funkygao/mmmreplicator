package main

import (
	"flag"
)

func parseFlags() {
	flag.StringVar(&confiFile, "conf", "etc/mmmd.cf", "config file")
	flag.BoolVar(&showVersion, "version", false, "show version and exit")

	flag.Parse()
}
