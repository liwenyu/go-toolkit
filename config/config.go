// .conf 类配置文件
package config

import "github.com/astaxie/beego/config"

// new ini config
func NewIniConfig(configFile string) (conf config.Configer, err error) {
	conf, err = config.NewConfig("ini", configFile)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
