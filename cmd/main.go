package main

import (
	"log"
	"os"

	"github.com/pluto/pkg/app"
)

var (
	logFile, _ = os.OpenFile("pluto.log", os.O_RDWR|os.O_CREATE, 0666)
)

func logConfig() {
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("start loging...")
}

func main() {
	logConfig()
	app.NewApp(os.Args[1]).Run()
}
