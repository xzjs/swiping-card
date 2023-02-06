package lib

import (
	"io/ioutil"
	"log"
	"swiping-card/model"
	"sync"

	"gopkg.in/yaml.v3"
)

var conf *model.Conf
var once sync.Once

func Conf() *model.Conf {
	once.Do(func() {
		f, err := ioutil.ReadFile("conf.yaml")
		if err != nil {
			log.Fatal(err)
		}
		conf = &model.Conf{}
		if err = yaml.Unmarshal(f, conf); err != nil {
			log.Fatal(err)
		}
	})
	return conf
}
