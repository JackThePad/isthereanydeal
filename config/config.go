package config

import (
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

	if tomlConfig.Config.NTFYTopic == "" {
		log.Fatalln("Ntfy topic not set")
	}
	if tomlConfig.Config.ITADAPIKey == "" {
		log.Fatal("No isthereadeal api key")
	}
	if tomlConfig.Config.SteamApiKey == "" {
		log.Fatal("No steam API key.")
	}

	return tomlConfig
}
