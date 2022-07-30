package config

import (
	"fmt"

	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/pkg/model"
)

func ReadConfigFile() {
	configAPI, err := polaris.NewConfigAPI()
	if err != nil {
		panic(err)
	}

	var namespace = "backend"
	var fileGroup = "online"

	configFile, err := configAPI.GetConfigFile(namespace, fileGroup, "backend-online.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(configFile.GetContent())

	configFile.AddChangeListener(changeListener)
}

func changeListener(event model.ConfigFileChangeEvent) {
	fmt.Println(fmt.Sprintf("received change event. %+v", event.NewValue))
}
