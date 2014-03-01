package main

import (
	"sync"
)

func main() {
	parseFlags()
	if showVersion {
		ShowVersionAndExit()
	}

	cf := loadConfig(confiFile)
	wg := new(sync.WaitGroup)
	for _, w := range cf.workers {
		wg.Add(1)
		worker := newWorker(w)
		worker.start(wg)
	}

	// should never happen
	wg.Wait()
}
