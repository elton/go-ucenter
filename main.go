package main

import (
	"log"
	"os"

	"ucenter/app/routes"

	"github.com/tabalt/gracehttp"
)

func main() {
	pid := os.Getpid()
	address := ":8080"

	log.Printf("process with pid %d serving %s.\n", pid, address)
	err := gracehttp.ListenAndServe(address, routes.New())
	if err != nil {
		log.Printf("process with pid %d stoped, error: %s.\n", pid, err)
	}

}
