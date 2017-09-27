package main

import (
	"fmt"

	"github.com/IBM-Bluemix/bluemix-cli-sdk/plugin"
)

type IoTPlugin struct{}

func main() {
	plugin.Start(new(IoTPlugin))
}

func (pluginDemo *IoTPlugin) Run(context plugin.PluginContext, args []string) {
	fmt.Println("Hi, this is my first plugin for IoT")
}

func (pluginDemo *IoTPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "WatsonIoTPlatform",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:        "watson-iot-platform",
				Alias:       "iot",
				Description: "say hello to IoT",
				Usage:       "bluemix iot",
			},
		},
	}
}

