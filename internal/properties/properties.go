package properties

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Properties struct {
	RunAddr    string `yml:"runaddr"`
	ConnString string `yml:"connstring"`
}

func New() Properties {
	var p Properties

	yamldata, err := os.ReadFile("config/props.yml")

	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamldata, &p)

	if err != nil {
		log.Fatal(err)
	}

	return p
}
