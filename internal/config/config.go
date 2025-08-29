package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return "", err
	}

	return homeDir + "/" + configFileName, nil
}

func Read() *Config {
	filePath, err := getConfigFilePath()
	if err != nil {
		log.Println("error getting the configuration file path: ", err)
		return &Config{}
	}

	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Println("error opening file:", err)
		return &Config{}
	}

	defer jsonFile.Close()

	bytes, _ := io.ReadAll(jsonFile)

	var config Config

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Println("error unmarshaling json:", err)
		return &Config{}
	}

	return &config
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username

	filePath, err := getConfigFilePath()
	if err != nil {
		log.Println(err)
		return err
	}

	dat, err := json.Marshal(cfg)
	if err != nil {
		log.Println(err)
		return err
	}

	err = os.WriteFile(filePath, dat, 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
