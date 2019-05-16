package confs

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Conf 服务器配置
type Conf struct {
	MSconf MSConf `yaml:"msconf"`
}

// MSConf 微软oauth的配置
type MSConf struct {
	Callback string `yaml:"callback"`
	Clientid string `yaml:"clientid"`
	Scope    string `yaml:"scope"`
	Sec      string `yaml:"sec"`
}

var (
	MyConf = &Conf{}
)

func init() {
	yamlFile, err := ioutil.ReadFile("confs/conf.yml")
	panicIfErr(err)
	err = yaml.Unmarshal(yamlFile, MyConf)
	panicIfErr(err)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
