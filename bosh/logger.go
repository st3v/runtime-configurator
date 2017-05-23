package bosh

import boshlog "github.com/cloudfoundry/bosh-utils/logger"

func NewLogger(level string) (boshlog.Logger, error) {
	if level == "" {
		return boshlog.NewLogger(boshlog.LevelNone), nil
	}

	l, err := boshlog.Levelify(level)
	if err != nil {
		return nil, err
	}

	return boshlog.NewLogger(l), err
}
