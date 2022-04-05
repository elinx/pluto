package main

import (
	"flag"
	"log"
	"os"

	"github.com/pluto/pkg/app"
	"github.com/pluto/pkg/util"
)

func logConfig(filename string) {
	logFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("start loging...")
}

func parseFlags() *app.Configuration {
	logFile := flag.String("log", "pluto.log", "log file")
	configFile := flag.String("config", "~/.config/pluto/config.yaml", "config file")
	startupBehavior := flag.String("startup", "left", "startup behavior")
	startupDir := flag.String("startup-dir", util.HomeDir(), "startup directory")
	flag.Parse()
	logConfig(*logFile)
	return app.NewConfiguration(*configFile,
		app.StartupBehavior(*startupBehavior),
		*startupDir)
}

func main() {
	config := parseFlags()
	// log.Println("startup directory:", config)
	app.NewApp(config).Run()
}
