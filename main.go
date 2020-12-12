package main

import (
	"filegea/config"
	"filegea/server"
	"fmt"
	"os"
)

func main() {
	if err := config.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conf := config.GetConfig()
	server.Init().Run(conf.Port)
}
