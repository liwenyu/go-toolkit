// .conf 类配置文件
package config

import (
	"github.com/astaxie/beego/config"
	"os"
)

// new ini config
func NewIniConfig(configFile string) (conf config.Configer, err error) {
	// Determine whether the configuration file exists
	_, err = os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}
	}

	conf, err = config.NewConfig("ini", configFile)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
