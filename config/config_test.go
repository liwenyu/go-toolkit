package config

import (
	"fmt"
	"os"
	"testing"
)

func TestIniConfig(t *testing.T) {
	t.Log("Testing")
	dir, err := os.Getwd()
	if err != nil {
		t.Log("error")
		return
	}

	appConfig := fmt.Sprintf("%s/%s", dir, "app.conf")
	conf, err := NewIniConfig(appConfig)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(conf.String("version"))
	t.Log(conf.GetSection("cas"))
	t.Log(conf.String("cas::login"))

	_ = conf.Set("db::port", "3306")
	_ = conf.SaveConfigFile(appConfig)
}
