package main

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

type DNSResult struct {
	Ips    []string
	Domain string
	Err    error
}

func resolver(ctx context.Context, i int, jobs <-chan string, dnsresult chan<- DNSResult, ratelimiter <-chan time.Time, wg *sync.WaitGroup) {
	//func resolver(domain string, res chan<- DNSResult, wg *sync.WaitGroup) {

	defer wg.Done()

	for domain := range jobs {

		select {

		case <-ctx.Done():
			return

		case <-ratelimiter:
		}

		var ips []string
		var err error

		for attempt := 1; attempt <= 3; attempt++ {
			ips, err = net.DefaultResolver.LookupHost(ctx, domain)

			if err == nil {
				break
			}

			time.Sleep(time.Duration(attempt) * 200 * time.Millisecond)

		}

		dnsresult <- DNSResult{
			Ips:    ips,
			Domain: domain,
			Err:    err,
		}

	}

	//addr, err := net.LookupHost(domain)

	// if err != nil {
	// 	res <- DNSResult{
	// 		Domain: domain,
	// 		Err:    err,
	// 	}
	// 	return
	// }

	// res <- DNSResult{
	// 	Ip:     addr,
	// 	Domain: domain,
	// 	Err:    err,
	// }
}

func main() {

	domains := []string{
		"google.com",
		"github.com",
		"youtube.com",
		"openai.org",
		"openai.com",
		"golang.org",
		"stackoverflow.com",
	}

	var dnschannel = make(chan DNSResult)
	var jobchannel = make(chan string)
	var wg sync.WaitGroup

	ratelimiter := time.Tick(time.Second / 5)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	workerpool := 3

	for i := 1; i <= workerpool; i++ {
		wg.Add(1)
		go resolver(ctx, i, jobchannel, dnschannel, ratelimiter, &wg)
	}

	go func() {

		for _, domain := range domains {
			jobchannel <- domain
		}
		close(jobchannel)
	}()

	// for _, domain := range domains {
	// 	wg.Add(1)
	// 	go resolver(domain, dnschannel, &wg)
	// }

	go func() {
		wg.Wait()
		close(dnschannel)
	}()

	for dnsch := range dnschannel {

		if dnsch.Err != nil {
			fmt.Println("Error resolving domain:", dnsch.Domain, dnsch.Err)
			continue
		}

		fmt.Printf("%s -> %v \n", dnsch.Domain, dnsch.Ips)
	}

}
