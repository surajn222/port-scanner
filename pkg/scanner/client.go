package scanner

import (
	"net"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func ScanPort(target string, timeoutDuration time.Duration, wg *sync.WaitGroup) {
	conn, err := net.DialTimeout("tcp", target, timeoutDuration)

	port := strings.Split(target, ":")[1]
	if err != nil {
		wg.Done()
		return
	}

	logrus.Infof("Port %s is open", port)
	conn.Close()
	wg.Done()
}
