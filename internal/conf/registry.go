package conf

import (
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Registry struct {
	ClientConf *constant.ClientConfig
	Endpoints  []*constant.ServerConfig
}

func LoadRegistry(path string) (reg *Registry, err error) {
	var raw []byte
	if raw, err = ioutil.ReadFile(path); err != nil {
		return
	}
	if err = yaml.Unmarshal(raw, reg); err != nil {
		return
	}
	err = reg.verify()
	return
}

func (r *Registry) verify() error {
	return nil
}
