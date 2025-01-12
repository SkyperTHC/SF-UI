package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

func ReadConfig() SfUI {
	// Any options not present in config.yaml will have default values
	sfuiConfig := getDefaultConfig()

	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Println("Failed to read file : ", err, ", Using default configs")
	}

	err = yaml.Unmarshal(data, &sfuiConfig)
	if err != nil {
		log.Println("Failed Unmarshal data", err)
	}

	sfuiConfig.CompiledClientConfig = getcompiledClientConfig(sfuiConfig)
	return sfuiConfig
}

func getDefaultConfig() SfUI {
	return SfUI{
		MaxWsTerminals:     10,
		ServerBindAddress:  "127.0.0.1:7171",
		XpraWSAddress:      "ws://127.0.0.1:2000/",
		Debug:              false,
		ShellCommand:       "bash",
		AddSfUIArgs:        false,
		SfUIOrigin:         "http://127.0.0.1:7171",
		DisableOriginCheck: true,
		DisableDesktop:     false,
	}
}

func getcompiledClientConfig(sfui SfUI) []byte {
	// Add any UI related configuration that has to be sent to client
	// Store it byte format, to prevent json marshalling on every request
	// See handleUIConfig()
	compConfig := []byte(fmt.Sprintf(
		`{"max_terminals":"%d","auto_login":%s,"desktop_disabled":%s}`,
		sfui.MaxWsTerminals,
		strconv.FormatBool(!sfui.AddSfUIArgs),   // Redirect client directly to dashboard if not in global mode.
		strconv.FormatBool(sfui.DisableDesktop), // Hide the GUI Option in UI
	))
	return compConfig
}
