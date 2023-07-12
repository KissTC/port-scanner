package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

// cli flag --site=
var site = flag.String("site", "scanme.nmap.org", "url to scann")

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	for port := 1; port < 65535; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open\n", port)
		}(port)
	}
	wg.Wait()
}
