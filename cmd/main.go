package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/surajn222/port-scanner/pkg/scanner"
)

func main() {
	var wg sync.WaitGroup

	targetDomain := flag.String("host", "www.google.com", "domain to be scanned")
	startPort := flag.Int("startPort", 0, "Starting Port")
	endPort := flag.Int("endPort", 100, "End port")
	flag.Parse()

	for i := *startPort; i < *endPort; i++ {
		wg.Add(1)
		target := fmt.Sprintf("%s:%s", *targetDomain, strconv.Itoa(i))
		timeoutDuration := time.Millisecond * 100
		go scanner.ScanPort(target, timeoutDuration, &wg)
	}

	wg.Wait()
}
