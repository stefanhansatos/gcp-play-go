package main

import (
	"fmt"
	"net"
	"net/http"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	ipv4 := GetLocalIP()
	fmt.Printf("ipv4: %s\n", ipv4)

	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(fmt.Sprintf("%s:8080", ipv4), nil)
	if err != nil {
		fmt.Printf("failed to listen and serve: %v\n", err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}
