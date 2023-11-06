package conf

import (
	"github.com/kelseyhightower/envconfig"
)

type conf struct {
	AuthToken string `required:"true" envconfig:"AUTH_TOKEN"`
	ApiKey    string `required:"true" envconfig:"API_KEY"`
}

var Conf conf

func init() {
	err := envconfig.Process("", &Conf)
	if err != nil {
		panic(err)
	}
}
