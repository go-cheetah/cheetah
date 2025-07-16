package main

import (
	"log"

	"github.com/go-cheetah/cheetah/cmd/cheetah/app"
)

func main() {
	cmd := app.NewServerCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
