package app

import (
	"log"
	"os"
	"strings"

	"github.com/pluto/pkg/util"
	yaml "gopkg.in/yaml.v2"
)

type YamlFormat struct {
	Global struct {
		LeftOffDir      string          `yaml:"leftOffDir"`
		startupBehavior StartupBehavior `yaml:"startupBehavior"`
		startupDir      string          `yaml:"startupDir"`
	}
}

type Configuration struct {
	filePath string
	data     YamlFormat
}

type StartupBehavior string

const (
	StartupBehaviorDefault   StartupBehavior = "left"
	StartupBehaviorHome      StartupBehavior = "home"
	StartupBehaviorSpecified StartupBehavior = "specified"
)

func NewConfiguration(filePath string,
	startupBehavior StartupBehavior,
	startupDir string) *Configuration {
	fullPath := strings.Replace(filePath, "~", util.HomeDir(), 1)
	yamlFile, err := os.ReadFile(fullPath)
	if err != nil {
		log.Panic(err)
	}
	data := YamlFormat{}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Panic(err)
	}
	data.Global.startupBehavior = startupBehavior
	data.Global.startupDir = startupDir
	return &Configuration{
		filePath: fullPath,
		data:     data,
	}
}

func (config *Configuration) resolveStartupDir() string {
	var resolve string
	startupBehavior := config.data.Global.startupBehavior
	startupDir := config.data.Global.startupDir
	if startupBehavior == StartupBehaviorSpecified {
		resolve = startupDir
	} else if startupBehavior == StartupBehaviorHome {
		resolve = util.HomeDir()
	} else {
		resolve = config.LeftOffDir()
	}
	if util.IsDir(resolve) {
		return resolve
	} else {
		return util.HomeDir()
	}
}

func (config *Configuration) LeftOffDir() string {
	return config.data.Global.LeftOffDir
}

func (config *Configuration) SetLeftOffDir(dir string) {
	config.data.Global.LeftOffDir = dir
}

func (config *Configuration) Serialize() {
	yamlFile, err := yaml.Marshal(&config.data)
	if err != nil {
		log.Panic(err)
	}
	err = os.WriteFile(config.filePath, yamlFile, 0644)
	if err != nil {
		log.Panic(err)
	}
}
