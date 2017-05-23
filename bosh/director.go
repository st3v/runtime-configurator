package bosh

import (
	boshdir "github.com/cloudfoundry/bosh-cli/director"
	boshuaa "github.com/cloudfoundry/bosh-cli/uaa"
	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshuit "github.com/cloudfoundry/bosh-cli/ui/task"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

func NewDirector(url, user, password, caCertPath string, uaa boshuaa.UAA, logger boshlog.Logger) (boshdir.Director, error) {
	config, err := boshdir.NewConfigFromURL(url)
	if err != nil {
		return nil, err
	}

	if config.CACert, err = readFile(caCertPath); err != nil {
		return nil, err
	}

	config.TokenFunc = boshuaa.NewClientTokenSession(uaa).TokenFunc

	if user != "" || password != "" {
		token, err := uaa.OwnerPasswordCredentialsGrant([]boshuaa.PromptAnswer{
			{"username", user},
			{"password", password},
		})

		if err != nil {
			return nil, err
		}

		config.TokenFunc = boshuaa.NewAccessTokenSession(token).TokenFunc
	}

	ui := boshui.NewNonInteractiveUI(boshui.NewConfUI(logger))
	defer ui.Flush()

	taskReporter := boshuit.NewReporter(ui, true)
	fileReporter := boshui.NewFileReporter(ui)

	return boshdir.NewFactory(logger).New(
		config,
		taskReporter,
		fileReporter,
	)
}
