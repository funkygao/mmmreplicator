package main

import (
	"sync"
)

type worker struct {
	cf *workerConfig
}

func newWorker(cf *workerConfig) *worker {
	this := new(worker)
	this.cf = cf
	return this
}

func (this *worker) start(wg *sync.WaitGroup) {
	defer wg.Done()
}
