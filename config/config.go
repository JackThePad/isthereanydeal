package config

import (
	"errors"
	mytypes "isthereanydeal/my-types"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

func LoadTOML() mytypes.TOMLConfig {
	var tomlConfig mytypes.TOMLConfig

	toml_content, err := os.ReadFile("config.toml")

	if err != nil {
		log.Fatal(err)
	}

	_, err = toml.Decode(string(toml_content), &tomlConfig)

	if err != nil {
		log.Fatal(err)
	}

	if tomlConfig.Config.NTFYUrl == "" {
		log.Fatalln("Ntfy url not set")
	}
	if tomlConfig.Config.ITADAPIKey == "" {
		log.Fatal("No isthereadeal api key")
	}
	if tomlConfig.Config.SteamApiKey == "" {
		log.Fatal("No steam API key.")
	}
	if tomlConfig.Config.JsonName == "" {
		tomlConfig.Config.JsonName = "games.json"
		log.Println("Not json file stated using games.json")
	}

	if _, err := os.ReadFile(tomlConfig.Config.JsonName); errors.Is(err, os.ErrNotExist) {
		os.WriteFile(tomlConfig.Config.JsonName, []byte{}, 0644)
	}

	return tomlConfig
}
