package main

import (
	"flag"
	"sync"
)

func main() {
	flag.Parse()
	cf := loadConfig(CONFIG_FILE)
	wg := new(sync.WaitGroup)
	for _, w := range cf.workers {
		wg.Add(1)
		worker := newWorker(w)
		worker.start(wg)
	}

	wg.Wait()
}
