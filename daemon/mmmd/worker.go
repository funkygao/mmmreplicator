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

func (this *worker) start(wg *sync.WaitGroup) {
	defer wg.Done()

	session, err := mgo.Dial(this.cf.dsn)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()

	log.Printf("worker[%s] started", this.cf.dsn)

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
