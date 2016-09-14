package httpserver

import (
	"fmt"
	"sync"

	"github.com/HailoOSS/go-service-layer/healthcheck"
)

const HealthCheckId = "com.HailoOSS.service.config.httpconnect"

var connectErr error = fmt.Errorf("Not yet connected")
var connectErrorLock sync.RWMutex

// HealthCheck asserts that the port has been bound
func HttpConnectHealthCheck() healthcheck.Checker {
	return checkHttpConnect
}

func checkHttpConnect() (map[string]string, error) {
	connectErrorLock.RLock()
	err := connectErr
	connectErrorLock.RUnlock()

	return nil, err
}

func SetConnectHealthCheck(err error) {
	connectErrorLock.Lock()
	connectErr = err
	connectErrorLock.Unlock()
}
