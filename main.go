package main

import "filegea/server"

func main() {
	port := "1270"
	server.Init().Run(port)
}
