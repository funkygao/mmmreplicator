package main

import (
	"fmt"
	conf "github.com/funkygao/jsconf"
)

type config struct {
	*conf.Conf
	workers []*workerConfig
}

type workerConfig struct {
	dsn string
}

func (this *workerConfig) loadConfig(cf *conf.Conf) {
	host := cf.String("host", "")
	if host == "" {
		panic("empty host")
	}
	port := cf.Int("port", 27017)
	this.dsn = fmt.Sprintf("mongodb://%s:%d/", host, port)
}

func loadConfig(fn string) *config {
	cf, err := conf.Load(fn)
	if err != nil {
		panic(err)
	}

	c := new(config)
	c.Conf = cf

	workers, err := cf.Section("workers")
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(workers.List("servers", nil)); i++ {
		section, err := workers.Section(fmt.Sprintf("servers[%d]", i))
		if err != nil {
			panic(err)
		}

		worker := new(workerConfig)
		worker.loadConfig(section)
		c.workers = append(c.workers, worker)
	}

	return c
}
