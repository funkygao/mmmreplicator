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
	host string
	port int
}

func (this *workerConfig) loadConfig(cf *conf.Conf) {
	this.host = cf.String("host", "")
	if this.host == "" {
		panic("empty host")
	}
	this.port = cf.Int("port", 27017)
}

func (this *workerConfig) dsn() string {
	return fmt.Sprintf("mongodb://%s:%d/", this.host, this.port)
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
