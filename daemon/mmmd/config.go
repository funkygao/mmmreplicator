package main

import (
	conf "github.com/funkygao/jsconf"
)

type config struct {
	*conf.Conf
	workers []*workerConfig
}

type workerConfig struct {
	dsn string
}

func loadConfig(fn string) *config {
	cf, err := conf.Load(fn)
	if err != nil {
		panic(err)
	}

	c := new(config)
	c.Conf = cf
	return c
}
