package frame

import (
	"acme/library/util"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// ConfigPathEnv 配置文件名
const ConfigPathEnv = "ACME_HTTP_CONFIG_PATH_ENV"
const ConfigName = "acme_http_config.json"

// App 应用配置
type App struct {
	Host           string            `json:"host"`
	Port           uint32            `json:"port"`
	Env            map[string]string `json:"env"`
	TrustedProxies []string          `json:"trustedProxies"`
	Production     bool              `json:"production"`
	LogPath        string            `json:"logPath"`
}

// Config 配置文件
type Config struct {
	App App `json:"app"`
}

// 初始化配置文件
func iniConfig() (error, *Config) {
	configEnv := os.Getenv(ConfigPathEnv)
	if configEnv != "" {
		configPath := configEnv + "/" + ConfigName
		// 判断文件是否存在
		content, err := util.ReadFile(configPath)
		if err != nil {
			return err, nil
		}
		// 解析配置
		config := &Config{}
		err = json.Unmarshal(content, config)
		if err != nil {
			return err, nil
		}
		return nil, config
	}
	return errors.New(fmt.Sprintf("配置文件环境变量[ %s ]未设置!", ConfigPathEnv)), nil
}
