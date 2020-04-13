package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", HelloServer)

	fmt.Printf("start to listen and serve at %s:%s\n", ipv4, port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", ipv4, port), nil)
	if err != nil {
		fmt.Printf("failed to listen and serve: %v\n", err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from inside a Cloud Run container!\n\n")
	fmt.Fprintf(w, "Here the current environment variables from inside the container:\n")

	for index, value := range os.Environ() {
		fmt.Fprintf(w, "%v:\t%s\n", index, value)
	}
}
