package configurator

import (
	"bytes"
	"strings"

	"github.com/st3v/runtime-configurator/rtconfig"
)

type getsetter struct {
	dir Director
}

func (gs *getsetter) Get(name string) (rtconfig.RuntimeConfig, error) {
	rc, err := gs.dir.LatestRuntimeConfig(name)
	if err != nil && err.Error() != "No runtime config" {
		return rtconfig.RuntimeConfig{}, err
	}

	return rtconfig.Read(strings.NewReader(rc.Properties))
}

func (gs *getsetter) Set(name string, rc rtconfig.RuntimeConfig) error {
	buf := new(bytes.Buffer)

	if err := rc.Write(buf); err != nil {
		return err
	}

	return gs.dir.UpdateRuntimeConfig(name, buf.Bytes())
}
