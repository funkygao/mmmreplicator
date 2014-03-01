package main

import (
	"log"
	"sync"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	parseFlags()
	if showVersion {
		ShowVersionAndExit()
	}

	cf := loadConfig(confiFile)
	log.Printf("configuration loaded from %+v", confiFile)

	wg := new(sync.WaitGroup)
	for _, w := range cf.workers {
		wg.Add(1)
		worker := newWorker(w)
		worker.start(wg)
	}

	// should never happen
	wg.Wait()
}
