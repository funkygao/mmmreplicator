package main

import (
	"github.com/funkygao/mmmreplicator/replicator"
	"labix.org/v2/mgo"
	"log"
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

func (this *worker) String() string {
	return "worker[" + this.cf.dsn() + "]"
}

func (this *worker) start(wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("%s connecting...", this)
	session, err := mgo.Dial(this.cf.dsn())
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()

	log.Printf("%s started", this)

	opChan, errChan := replicator.Tail(session, &replicator.Options{nil, nil})
	for {
		select {
		case err = <-errChan:
			log.Println(err)

		case op := <-opChan:
			log.Printf("%+v", op)
		}

	}
}
