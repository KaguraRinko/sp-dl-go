package config

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/XiaoMengXinX/sp-dl-go/logger"
	"os"
)

var CM = NewConfigManager()

func init() {
	CM.Initialize()
}

type Data struct {
	SpDc              string `json:"sp_dc"`
	AccessToken       string `json:"accessToken"`
	AccessTokenExpire int64  `json:"accessTokenExpire"`
	MetadataLanguage  string `json:"metadataLanguage"`
}

type Manager struct {
	configFileName string
	config         Data
	defaults       Data
}

func NewConfigManager() *Manager {
	defaults := Data{
		SpDc:              "",
		AccessToken:       "",
		AccessTokenExpire: -1,
		MetadataLanguage:  "zh-CN,zh;q=0.9,zh-HK;q=0.8,zh-TW;q=0.7, ja-JP;q=0.6,en-US;q=0.5,en;q=0.4",
	}
	return &Manager{
		configFileName: "config.json",
		config:         defaults,
		defaults:       defaults,
	}
}

func (cm *Manager) Initialize() {
	if _, err := os.Stat(cm.configFileName); errors.Is(err, os.ErrNotExist) {
		cm.writeConfig()
	}
}

func (cm *Manager) ReadConfig() error {
	data, err := os.ReadFile(cm.configFileName)
	if err != nil {
		return fmt.Errorf("unable to read config file: %w", err)
	}

	if err := json.Unmarshal(data, &cm.config); err != nil {
		return fmt.Errorf("unable to parse json config file: %w", err)
	}
	return nil
}

func (cm *Manager) Get() Data {
	return cm.config
}

func (cm *Manager) ReadAndGet() (Data, error) {
	if err := cm.ReadConfig(); err != nil {
		return Data{}, err
	}
	return cm.Get(), nil
}

func (cm *Manager) GetDefault() Data {
	return cm.defaults
}

func (cm *Manager) writeConfig() {
	data, err := json.MarshalIndent(cm.config, "", "  ")
	if err != nil {
		log.Fatalf("Unable to marshal config to json: %v", err)
	}

	if err := os.WriteFile(cm.configFileName, data, 0644); err != nil {
		log.Fatalf("Unable to write config to file: %v", err)
	}
}

func (cm *Manager) Set(newConfig Data) {
	cm.config = newConfig
	cm.writeConfig()
}
