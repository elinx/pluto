package main

import (
	"log"
	"os"

	"github.com/pluto/pkg/app"
)

func logConfig() {
	f, err := os.OpenFile("pluto.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("start loging...")
}

func main() {
	logConfig()
	app.NewApp(os.Args[1]).Run()
}
