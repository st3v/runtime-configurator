package bosh

import (
	boshuaa "github.com/cloudfoundry/bosh-cli/uaa"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

func NewUAA(url, clientID, clientSecret, caCertPath string) (boshuaa.UAA, error) {
	config, err := boshuaa.NewConfigFromURL(url)
	if err != nil {
		return nil, err
	}

	config.Client, config.ClientSecret = clientID, clientSecret

	if config.CACert, err = readFile(caCertPath); err != nil {
		return nil, err
	}

	return boshuaa.NewFactory(boshlog.NewLogger(boshlog.LevelError)).New(config)
}
