package configurator

import (
	boshdir "github.com/cloudfoundry/bosh-cli/director"

	"github.com/st3v/runtime-configurator/rtconfig"
)

type getter interface {
	Get(string) (rtconfig.RuntimeConfig, error)
}

type setter interface {
	Set(string, rtconfig.RuntimeConfig) error
}

type Director interface {
	LatestRuntimeConfig(string) (boshdir.RuntimeConfig, error)
	UpdateRuntimeConfig(string, []byte) error
}

type Infoer interface {
	Info(string, string, ...interface{})
}

func New(director Director, infoer Infoer) *configurator {
	gs := &getsetter{director}
	return &configurator{getr: gs, setr: gs, log: infoer}
}

type configurator struct {
	getr getter
	setr setter
	log  Infoer
}

func (c *configurator) Add(name, runtimeConfigPath string) error {
	c.log.Info("configurator", "Adding runtime config - name: %q, path: %q", name, runtimeConfigPath)
	defer c.log.Info("configurator", "Done")

	rc, err := rtconfig.FromFile(runtimeConfigPath)
	if err != nil {
		return err
	}

	return c.process(name, func(orig rtconfig.RuntimeConfig) rtconfig.RuntimeConfig {
		return orig.Add(rc)
	})
}

func (c *configurator) Remove(name, runtimeConfigPath string) error {
	c.log.Info("configurator", "Removing runtime config - name: %q, path: %q", name, runtimeConfigPath)
	defer c.log.Info("configurator", "Done")

	rc, err := rtconfig.FromFile(runtimeConfigPath)
	if err != nil {
		return err
	}

	return c.process(name, func(orig rtconfig.RuntimeConfig) rtconfig.RuntimeConfig {
		return orig.Remove(rc)
	})
}

type updateFunc func(rtconfig.RuntimeConfig) rtconfig.RuntimeConfig

func (c *configurator) process(name string, update updateFunc) error {
	rc, err := c.getr.Get(name)
	if err != nil {
		return err
	}

	return c.setr.Set(name, update(rc))
}
