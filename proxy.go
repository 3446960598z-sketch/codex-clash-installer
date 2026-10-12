package main

import (
	"fmt"
	"net"
)

func detectProxies() []string {
	var proxies []string
	// Check env, system, common ports 7890, 7897, 1080 etc.
	ports := []string{"7890", "7897", "7891", "1080", "10809"}
	for _, p := range ports {
		conn, err := net.Dial("tcp", "127.0.0.1:"+p)
		if err == nil {
			conn.Close()
			proxies = append(proxies, "http://127.0.0.1:"+p)
			fmt.Printf("Found local proxy on port %s\n", p)
		}
	}
	return proxies
}
