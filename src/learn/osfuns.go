package main

import "os"

func main() {
	hostname, err := os.Hostname()
	println(hostname, " ", err)

}
