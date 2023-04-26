/*
 * @Author: iRorikon
 * @Date: 2023-04-26 11:12:41
 * @FilePath: \api-service\service\download\clinet.go
 */
package download

import (
	"context"
	"net"
	"net/http"
	"runtime"
	"time"
)

func newDownloadClinet(maxIdleConnsPerHost int) *http.Client {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	dialer := newDialRateLimiter(&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	})
	tr.DialContext = dialer.DialContext
	tr.MaxIdleConns = 0 // no limit
	tr.MaxIdleConnsPerHost = maxIdleConnsPerHost
	return &http.Client{
		Transport: tr,
	}
}

func newClient(client *http.Client) *http.Client {
	if client == nil {
		return http.DefaultClient
	}
	return client
}

type dialRateLimiter struct {
	dialer *net.Dialer
	sem    chan struct{}
}

func newDialRateLimiter(dialer *net.Dialer) *dialRateLimiter {
	// exact value doesn't matter too much, but too low will be too slow,
	// and too high will reduce the beneficial effect on thread count
	const concurrentDialsPerCpu = 10

	return &dialRateLimiter{
		dialer: dialer,
		sem:    make(chan struct{}, concurrentDialsPerCpu*runtime.NumCPU()),
	}
}

func (d *dialRateLimiter) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	d.sem <- struct{}{}
	defer func() { <-d.sem }()
	return d.dialer.DialContext(ctx, network, addr)
}
