package server

import (
	"io/ioutil"

	"github.com/linyuan-t/quick-search/backend/config"
	"github.com/linyuan-t/quick-search/backend/services"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func initServer(config *config.Config) (*services.Services, error) {
	// 初始化配置
	if err := InitConfigFromFile(config); err != nil {
		return nil, err
	}

	s := services.NewServices(config)

	return s, nil
}

// InitConfigFromFile ...
func InitConfigFromFile(config *config.Config) error {

	// 1.解析 YAML 配置文件
	bytes, err := ioutil.ReadFile(config.FileName)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return nil
	}

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil
	}

	return nil
}
