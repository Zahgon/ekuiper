package main

import (
	"os"
)

func getClient(host, key string) { _ = "STUB: not implemented"; return }

// no password set
// use default DB

func setClient(host, key string) { _ = "STUB: not implemented"; return }

// no password set
// use default DB

func main() {
	if len(os.Args) == 4 {
		if v := os.Args[1]; v == "set" {
			// The 2nd parameter is MQTT broker server address
			setClient(os.Args[2], os.Args[3])
		}
		if v := os.Args[1]; v == "get" {
			// The 2nd parameter is MQTT broker server address
			getClient(os.Args[2], os.Args[3])
		}
	}
}
